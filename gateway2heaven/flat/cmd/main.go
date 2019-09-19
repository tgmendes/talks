package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	account "github.com/tgmendes/gateway2heaven/flat"
	"log"
	"net/http"
	"os"
)

type TokenClient struct {
}

func (tc TokenClient) ServiceToken() (*account.ServiceToken, error) {
	return &account.ServiceToken{
		ExpiresIn: 3600,
		AccessTkn: "communication-breakdown",
	}, nil
}

func main() {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		log.Fatal("client ID or secret not set")
	}

	serviceHost := os.Getenv("SERVICE_HOST")
	if serviceHost == "" {
		log.Fatal("host not set")
	}

	tknCl := TokenClient{}
	accountClient := account.NewClient(serviceHost, clientID, tknCl)
	serviceHandler := ServiceHandler{accountClient}

	router := chi.NewRouter()

	corsCfg := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	router.Use(corsCfg.Handler, middleware.StripSlashes)
	router.Route("/accounts/{email}", func(r chi.Router) {
		r.Get("/status", serviceHandler.AccountStatus)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}

type ServiceHandler struct {
	accountClient *account.Client
}

func (sh ServiceHandler) AccountStatus(w http.ResponseWriter, r *http.Request) {
	email := chi.URLParam(r, "email")

	accStat, err := sh.accountClient.VerificationStatus(email)
	if err != nil {
		// Not checking error type on purpose
		log.Printf("error retrieving verification status, %v", err.Error())
		w.WriteHeader(http.StatusBadGateway)
		return
	}

	b, err := json.Marshal(accStat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
