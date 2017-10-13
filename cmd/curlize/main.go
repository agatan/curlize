package main

import (
	"log"
	"net/http"

	"github.com/agatan/curlize"
	"github.com/agatan/groxy"
)

func main() {
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

	if err := http.ListenAndServe(":8080", proxy); err != nil {
		panic(err)
	}
}
