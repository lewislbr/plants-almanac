package main

import (
	"log"

	"plants/src/add"
	"plants/src/api"
	"plants/src/delete"
	"plants/src/edit"
	"plants/src/list"
	"plants/src/storage"
)

func main() {
	db := storage.ConnectDatabase()
	repository := storage.NewRepository(db)
	addService := add.NewAddService(repository)
	listService := list.NewListService(repository)
	editService := edit.NewEditService(listService, repository)
	deleteService := delete.NewDeleteService(repository)

	if err := api.Start(addService, listService, editService, deleteService); err != nil {
		log.Fatalln(err)
	}
}
