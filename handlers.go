package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func HandlerHello(ctx *gin.Context) {
    ctx.String(http.StatusOK, "Hello string")
}

func HandlerPing(ctx *gin.Context) {
    pong, err := Cache.Ping()

    if err != nil {
        ctx.String(http.StatusInternalServerError, err.Error())
        return
    }

    ctx.String(http.StatusOK, pong)
}

// Allow info - name, email, password
func HandlerRegister(ctx *gin.Context) {
    var user UserDB
    err := ctx.ShouldBindJSON(&user)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newPassword, err := CryptPassword(user)

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    user.Password = newPassword
    newId, err := DbClient.CreateNewUser(user)

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    newToken, err := GenerateToken()

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    err = Cache.Save(newToken, newId, 0)
    
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"token": newToken})
}

// For login take only email (it's unique)
// Token may not be in cache, (expired or lost)
//  so need to create new one and save
func HandlerLogin(ctx *gin.Context) {
    var user UserDB
    err := ctx.ShouldBindJSON(&user)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
}

func HandlerGetByToken(ctx *gin.Context) {
    ctx.JSON(http.StatusOK, gin.H{"error": "Not implemented"})
}

func HandlerGetByEmail(ctx *gin.Context) {
    var record map[string]string
    err := ctx.ShouldBindJSON(&record)

    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, err := DbClient.GetUserByEmail(record["email"])

    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    ctx.JSON(http.StatusOK, user)
}
