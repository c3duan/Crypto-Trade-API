package helper

import (
	"log"
	"net/http"
	"strings"
)

func WithSlashTrimming(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
        log.Println(r.RequestURI)
        next.ServeHTTP(w, r)
    })
}
