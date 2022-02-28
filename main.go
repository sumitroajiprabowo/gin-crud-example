package main

import (
	"github.com/sumitroajiprabowo/gin-crud-example/app"
	"github.com/sumitroajiprabowo/gin-crud-example/model/entity"
	"github.com/sumitroajiprabowo/gin-crud-example/routes"
)

func main() {

	db := app.SetupDB()

	// Migrate the schema
	db.AutoMigrate(&entity.Package{})

	r := routes.SetupRoutes(db)
	r.Run()
}
