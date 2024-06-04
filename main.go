package main

import (
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/biohackerellie/goMailer/env"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {
	r := chi.NewRouter()
	urls := strings.Split(env.GetConfig().AllowedDomains, ",")
	ipv4_regex := `^(localhost|10(\.\d{1,3}){3}|192\.168(\.\d{1,3}){2})$`
	re := regexp.MustCompile(ipv4_regex)

	var allowedOrigins []string
	for _, url := range urls {
		if re.MatchString(url) || strings.Contains(url, "*") {
			allowedOrigins = append(allowedOrigins, url)
		}
	}
	allowedOrigins = append(allowedOrigins, ipv4_regex)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: allowedOrigins,
		AllowedMethods: []string{"GET", "POST"},
	}))
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("🏳️‍🌈🏳️‍🌈🏳️‍🌈"))
	})

	r.Mount("/email", EmailRequest{}.Routes())
	http.ListenAndServe(":6969", r)

}
