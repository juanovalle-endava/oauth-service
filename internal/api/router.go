package api

import (
	"github.com/gin-gonic/gin"
	"github.com/juanovalle-endava/oauth-service/internal/controller"
	"go.uber.org/fx"
)

type Api struct {
	OAuthController controller.OAuthController
}

func (api *Api) SetupRoutes() {
	r := gin.Default()

	apiRoute := r.Group("/api")
	apiRoute.GET("/:id", api.OAuthController.Get)
	apiRoute.POST("/token", api.OAuthController.CreateToken)

	r.Run(":8080")

}

func NewOAuthRoutes(oAuthController controller.OAuthController) Api {
	return Api{OAuthController: oAuthController}
}

var Module = fx.Provide(NewOAuthRoutes)
