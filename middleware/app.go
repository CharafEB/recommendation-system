package middlewares

import (
	"github.com/recommendation-system/model"
)

type Application struct {
	Address string
	Storge  model.Store
}
