package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TODO: all sending responce must be like:
//       status: ok | error
//       message: "..." (message exist only if status = error)

func HandlerHello(ctx *gin.Context) {
    ctx.String(http.StatusOK, "Hello string")
}

// func HandlerPing(ctx *gin.Context) {
//     pong, err := Cache.Ping()
//
//     if err != nil {
//         ctx.String(http.StatusInternalServerError, err.Error())
//         return
//     }
//
//     ctx.String(http.StatusOK, pong)
// }

/*
Register a new user if not already exist in DB
Send {"message": "ok"} when new user registered
*/
func HandlerRegister(ctx *gin.Context) {
    var user_in UserIn
    err := ctx.BindJSON(&user_in) // convert request body from json to struct

    if err != nil || user_in.Email == "" || user_in.Password == "" {
        ctx.JSON(
            http.StatusBadRequest,
            gin.H{"status": "error", "message": "Invalid data"})
        return
    }

    password_crypted, err := CryptPassword(user_in.Password)

    if err != nil {
        ctx.JSON(
            http.StatusBadRequest,
            gin.H{"status": "error", "message": "Invalid data"})
        return
    }

    var user_db UserDB = UserDB{
        Email: user_in.Email,
        Password: password_crypted,
    }

    err = DbClient.CreateNewUser(user_db)

    if err != nil {
        message := err.Error()

        if DbClient.IsUserEmailExist(user_in.Email) {
            message = "User already exist"
        }

        ctx.JSON(
            http.StatusBadRequest,
            gin.H{"status": "error", "message": message})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}

/*
Login a user 
Return {"status": "allowed" | "denied"}
*/
func HandlerLogin(ctx *gin.Context) {
    var user_in UserIn
    err := ctx.ShouldBindJSON(&user_in)

    if err != nil || user_in.Email == "" || user_in.Password == "" {
        ctx.JSON(
            http.StatusBadRequest,
            gin.H{"status": "denied", "message": err.Error()})
        return
    }

    user_db, err := DbClient.GetUserByEmail(user_in.Email)

    if err != nil {
        ctx.JSON(
            http.StatusBadRequest,
            gin.H{"status": "denied", "message": "User not exist"})
        return
    }

    password_crypted, err := CryptPassword(user_in.Email)

    if err != nil {
        ctx.JSON(
            http.StatusBadRequest,
            gin.H{"status": "denied", "message": err.Error()})
        return
    }

    if user_db.Password != password_crypted {
        ctx.JSON(
            http.StatusUnauthorized,
            gin.H{"status": "denied", "message": "Wrong password"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "allowed"})
}

// func HandlerGetByToken(ctx *gin.Context) {
//     ctx.JSON(http.StatusOK, gin.H{"error": "Not implemented"})
// }

/*
Check if user exist in DB
Return {"status": "exist" | "not exist"}
*/
func HandlerIsExist(ctx *gin.Context) {
    var user_in UserIn // change to UserOut
    err := ctx.ShouldBindJSON(&user_in)

    if err != nil || user_in.Email == "" {
        fmt.Printf("IsExistERROR: '%#v'\n", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "not exist", "message": "Invalid data"})
        return
    }

    if !DbClient.IsUserEmailExist(user_in.Email) {
        ctx.JSON(http.StatusBadRequest, gin.H{"status": "not exist", "message": "Can't find user in DB"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
}
