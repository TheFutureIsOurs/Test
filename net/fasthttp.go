package net

import (
	"encoding/json"
	"log"
	_ "net/http/pprof"
	"time"

	"github.com/valyala/fasthttp"
)

var (
	strContentType     = []byte("Content-Type")
	strApplicationJSON = []byte("application/json; charset=utf-8")
	strConnection      = []byte("Connection")
	strKeepAlive       = []byte("keep-alive")
)

type MyHandler struct {
}

type Response struct {
	ErrCode int    `json:"errCode"`
	Message string `json:"message"`
	Hit     int    `json:"hit"`
}

func (h *MyHandler) HandleFastHTTP(ctx *fasthttp.RequestCtx) {
	response := Response{
		ErrCode: 0,
		Message: "success",
		Hit:     1,
	}
	outPutJson(ctx, response)

	//ctx.Logger().Printf("status:%v", ctx.Response.Header.StatusCode())
}

func outPutJson(ctx *fasthttp.RequestCtx, obj interface{}) {
	ctx.Response.Header.SetCanonical(strContentType, strApplicationJSON)
	ctx.Response.SetStatusCode(200)
	json.NewEncoder(ctx).Encode(obj)
}

func main() {
	myHandler := &MyHandler{}

	s := &fasthttp.Server{
		Name:               "map server",
		Handler:            myHandler.HandleFastHTTP,
		ReadTimeout:        1 * time.Second,
		WriteTimeout:       3 * time.Second,
		IdleTimeout:        60 * time.Second,
		TCPKeepalive:       true,
		TCPKeepalivePeriod: 60 * time.Second,
	}
	log.Fatal(s.ListenAndServe(":8081"))
}
