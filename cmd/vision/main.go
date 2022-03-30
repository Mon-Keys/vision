package main

import (
	app "github.com/perlinleo/vision/internal/app"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Error().Msgf(app.Start().Error())
}
