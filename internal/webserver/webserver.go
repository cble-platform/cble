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
	"github.com/cble-platform/cble-backend/config"
	"github.com/cble-platform/cble-backend/ent"
	"github.com/cble-platform/cble-backend/graph"
	"github.com/cble-platform/cble-backend/providers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type CBLEWebserver struct {
	webserver *http.Server
	config    *config.Config
}

func graphqlHandler(config *config.Config, client *ent.Client, cbleServer *providers.CBLEServer) gin.HandlerFunc {
	cors_urls := []string{"http://localhost", "http://localhost:8080", "http://localhost:3000"}
	if len(config.Server.AllowedOrigins) > 0 {
		cors_urls = config.Server.AllowedOrigins
	}

	srv := handler.New(graph.NewSchema(client, cbleServer))

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

func New(config *config.Config, client *ent.Client, cbleServer *providers.CBLEServer) *CBLEWebserver {
	r := gin.Default()

	api := r.Group("/api")

	gql := api.Group("/graphql")

	gql.Any("/query", graphqlHandler(config, client, cbleServer))
	if config.Debug {
		gql.Any("/playground", playgroundHandler())
	}

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
