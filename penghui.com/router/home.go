package router

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"penghui.com/config"
	"penghui.com/controllers"
)

var router = httprouter.New()

func init() {
	router.NotFound = http.FileServer(http.Dir("public"))
	router.GET("/", basicAuth(new(controllers.HomeController).TechIndex))
	router.GET("/news/:id", basicAuth(new(controllers.HomeController).NewsIndex))
	router.POST("/my/", new(controllers.HomeController).MyLog)
	router.GET("/login", new(controllers.HomeController).Login)
}

func Register() *httprouter.Router {
	return router
}

func basicAuth(h func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		tokenStr := r.Header.Get("token")
		hmacSampleSecret := []byte(config.Conf.Token.Token)

		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			}
			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("this is name:", claims["name"])
			h(w, r, ps)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
