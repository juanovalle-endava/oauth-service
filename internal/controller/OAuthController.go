package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"net/http"
)

type OAuthController interface {
	Get(ctx *gin.Context)
}

type oAuthController struct {
}

func NewOAuthController() OAuthController {
	return &oAuthController{}
}

func (controller *oAuthController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello user with id: %s", id),
	})
}

var Module = fx.Provide(NewOAuthController)
