package controllers

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
)

func GraphqlHandler(h *handler.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("set-cookie", "user=12345")
		h.ServeHTTP(ctx.Writer, ctx.Request)
	}
}

func GraphPlayGroundHandler(p http.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
