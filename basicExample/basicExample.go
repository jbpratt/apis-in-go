package main

import (
	"fmt"
	"io"
	"net/http"
	"time"

	restful "github.com/emicklei/go-restful"
)

func main() {
	// create a web service
	webservice := new(restful.WebService)
	// new route and attach handler
	webservice.Route(webservice.GET("/ping").To(pingTime))
	// add service to app
	restful.Add(webservice)
	http.ListenAndServe(":8000", nil)
}

func pingTime(req *restful.Request, resp *restful.Response) {
	// write to the response
	io.WriteString(resp, fmt.Sprintf("%s", time.Now()))
}
