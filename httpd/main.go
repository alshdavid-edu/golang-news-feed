package main

import (
	"fmt"
	"net/http"
	"newsfeeder/httpd/environment"
	handlersIndexGet "newsfeeder/httpd/handlers/index_get"
	handlersNewsfeedGet "newsfeeder/httpd/handlers/newsfeed_get"
	handlersNewsfeedPost "newsfeeder/httpd/handlers/newsfeed_post"
	servicecontainer "newsfeeder/httpd/service-container"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	env := environment.NewEnvironment()
	services := servicecontainer.New()
	app := chi.NewRouter()

	app.Use(middleware.Logger)

	app.Get("/", handlersIndexGet.Handler())
	app.Get("/newsfeed", handlersNewsfeedGet.Handler(services))
	app.Post("/newsfeed", handlersNewsfeedPost.Handler(services))

	fmt.Println(fmt.Sprintf("Serving on http://localhost:%s", env.Port))
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%s", env.Port), app)
}
