package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"penghui.com/controllers"
)

var router = httprouter.New()

func init() {
	router.NotFound = http.FileServer(http.Dir("public"))
	router.GET("/", new(controllers.HomeController).TechIndex)
	router.GET("/news/:id", new(controllers.HomeController).NewsIndex)
	router.POST("/my/", new(controllers.HomeController).MyLog)
}

func Register() *httprouter.Router {
	return router
}
