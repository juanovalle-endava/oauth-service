package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/juanovalle-endava/oauth-service/internal/service"
	"github.com/juanovalle-endava/oauth-service/internal/token"
	"go.uber.org/fx"
	"log"
	"net/http"
)

type VerifyParams struct {
	Token string `json:"token"`
}

type OAuthController interface {
	ListTokens(ctx *gin.Context)
	CreateToken(ctx *gin.Context)
	VerifyToken(ctx *gin.Context)
}

type oAuthController struct {
	token   token.Token
	service service.OAuthService
}

func NewOAuthController(token token.Token, authService service.OAuthService) OAuthController {
	return &oAuthController{
		token:   token,
		service: authService,
	}
}

func (c *oAuthController) ListTokens(ctx *gin.Context) {

	tokens, err := c.service.ListTokens()
	if err != nil {
		log.Fatalln(err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"tokens": tokens,
	})
}

func (c *oAuthController) CreateToken(ctx *gin.Context) {
	clientId, clientSecret, ok := ctx.Request.BasicAuth()
	if !ok {
		log.Fatalln("there was an error getting the client information")
		ctx.AbortWithStatus(500)
		return
	}

	if clientId != "endava" || clientSecret != "secretpass" {
		log.Fatalln("invalid client or password")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid client or password"})
		return
	}

	token, err := c.token.CreateToken(clientId)
	if err != nil {
		log.Fatalln(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.service.SaveToken(token)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (c *oAuthController) VerifyToken(ctx *gin.Context) {
	var params VerifyParams
	if err := ctx.BindJSON(&params); err != nil {
		log.Fatalln(err)
		ctx.AbortWithStatus(500)
		return
	}

	err := c.token.VerifyToken(params.Token)
	if err != nil {
		log.Fatalln(err)
		ctx.AbortWithStatus(500)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{})

}

var Module = fx.Provide(NewOAuthController)
