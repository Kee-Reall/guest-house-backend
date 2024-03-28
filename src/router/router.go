package router

import (
	"example.com/m/v2/src/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"runtime/debug"
)

func setRoutes(router *mux.Router) {
	controller := controllers.CreateController()
	router.Use(panicRecovery)
	router.Use(setJson)
	router.HandleFunc("/like", controller.HelloWorld).Methods("GET")
	router.HandleFunc("/", controller.PostHandle).Methods("POST")
}

func setJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		const key string = "Content-type"
		const appType string = "application/json"
		writer.Header().Set(key, appType)
		next.ServeHTTP(writer, request)
	})

}

func panicRecovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Println(string(debug.Stack()))
			}
		}()
		next.ServeHTTP(w, req)
	})
}

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	setRoutes(router)
	return router
}
