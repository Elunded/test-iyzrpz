package main

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("service", "image-metadata-processor").Logger()

	logger.Info().Msg("Image Metadata Processor started...")
	time.Sleep(1 * time.Hour)
}
