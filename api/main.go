package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-mongo-api/config"
	"go-mongo-api/controllers"
	"go-mongo-api/database"
	"go-mongo-api/services"
	"log"
	"net/http"
)

func main() {
	postService := &services.PostService{
		Col: database.GetDB().Collection(config.GetConfig().Mongo.Collections.Posts),
		Ctx: database.GetGlobalContext(),
	}

	r := mux.NewRouter()

	r.HandleFunc("/posts", controllers.FindPosts(postService)).Methods("GET")
	r.HandleFunc("/posts/{id}", controllers.GetPost(postService)).Methods("GET")
	r.HandleFunc("/posts", controllers.CreatePost(postService)).Methods("POST")
	r.HandleFunc("/posts/{id}", controllers.PatchPost(postService)).Methods("PATCH")
	r.HandleFunc("/posts/{id}", controllers.DeletePost(postService)).Methods("DELETE")

	log.Println(fmt.Sprintf("App's running on port: 8080"))
	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}
