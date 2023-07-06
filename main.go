package main

import (
	"fmt"
	"github.com/jessehorne/newsengine/db"
	"github.com/jessehorne/newsengine/routes"
	"github.com/jessehorne/newsengine/util"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// Init Config
	if err := util.InitConf(); err != nil {
		log.Fatal(err)
	}

	// Init DB
	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}

	// Init Validator
	util.InitValidator()

	// Routes
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/articles", func(r chi.Router) {
		r.Route("/{articleID}", func(r chi.Router) {
			r.Get("/", routes.GetArticleByID)
			r.Put("/", routes.UpdateArticleByID)
			r.Delete("/", routes.DeleteArticleByID)
		})

		r.Get("/", routes.GetArticles)
		r.Post("/", routes.CreateArticle)
	})

	version := os.Getenv("APP_VERSION")
	port := os.Getenv("APP_PORT")

	fmt.Printf("newsengine v%v running on port %v\n", version, port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), r)
}
