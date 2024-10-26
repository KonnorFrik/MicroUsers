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
)

var (
    server_ip string
    server_port_http int
    server_address_http string

    server_port_https int
    server_address_https string
)

func main() {
    router := gin.Default()

    router.GET("/hello", HandlerHello)

    router.Run(server_address_http)
}

func init() {
    flag.StringVar(&server_ip, "ip", "0.0.0.0", "IP adress for server")
    flag.IntVar(&server_port_http, "http", DEFAULT_HTTP, "http port for listen")
    flag.IntVar(&server_port_https, "https", DEFAULT_HTTPS, "https port for listen")
    flag.Parse()

    server_address_http = fmt.Sprintf("%s:%d", server_ip, server_port_http)
    server_address_https = fmt.Sprintf("%s:%d", server_ip, server_port_https)

    // gin.SetMode(gin.ReleaseMode)

    log.Printf("Server           PID: '%d'\n", os.Getgid())
    log.Printf("Server http  address: '%s'\n", server_address_http)
    log.Printf("Server https address: '%s'\n", server_address_https)
}
