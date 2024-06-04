package main

import (
	"encoding/json"
	"net/http"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/biohackerellie/goMailer/env"
	"github.com/go-chi/chi/v5"
)

type EmailRequest struct {
	Key     string
	To      string
	From    string
	Subject string
	HTML    string
}

type EmailRoute struct{}

func (rs EmailRequest) Routes() chi.Router {
	r := chi.NewRouter()
	r.Post("/", rs.sendEmail)
	return r
}

func (rs EmailRequest) sendEmail(w http.ResponseWriter, r *http.Request) {
	var e EmailRequest
	var (
		config = env.GetConfig()
	)

	auth := smtp.PlainAuth("", config.User, config.Pass, "smtp.gmail.com")

	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Key := e.Key
	if Key != config.APIKey {
		http.Error(w, "Invalid API Key", http.StatusUnauthorized)
		return
	}

	from := mail.Address{Name: e.From, Address: config.User}
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	subject := "Subject: " + e.Subject + "!\n"
	header := "From: " + from.Name + "<" + from.Address + ">\n"
	msg := []byte(header + subject + mime + "\n" + e.HTML)

	to := e.To
	recipients := strings.Split(to, ", ")

	errs := smtp.SendMail("smtp.gmail.com:587", auth, config.User, recipients, msg)
	if errs != nil {
		http.Error(w, errs.Error(), http.StatusInternalServerError)
		return
	}
}
