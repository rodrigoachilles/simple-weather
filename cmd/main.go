package main

import (
	"fmt"
	"github.com/rodrigoachilles/simple-weather/internal/infra/web"
	"github.com/rodrigoachilles/simple-weather/internal/usecase"
	"github.com/rodrigoachilles/simple-weather/pkg/log"
	"net/http"
)

func main() {
	port := ":8080"
	log.Logger.Info().Msg(fmt.Sprintf("Starting server on port '%s'...", port[1:]))

	mux := http.NewServeMux()
	localeFinder := usecase.NewLocaleFinder(http.DefaultClient)
	weatherFinder := usecase.NewWeatherFinder(http.DefaultClient)
	mux.HandleFunc("GET /{cep}", web.NewLocaleHandler(localeFinder, weatherFinder).Handle)

	log.Logger.Fatal().Err(http.ListenAndServe(port, mux))
}
