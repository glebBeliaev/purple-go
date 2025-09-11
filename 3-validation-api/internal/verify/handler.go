package verify

import (
	"encoding/json"
	"fmt"
	"http/3-validation-api/configs"
	"http/3-validation-api/internal/repository"
	"net"
	"net/http"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type VerifyHandler struct {
	Config *configs.Config
}
type VerifyHandlerDeps struct {
	*configs.Config
}

const storePath = "3-validation-api/internal/repository/verifications.json"

func NewAuthHandler(router *http.ServeMux, deps VerifyHandlerDeps) {
	handler := &VerifyHandler{
		Config: deps.Config}

	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VerifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var payload SendRequest
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		hash, err := generateHash(16)
		if err != nil {
			http.Error(w, "failed to generate hash", http.StatusInternalServerError)
			return
		}
		verifications, err := repository.LoadVerification(storePath)
		if err != nil {
			http.Error(w, "failed to load store", http.StatusInternalServerError)
			fmt.Println("ошибка загрузки")
			return
		}
		verifications = append(verifications, repository.Verification{
			Email: payload.Email,
			Hash:  hash,
		})
		if err := repository.SaveVerifications(storePath, verifications); err != nil {
			http.Error(w, "failed to save store", http.StatusInternalServerError)
			fmt.Println("ошибка сохранения")
			return
		}

		baseURL := "http://localhost:8081"
		verifyURL := fmt.Sprintf("%s/verify/%s", baseURL, hash)

		e := email.NewEmail()
		e.From = handler.Config.Mail.Address
		e.To = []string{payload.Email}
		e.Subject = "Email verification"
		e.Text = []byte("Перейдите по ссылке для подтверждения: " + verifyURL)
		e.HTML = []byte(fmt.Sprintf(
			`<p>Перейдите по ссылке для подтверждения:</p><p><a href="%s">%s</a></p>`,
			verifyURL, verifyURL))

		smtpAddr := handler.Config.Mail.SMTP
		host, _, _ := net.SplitHostPort(smtpAddr)
		auth := smtp.PlainAuth("", handler.Config.Mail.Address, handler.Config.Mail.Password, host)

		if err := e.Send(smtpAddr, auth); err != nil {
			http.Error(w, "failed to send email: "+err.Error(), http.StatusInternalServerError)
			return
		}

	}
}

func (handler *VerifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		if hash == "" {
			http.Error(w, "hash is required", http.StatusBadRequest)
			return
		}

		verifications, err := repository.LoadVerification(storePath)
		if err != nil {
			http.Error(w, "failed to load store", http.StatusInternalServerError)
			return
		}

		found := false
		out := verifications[:0]
		for _, v := range verifications {
			if v.Hash == hash {
				found = true
				continue
			}
			out = append(out, v)
		}

		if found {
			if err := repository.SaveVerifications(storePath, out); err != nil {
				http.Error(w, "failed to update store", http.StatusInternalServerError)
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]bool{
			"verified": found,
		})
	}
}
