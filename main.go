package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("hello world")

	godotenv.Load(".env")

	portString := os.Getenv("PORT")

	if portString==""{
		log.Fatal("Port is not found in env")
	}
	
	router :=chi.NewRouter();

	srv :=&http.Server{
		Handler:router,
		Addr: ":" +portString,
	}

	err := srv.ListenAndServe()

	if err!=nil{
		log.Fatal(err)
	}
	fmt.Println("Port", portString)

	
}