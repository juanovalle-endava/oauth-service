package main

import (
	"github.com/juanovalle-endava/oauth-service/internal/bootstrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(bootstrap.Module).Run()
}
