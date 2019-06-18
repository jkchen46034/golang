/* how to figure out "muliptle header calls" error */
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
)

type debugLogger struct{}

func (d debugLogger) Write(p []byte) (n int, err error) {
	s := string(p)
	if strings.Contains(s, "multiple response.WriteHeader") {
		debug.PrintStack()
	}
	return os.Stderr.Write(p)
}

func myhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	go HandleIndex(w, r)
}
func main() {

	logger := log.New(debugLogger{}, "", 0)

	mux := http.NewServeMux()
	th := http.HandlerFunc(myhandler)
	mux.Handle("/", th)

	server := &http.Server{
		Addr:     ":3001",
		Handler:  mux,
		ErrorLog: logger,
	}

	fmt.Println("Starting Server...")
	server.ListenAndServe()
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Hello, World!"))
}

/*
$ go run main.go
$ curl localhost:3001
Starting Server...
/
goroutine 21 [running]:
runtime/debug.Stack(0xc000032d38, 0x4be9a9, 0xc000124090)
	/usr/local/go/src/runtime/debug/stack.go:24 +0xa7
runtime/debug.PrintStack()
	/usr/local/go/src/runtime/debug/stack.go:16 +0x22
main.debugLogger.Write(0xc000124060, 0x2a, 0x30, 0x30, 0x0, 0xc000124060)
	/home/jk/dev/gocode/src/github.com/golang/multiple_headers/main.go:17 +0xde
log.(*Logger).Output(0xc00009e140, 0x2, 0xc000124030, 0x29, 0x0, 0x0)
	/usr/local/go/src/log/log.go:172 +0x204
log.(*Logger).Printf(0xc00009e140, 0x6b61e5, 0x29, 0x0, 0x0, 0x0)
	/usr/local/go/src/log/log.go:179 +0x7e
net/http.(*Server).logf(0xc000088f70, 0x6b61e5, 0x29, 0x0, 0x0, 0x0)
	/usr/local/go/src/net/http/server.go:2977 +0x6a
net/http.(*response).WriteHeader(0xc00011c000, 0xc8)
	/usr/local/go/src/net/http/server.go:1102 +0x26a
main.HandleIndex(0x6f1660, 0xc00011c000, 0xc000116000)
	/home/jk/dev/gocode/src/github.com/golang/multiple_headers/main.go:45 +0x3e
created by main.myhandler
	/home/jk/dev/gocode/src/github.com/golang/multiple_headers/main.go:24 +0x99
http: multiple response.WriteHeader calls
*/

