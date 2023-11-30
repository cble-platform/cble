package webserver

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cble-platform/cble-backend/auth"
	"github.com/cble-platform/cble-backend/config"
	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/graph"
	"github.com/cble-platform/cble-backend/internal/permissionengine"
	"github.com/cble-platform/cble-backend/providers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type CBLEWebserver struct {
	webserver *http.Server
	config    *config.Config
}

func graphqlHandler(config *config.Config, client *ent.Client, cbleServer *providers.CBLEServer, permissionEngine *permissionengine.PermissionEngine) gin.HandlerFunc {
	cors_urls := []string{"http://localhost", "http://localhost:8080", "http://localhost:3000"}
	if len(config.Server.AllowedOrigins) > 0 {
		cors_urls = config.Server.AllowedOrigins
	}

	srv := handler.New(graph.NewSchema(client, cbleServer, permissionEngine))

	srv.AddTransport(&transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				for _, url := range cors_urls {
					if r.Header.Get("origin") == url {
						return true
					}
				}
				return false
			},
		},
	})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	if config.Server.GQlTrace {
		srv.Use(&debug.Tracer{})
	}

	if config.Server.GQlIntrospection {
		srv.Use(extension.Introspection{})
	}

	return func(c *gin.Context) {
		srv.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("CBLE GraphQL Playground", "/api/graphql/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func New(config *config.Config, client *ent.Client, cbleServer *providers.CBLEServer, permissionEngine *permissionengine.PermissionEngine) *CBLEWebserver {
	// Use default gin route settings
	r := gin.Default()

	// Set the gin mode based on debug value
	if config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Set up CORS (see https://github.com/gin-contrib/cors)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     config.Server.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
	}))

	// General API route group
	api := r.Group("/api")

	// Authentication route group (/api/auth)
	authGroup := api.Group("/auth")
	// Login endpoint
	authGroup.POST("/login", auth.Login(config, client))
	// Logout endpoint
	authGroup.DELETE("/logout", auth.Logout(config))

	// GraphQL route group (/api/graphql)
	gqlGroup := api.Group("/graphql")
	// Inject gin context into graphql context
	gqlGroup.Use(graph.GinContextToContextMiddleware())
	// Authenticate all graphql requests
	gqlGroup.Use(auth.AuthMiddleware(config, client))
	// Direct all graphql queries to graphql handler
	gqlGroup.Any("/query", graphqlHandler(config, client, cbleServer, permissionEngine))
	// Only enable graphql playground on debug
	if config.Debug {
		gqlGroup.Any("/playground", playgroundHandler())
	}

	// Set up http server for gin (allows interactive shutdown)
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Server.Hostname, config.Server.Port),
		Handler: r,
	}

	return &CBLEWebserver{
		webserver: srv,
		config:    config,
	}
}

func (w *CBLEWebserver) Listen(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	go func() {
		if err := w.webserver.ListenAndServe(); err != http.ErrServerClosed {
			logrus.Fatalf("failed to start webserver")
		}
	}()

	<-ctx.Done()
	logrus.Warnf("Gracefully shutting down CBLE webserver...")
	w.webserver.Shutdown(ctx)
}
