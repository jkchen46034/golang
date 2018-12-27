// a web server with regular expression as routes, with done pattern
// dynamic route inherited from this link: https://gist.github.com/reagent/043da4661d2984e9ecb1ccb5343bf438

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
)

func main() {
	done := serveHttp()
	<-done
}

func serveHttp() chan int {
	done := make(chan int)
	go func() {
		httpServer()
		done <- 1
	}()
	return done
}

func httpServer() {
	webServer := NewWebServer()

	webServer.Handle(`^/hello$`, func(resp *Response, req *Request) {
		resp.Text(http.StatusOK, "Hello world")
	})

	webServer.Handle(`/hello/([\w\._-]+)$`, func(resp *Response, req *Request) {
		resp.Text(http.StatusOK, fmt.Sprintf("Hello %s", req.Params[0]))
	})

	webServer.Handle(`/$`, func(resp *Response, req *Request) {
		resp.Text(http.StatusOK, "Index route")
	})

	err := http.ListenAndServe(":9999", webServer)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}
	log.Println("clubhouse server exited")
}

type Handler func(*Response, *Request)

type Route struct {
	Pattern *regexp.Regexp
	Handler Handler
}

type WebServer struct {
	Routes       []Route
	DefaultRoute Handler
}

func NewWebServer() *WebServer {
	webServer := &WebServer{
		DefaultRoute: func(resp *Response, req *Request) {
			resp.Text(http.StatusNotFound, "Not found")
		},
	}

	return webServer
}

func (server *WebServer) Handle(pattern string, handler Handler) {
	re := regexp.MustCompile(pattern)
	server.Routes = append(server.Routes, Route{Pattern: re, Handler: handler})
}

func (server *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := &Request{Request: r}
	resp := &Response{w}

	for _, rt := range server.Routes {
		if matches := rt.Pattern.FindStringSubmatch(r.URL.Path); len(matches) > 0 {
			if len(matches) > 1 {
				req.Params = matches[1:]
			}

			rt.Handler(resp, req)
			return
		}
	}

	server.DefaultRoute(resp, req)
}

type Request struct {
	*http.Request
	Params []string
}

type Response struct {
	http.ResponseWriter
}

func (r *Response) Text(code int, body string) {
	r.Header().Set("Content-Type", "text/plain")
	r.WriteHeader(code)

	io.WriteString(r, fmt.Sprintf("%s\n", body))
}
