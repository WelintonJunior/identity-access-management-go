package auth

import (
	"log"
	"net/http"

	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
)

func SetupSSO() {
	go func() {
		manager := manage.NewDefaultManager()
		manager.MustTokenStorage(store.NewMemoryTokenStore())

		clientStore := store.NewClientStore()
		clientStore.Set("000000", &models.Client{
			ID:     "000000",
			Secret: "999999",
			Domain: "http://localhost",
		})
		manager.MapClientStorage(clientStore)

		srv := server.NewDefaultServer(manager)
		srv.SetAllowGetAccessRequest(true)
		srv.SetClientInfoHandler(server.ClientFormHandler)

		srv.UserAuthorizationHandler = func(w http.ResponseWriter, r *http.Request) (userID string, err error) {
			return "000000", nil
		}

		srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
			log.Println("Internal Error:", err.Error())
			return
		})

		srv.SetResponseErrorHandler(func(re *errors.Response) {
			log.Println("Response Error:", re.Error.Error())
		})

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/plain")
			w.Write([]byte("Welcome to the OAuth2 SSO server!"))
		})

		http.HandleFunc("/authorize", func(w http.ResponseWriter, r *http.Request) {
			err := srv.HandleAuthorizeRequest(w, r)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
		})

		http.HandleFunc("/token", func(w http.ResponseWriter, r *http.Request) {
			srv.HandleTokenRequest(w, r)
		})

		log.Fatal(http.ListenAndServe(":9096", nil))
	}()
}
