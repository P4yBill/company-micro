package main

import (
	"company-micro/api/router"
	"company-micro/config"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func main() {

	config := config.App()
	env := config.Env

	db := config.Mongo.Database(env.DBName)
	defer config.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second
	fmt.Println(env)
	mux := router.GetRouter(env, timeout, db)

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%s", env.ServerAddress, env.Port),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			fmt.Printf("error running http server: %s\n", err)
		}
	}
}
