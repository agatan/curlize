package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/agatan/curlize"
	"github.com/agatan/groxy"
)

type options struct {
	address string
}

func main() {
	var op options

	flag.StringVar(&op.address, "addr", ":8080", "listening address")

	flag.Parse()

	proxy := &groxy.ProxyServer{
		HTTPSAction: groxy.HTTPSActionMITM,
	}
	proxy.Use(func(h groxy.Handler) groxy.Handler {
		return func(r *http.Request) (*http.Response, error) {
			curl, err := curlize.Curlize(r)
			if err != nil {
				log.Println(err)
			} else {
				log.Println(curl.String())
			}
			return h(r)
		}
	})

	if err := http.ListenAndServe(op.address, proxy); err != nil {
		panic(err)
	}
}
