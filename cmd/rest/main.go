package main

import (
	"QRCodeGenerator/internal/http"
	"QRCodeGenerator/internal/store/inmemory"
	"context"
	"log"
)

func main() {
	store := inmemory.Init()

	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		log.Println(err)
	}

	srv.WaitForGracefulTermination()

}
