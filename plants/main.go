package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"plants/add"
	"plants/delete"
	"plants/edit"
	"plants/list"
	"plants/server"
	"plants/storage"
)

var db = storage.ConnectDatabase()

func main() {
	r := storage.NewRepository(db)
	ad := add.NewAddService(r)
	ls := list.NewListService(r)
	ed := edit.NewEditService(ls, r)
	dl := delete.NewDeleteService(r)

	go gracefulShutdown()

	err := server.Start(ad, ls, ed, dl)
	if err != nil {
		log.Panic(err)
	}
}

func gracefulShutdown() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := storage.DisconnectDatabase(ctx, db)
	if err != nil {
		fmt.Print(err)
	}

	err = server.Stop(ctx)
	if err != nil {
		fmt.Print(err)
	}
}
