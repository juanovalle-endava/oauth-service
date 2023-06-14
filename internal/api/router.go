package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanovalle-endava/oauth-service/internal/controller"
	"github.com/juanovalle-endava/oauth-service/internal/utils"
	"go.uber.org/fx"
	"log"
)

type Api struct {
	OAuthController controller.OAuthController
}

var accounts = gin.Accounts{
	"endava": "secretpass",
	"user":   "root",
}

func (api *Api) SetupRoutes() {
	r := gin.Default()

	apiRoute := r.Group("/api")
	apiRoute.GET("/signing-keys", api.OAuthController.ListSigningKeys)
	apiRoute.POST("/token", utils.AuthMiddleware(accounts), api.OAuthController.CreateToken)
	apiRoute.GET("/introspect", utils.VerifyMiddleware(), api.OAuthController.VerifyToken)

	if err := r.Run(":8080"); err != nil {
		log.Fatalln("Failed to listen and serve on port 8080: " + err.Error())
		panic(err)
	}

}

func NewOAuthRoutes(oAuthController controller.OAuthController) Api {
	return Api{OAuthController: oAuthController}
}

var Module = fx.Provide(NewOAuthRoutes)
