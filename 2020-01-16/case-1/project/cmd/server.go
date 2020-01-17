package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"tech.avito.ru/service/internal/handler"
	"tech.avito.ru/service/internal/service"
)

func main() {
	http.HandleFunc("/echo", func(rw http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if delay := req.Form.Get("delay"); delay != "" {
			if d, err := time.ParseDuration(delay); err == nil {
				time.Sleep(d)
			}
		}
		out := io.MultiWriter(rw, os.Stdout)
		io.Copy(out, req.Body)
		fmt.Fprintln(out)
	})
	http.Handle("/", handler.Proxy(service.New()))
	log.Fatalln(http.ListenAndServe(":80", http.DefaultServeMux))
}
