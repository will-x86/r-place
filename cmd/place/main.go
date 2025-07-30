package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	"github.com/will-x86/r-place/pkg/router"
	"github.com/will-x86/r-place/pkg/vk"
)

var XMAX int
var YMAX int

func init() {
	vk.InitialiseValkey()
	var err error
	XMAX, err = strconv.Atoi(os.Getenv("X_MAX"))
	if err != nil {
		panic(err)
	}
	YMAX, err = strconv.Atoi(os.Getenv("Y_MAX"))

	if err != nil {
		panic(err)
	}

}
func main() {
	for k,v := range 
	defer vk.CloseValkey()
	/*
		Important for "mission critical" golang programs, if the program panics, start again.

		Could lead to missing panics, but no program vs one that crashes is not my debate
	*/
	defer func() {
		if r := recover(); r != nil {

			log.Println("Recovered. Error:\n", r)
		}
	}()
	log.Printf("App enviroment: %s", os.Getenv("APP_ENV"))
	// get router that has all requests in
	r := router.NewRouter()
	if err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), r); err != nil {
		log.Println("Fatal ListenAndServe error: %w", err)
	}

}
