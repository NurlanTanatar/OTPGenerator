package main

import (
	"HW_8/internal/http"
	inmemory "HW_8/internal/store/mongg"
	"context"
	"log"
)

func main() {
	store := inmemory.Init()

	// cacheUser := cache.NewUserCache("localhost:6379", 0, 30)

	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		log.Println(err)
	}

	srv.WaitForGracefulTermination()

}
