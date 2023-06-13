package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/juanovalle-endava/oauth-service/internal/token"
	"go.uber.org/fx"
	"log"
	"net/http"
)

type VerifyParams struct {
	Token string `json:"token"`
}

type OAuthController interface {
	Get(ctx *gin.Context)
	CreateToken(ctx *gin.Context)
	VerifyToken(ctx *gin.Context)
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
		log.Fatalln("there was an error getting the client information")
		ctx.AbortWithStatus(500)
		return
	}

	if clientId != "endava" || clientSecret != "secretpass" {
		log.Fatalln("invalid client or password")
		ctx.AbortWithStatus(500)
		return
	}

	token, err := o.token.CreateToken(clientId)
	if err != nil {
		log.Fatalln(err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (o *oAuthController) VerifyToken(ctx *gin.Context) {
	var params VerifyParams
	if err := ctx.BindJSON(&params); err != nil {
		log.Fatalln(err)
		ctx.AbortWithStatus(500)
		return
	}

	err := o.token.VerifyToken(params.Token)
	if err != nil {
		log.Fatalln(err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})

}

var Module = fx.Provide(NewOAuthController)
