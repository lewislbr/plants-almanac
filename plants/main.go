package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"plants/add"
	"plants/delete"
	"plants/edit"
	"plants/list"
	"plants/server"
	"plants/storage"
)

var db, err = storage.ConnectDatabase()

func main() {
	defer func() {
		r := recover()
		if r != nil {
			cleanUp()
			debug.PrintStack()
			os.Exit(1)
		}
	}()

	if err != nil {
		log.Panic(err)
	}

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

func cleanUp() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if db != nil {
		err := storage.DisconnectDatabase(ctx, db)
		if err != nil {
			fmt.Println(err)
		}
	}

	err = server.Stop(ctx)
	if err != nil {
		fmt.Println(err)
	}
}

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	cleanUp()
}
