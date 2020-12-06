package main

import "fmt"

import ("routeauth/src/security"
		"github.com/go-chi/chi"
		"github.com/go-chi/chi/middleware"
		"routeauth/src/resource"
		"time"
		"net/http")



func main() {

	r := chi.NewRouter()

	r.Use(middleware.Timeout(60 * time.Second))

	resource.InitAuthResource(r);

	fmt.Printf("%+v\n", r)

	http.ListenAndServe(":8080", r)
	
	fmt.Println(security.CreateJWT("secret"))
}