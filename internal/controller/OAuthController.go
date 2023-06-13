package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.uber.org/fx"
	"net/http"
	"strconv"
)

type OAuthController interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type oAuthController struct {
}

func NewOAuthController() OAuthController {
	return &oAuthController{}
}

func (controller *oAuthController) Get(w http.ResponseWriter, r *http.Request) {
	var (
		id  int64
		err error
	)

	vars := mux.Vars(r)
	if id, err = strconv.ParseInt(vars["id"], 10, 64); err != nil {
		// handle error
		return
	}

	// send response
	var response []byte
	if response, err = json.Marshal(fmt.Sprintf("Hello user with id: %d", id)); err != nil {
		// handle error
		return
	}
	w.WriteHeader(200)
	w.Write(response)
}

var Module = fx.Provide(NewOAuthController)
