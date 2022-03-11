package routes

import (
	controller "impl-api/controller"
	service "impl-api/service"
	"net/http"
	"strings"

	_ "impl-api/docs"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

var routes []Route

func init() {
	/* user  */
	//GET
	register("GET", "/userList/{page:[0-9]+}", controller.UserList, authMiddleware)
	register("GET", "/userSearch/{key}", controller.UserSearch, authMiddleware)
	register("GET", "/userDetail", controller.UserDetail, authMiddleware)
	//POST
	register("POST", "/userSignup", controller.UserSignup, nil)
	register("POST", "/userSignin", controller.UserSignin, nil)
	//DELETE
	register("DELETE", "/userDel", controller.UserDel, authMiddleware)
	//PUT
	register("PUT", "/userUpdate", controller.UserUpdate, authMiddleware)
	//PATCH
	register("PATCH", "/userUpdateFullname", controller.UserUpdateFullname, authMiddleware)
}
func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	apiRouter := r.PathPrefix("/api/").Subrouter()
	apiRouter.Use(authMiddleware)
	for _, route := range routes {
		if route.Middleware != nil {
			apiRouter.Methods(route.Method).
				Path(route.Pattern).
				Handler(route.Handler)
		} else {
			r.Methods(route.Method).Path("/api" + route.Pattern).Handler(route.Handler)
		}
	}

	return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
	routes = append(routes, Route{"OPTIONS", pattern, controller.OptionRespons, nil})
}
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Do stuff here

		token := r.Header.Get("Authorization")
		token = strings.Replace(token, "Bearer ", "", -1)
		acct, error := service.JWTcheck(token)
		if r.Method == "OPTIONS" {
			next.ServeHTTP(w, r)
		} else if error != nil {
			http.Error(w, error.Error(), http.StatusBadRequest)
		} else {
			context.Set(r, "acct", acct)
			next.ServeHTTP(w, r)
		}
	})
}
