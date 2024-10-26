package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func HandlerHello(ctx *gin.Context) {
    ctx.String(http.StatusOK, "Hello string")
}

