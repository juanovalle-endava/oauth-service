package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/juanovalle-endava/oauth-service/internal/controller"
	"go.uber.org/fx"
	"net/http"
)

type Api struct {
	OAuthController controller.OAuthController
}

func (api *Api) SetupRoutes() {
	r := mux.NewRouter()
	apiRoute := r.PathPrefix("/api").Subrouter().StrictSlash(false)

	apiRoute.HandleFunc("/{id}", api.OAuthController.Get).Methods(http.MethodGet)

	if err := http.ListenAndServe(":"+"8080", r); err != nil {
		fmt.Errorf("Failed to listen and serve on port 8080: " + err.Error())
		panic(err)
	}

}

func NewOAuthRoutes(oAuthController controller.OAuthController) Api {
	return Api{OAuthController: oAuthController}
}

var Module = fx.Provide(NewOAuthRoutes)
