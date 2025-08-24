package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	r := router.New()

	r.GET("/", Welcome)
	r.GET("/{userId}", GetUserById)
	log.Println("Application Started ms user api: 8081")
	log.Fatal(fasthttp.ListenAndServe(":8081", r.Handler))
}

func Welcome(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Welcome to Demon's World!")
}

func GetUserById(ctx *fasthttp.RequestCtx) {
	userId := ctx.UserValue("userId").(string)
	fmt.Fprintf(ctx, "Hello, %s!", userId)
}
