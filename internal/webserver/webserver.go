package webserver

import (
	"fmt"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/debug"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cble-platform/backend/config"
	"github.com/cble-platform/backend/ent"
	"github.com/cble-platform/backend/graph"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type CBLEWebserver struct {
	router *gin.Engine
	config *config.Config
}

func graphqlHandler(config *config.Config, client *ent.Client) gin.HandlerFunc {
	cors_urls := []string{"http://localhost", "http://localhost:8080", "http://localhost:3000"}
	if len(config.Server.AllowedOrigins) > 0 {
		cors_urls = config.Server.AllowedOrigins
	}

	srv := handler.New(graph.NewSchema(client))

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

func New(config *config.Config, client *ent.Client) *CBLEWebserver {
	r := gin.Default()

	api := r.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	gql := api.Group("/graphql")

	gql.Any("/query", graphqlHandler(config, client))
	if config.Debug {
		gql.Any("/playground", playgroundHandler())
	}

	return &CBLEWebserver{
		router: r,
		config: config,
	}
}

func (w *CBLEWebserver) Listen() {
	addr := fmt.Sprintf("%s:%d", w.config.Server.Hostname, w.config.Server.Port)
	if err := w.router.Run(addr); err != nil {
		logrus.Fatalf("failed to start gin router")
	}
}
