package router

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"penghui.com/config"
	"penghui.com/controllers"
	"time"
)

var router = httprouter.New()

func init() {
	router.NotFound = http.FileServer(http.Dir("public"))
	router.GET("/", basicAuth(new(controllers.HomeController).TechIndex))
	router.GET("/news/:id", basicAuth(new(controllers.HomeController).NewsIndex))
	router.POST("/my", new(controllers.HomeController).MyLog)
	router.GET("/login", new(controllers.HomeController).Login)
	router.GET("/redis/:id", new(controllers.HomeController).MyRedis)
}

func Register() *httprouter.Router {
	return router
}

func basicAuth(h func(w http.ResponseWriter, r *http.Request, ps httprouter.Params)) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenStr := r.Header.Get("token")

		if tokenStr == "" || len(tokenStr) <= 0 {
			tokenStr = "ayJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhZG1pbiI6dHJ1ZSwiZXhwIjoxNTQ2MzA3OTM0LCJuYW1lIjoiSm9uIFNub3cifQ.7aQb5BXrLaEgkXOYO7ykqatrTfNyD1s2W_leNNVUkd8"
		}

		hmacSampleSecret := []byte(config.Conf.Token.Token)

		token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			}
			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println("this is name:", claims["name"])
			r.Header.Add("starttime", time.Now().Format(""))
			ps = append(ps, httprouter.Param{"names", claims["name"].(string)})
			h(w, r, ps)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
