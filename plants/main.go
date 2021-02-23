package main

import (
	"log"

	"plants/add"
	"plants/api"
	"plants/delete"
	"plants/edit"
	"plants/list"
	"plants/storage"
)

func main() {
	db := storage.ConnectDatabase()
	r := storage.NewRepository(db)
	ad := add.NewAddService(r)
	ls := list.NewListService(r)
	ed := edit.NewEditService(ls, r)
	dl := delete.NewDeleteService(r)

	if err := api.Start(ad, ls, ed, dl); err != nil {
		log.Fatalln(err)
	}
}
