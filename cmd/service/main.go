package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/Elunded/lab3/internal/processor"
)

func main() {
	go func() {
		log.Println("Pprof server started on :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	log.Println("Image Metadata Processor started...")
	processor.RunWorkerPool(5)
}
