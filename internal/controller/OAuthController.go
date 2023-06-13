package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juanovalle-endava/oauth-service/internal/token"
	"go.uber.org/fx"
	"net/http"
)

type OAuthController interface {
	Get(ctx *gin.Context)
	CreateToken(ctx *gin.Context)
}

type oAuthController struct {
	token token.Token
}

func NewOAuthController(token token.Token) OAuthController {
	return &oAuthController{
		token: token,
	}
}

func (controller *oAuthController) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Hello user with id: %s", id),
	})
}

func (o *oAuthController) CreateToken(ctx *gin.Context) {
	clientId, clientSecret, ok := ctx.Request.BasicAuth()
	if !ok {
		ctx.AbortWithStatus(500)
		return
	}

	if clientId != "endava" || clientSecret != "secretpass" {
		ctx.AbortWithStatus(500)
		return
	}

	token, err := o.token.CreateToken(clientId)
	if err != nil {
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

var Module = fx.Provide(NewOAuthController)
