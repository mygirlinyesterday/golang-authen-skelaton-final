package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/playground/api/config"
	"github.com/playground/api/router"
	"github.com/playground/auto"
)

func Run() {
	config.Load()
	auto.Load()
	fmt.Printf("\n\tListening [::]:%d\n", config.PORT)
	listen(config.PORT)
}

func listen(port int) {
	r := router.New()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
