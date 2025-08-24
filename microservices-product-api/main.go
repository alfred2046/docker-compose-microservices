package main

import (
	"fmt"
	"log"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func main() {
	r := router.New()

	r.GET("/products", ListProducts)
	r.GET("/product/{productId}", GetProductById)
	log.Println("Application Started ms product api: 8082")
	log.Fatal(fasthttp.ListenAndServe(":8082", r.Handler))
}

func ListProducts(ctx *fasthttp.RequestCtx) {
	ctx.WriteString("Oracle, IBM, Microsoft, Amazon, Shopify")
}

func GetProductById(ctx *fasthttp.RequestCtx) {
	productId := ctx.UserValue("productId").(string)
	fmt.Fprintf(ctx, "Product by Id %s = Oracle, IBM, Microsoft", productId)
}
