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
	r := storage.NewRepository(db)
	ad := add.NewAddService(r)
	ls := list.NewListService(r)
	ed := edit.NewEditService(ls, r)
	dl := delete.NewDeleteService(r)

	if err := api.Start(ad, ls, ed, dl); err != nil {
		log.Fatalln(err)
	}
}
