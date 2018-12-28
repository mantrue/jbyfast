package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"penghui.com/controllers"
)

var router = httprouter.New()

func init() {
	router.NotFound = http.FileServer(http.Dir("public"))
	router.GET("/", basicAuth(new(controllers.HomeController).TechIndex))
	router.GET("/news/:id", basicAuth(new(controllers.HomeController).NewsIndex))
	router.POST("/my/", new(controllers.HomeController).MyLog)
}

func Register() *httprouter.Router {
	return router
}

func basicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.URL.Path == "/" {
			h(w, r, ps)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
