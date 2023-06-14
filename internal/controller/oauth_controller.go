package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/juanovalle-endava/oauth-service/internal/service"
	log "github.com/sirupsen/logrus"
	"go.uber.org/fx"
	"net/http"
)

type OAuthController interface {
	ListSigningKeys(ctx *gin.Context)
	CreateToken(ctx *gin.Context)
	VerifyToken(ctx *gin.Context)
}

type oAuthController struct {
	service service.OAuthService
}

func NewOAuthController(authService service.OAuthService) OAuthController {
	return &oAuthController{
		service: authService,
	}
}

func (c *oAuthController) ListSigningKeys(ctx *gin.Context) {

	keyResponse, err := c.service.ListSigningKeys()
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": keyResponse,
	})
}

func (c *oAuthController) CreateToken(ctx *gin.Context) {
	username := ctx.GetString("username")

	token, err := c.service.CreateToken(username)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}

func (c *oAuthController) VerifyToken(ctx *gin.Context) {

	tokenString := ctx.GetString("token")

	err := c.service.VerifyToken(tokenString)
	if err != nil {
		log.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)

}

var Module = fx.Provide(NewOAuthController)
