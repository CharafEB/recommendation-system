package controller

import (
	"net/http"

	middlewares "github.com/recommendation-system/middleware"
	corn "github.com/robfig/cron/v3"
)

type Application struct {
	Application middlewares.Application
}

type TrakerCron struct {
	Cron        *corn.Cron
	Application *middlewares.Application
}

type HandlerController struct {
	Error interface {
	RespondWithError(w http.ResponseWriter, code int, message string, err error)
}
}

