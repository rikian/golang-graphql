package app

import (
	"golang/graphql/app/config"
	"golang/graphql/app/controllers"
	"golang/graphql/app/graph"
	"golang/graphql/app/graph/generated"
	"golang/graphql/app/middlewares"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LintenAndServe(address string) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config.ConnectDB(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	// graph init
	var h *handler.Server = handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	// graph playground init
	var p http.HandlerFunc = playground.Handler("GraphQL", "/query")

	gin.SetMode(gin.ReleaseMode)

	var server *gin.Engine = gin.New()

	server.Use(middlewares.CORSMiddleware())

	server.POST("/query", controllers.GraphqlHandler(h))
	server.GET("/", controllers.GraphPlayGroundHandler(p))

	log.Print("listen and server at " + address)
	server.Run(address)
}
