/*
Allowed port:
    8000 - for http
    4430 - for https
*/
package main

import (
    "flag"
    "os"
    "log"
    "fmt"

    "github.com/gin-gonic/gin"
)

const (
    DEFAULT_HTTP = 8000
    DEFAULT_HTTPS = 4430

    // name must be same as in the compose file, links section
    REDIS_DEFAULT_IP = "redis"
    REDIS_DEFAULT_PORT = 6379
    REDIS_DEFAULT_PASSWORD = ""
    REDIS_DEFAULT_DB = 0
)

var (
    server_ip string
    server_port_http int
    server_address_http string

    server_port_https int
    server_address_https string

    RedisClient *RedisDB
    redisAddres string
    DbClient *DB
)

func main() {
    router := gin.Default()

    router.GET("/hello", HandlerHello)
    router.GET("/ping", HandlerPing)

    router.Run(server_address_http)
}

func init() {
    // gin.SetMode(gin.ReleaseMode)

    flag.StringVar(&server_ip, "ip", "0.0.0.0", "IP adress for server")
    flag.IntVar(&server_port_http, "http", DEFAULT_HTTP, "http port for listen")
    flag.IntVar(&server_port_https, "https", DEFAULT_HTTPS, "https port for listen")
    // TODO: add flags for custom redis settings
    flag.Parse()

    server_address_http = fmt.Sprintf("%s:%d", server_ip, server_port_http)
    server_address_https = fmt.Sprintf("%s:%d", server_ip, server_port_https)
    redisAddres = fmt.Sprintf("%s:%d", REDIS_DEFAULT_IP, REDIS_DEFAULT_PORT)

    RedisClient = RedisDBNew().Connect(redisAddres, REDIS_DEFAULT_PASSWORD, REDIS_DEFAULT_DB)
    DbClient = DBNew()
    err := DbClient.Connect(DB_DEFAULT_HOST, DB_DEFAULT_USER, DB_DEFAULT_PASSWORD, DB_DEFAULT_DBNAME, DB_DEFAULT_PORT, DB_DEFAULT_SSLMODE, DB_DEFAULT_TIMEZONE)

    log.Printf("Server           PID: '%d'\n", os.Getgid())
    log.Printf("Server http  address: '%s'\n", server_address_http)
    log.Printf("Server https address: '%s'\n", server_address_https)
    log.Printf("Redis DB     address: '%s'\n", redisAddres)
    log.Printf("Redis DB  client ptr: '%p'\n", RedisClient.client)

    if err != nil {
        log.Printf("DB connect ERROR: '%s'\n", err.Error())
    } else {
        log.Printf("DB connect Success\n")
    }
}
