package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanovalle-endava/oauth-service/internal/controller"
	"go.uber.org/fx"
	"log"
)

type Api struct {
	OAuthController controller.OAuthController
}

func (api *Api) SetupRoutes() {
	r := gin.Default()

	apiRoute := r.Group("/api")
	apiRoute.GET("/token", api.OAuthController.ListTokens)
	apiRoute.POST("/token", api.OAuthController.CreateToken)
	apiRoute.POST("/introspect", api.OAuthController.VerifyToken)

	if err := r.Run(":8080"); err != nil {
		log.Fatalln("Failed to listen and serve on port 8080: " + err.Error())
		panic(err)
	}

}

func NewOAuthRoutes(oAuthController controller.OAuthController) Api {
	return Api{OAuthController: oAuthController}
}

var Module = fx.Provide(NewOAuthRoutes)
