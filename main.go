package main

import (
    routers "impl-api/router"
    "net/http"
)

// @title mux swagger
// @version 2.0
// @description mux swagger

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// schemes http
func main() {

    router := routers.NewRouter()
    http.ListenAndServe(":3000", router)

}
