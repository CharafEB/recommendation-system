package router

import (
	"github.com/recommendation-system/controller"
	middlewares "github.com/recommendation-system/middleware"
)

type Application struct {
	middlewares.Application
}

type Controller struct {
	HandlerController controller.HandlerController
}

type Control struct {
	Controller      *controller.Application
}
