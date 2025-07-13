package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/smtp"
	"net/url"
	"os"

	"github.com/go-chi/chi"
)

type RecaptchaResponse struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
	Score       float64  `json:"score"`
	Action      string   `json:"action"`
}

func verifyRecaptcha(token string) (bool, error) {
	secretKey := os.Getenv("RECAPTCHA_SECRET")
	if secretKey == "" {
		return false, fmt.Errorf("RECAPTCHA_SECRET environment variable not set")
	}

	resp, err := http.PostForm("https://www.google.com/recaptcha/api/siteverify", url.Values{
		"secret":   {secretKey},
		"response": {token},
	})
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var recaptchaResp RecaptchaResponse
	if err := json.NewDecoder(resp.Body).Decode(&recaptchaResp); err != nil {
		return false, err
	}

	return recaptchaResp.Success && recaptchaResp.Score >= 0.5, nil
}

func main() {
	r := chi.NewRouter()

	// Serve the static HTML file with the form
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "form.html")
	})

	// Handle form submissions
	r.Post("/submit", func(w http.ResponseWriter, r *http.Request) {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get the reCAPTCHA token
		recaptchaToken := r.Form.Get("g-recaptcha-response")
		if recaptchaToken == "" {
			http.Error(w, "reCAPTCHA token missing", http.StatusBadRequest)
			return
		}

		// Verify reCAPTCHA
		valid, err := verifyRecaptcha(recaptchaToken)
		if err != nil {
			http.Error(w, "reCAPTCHA verification failed: "+err.Error(), http.StatusInternalServerError)
			return
		}
		if !valid {
			http.Error(w, "reCAPTCHA verification failed", http.StatusBadRequest)
			return
		}

		// Get the email address and name from the form
		email := r.Form.Get("email")
		name := r.Form.Get("name")
		message := r.Form.Get("message")

		password := os.Getenv("SMTP_PASSWORD")
		from := "pavel.anni@gmail.com"
		to := "pavel.anni@gmail.com"

		// Compose the email message
		subject := "Feedback from " + name
		body := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s\n", name, email, message)
		msg := []byte("To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"\r\n" +
			body + "\r\n")

		// Send the email
		err = smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", "pavel.anni@gmail.com", password, "smtp.gmail.com"),
			from, []string{to}, msg)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Redirect to the thank you page
		http.Redirect(w, r, "https://pythonicadventure.com/thankyou", http.StatusSeeOther)
	})

	// Start the server
	http.ListenAndServe("0.0.0.0:8080", r)
}
