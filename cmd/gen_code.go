// Package utils Generate `DAO` code from entities using `gen` package.
//
// Run the following command in the root folder:
// go run ./cmd/gen_code.go
// The generated code will be placed in `dao/` folder.
package main

import (
	"DouyinSimpleProject/config"
	"DouyinSimpleProject/entity"

	"gorm.io/gen"
)

func genCode() {
	db := config.LoadDB()

	g := gen.NewGenerator(gen.Config{
		OutPath:        "./dao",
		Mode:           gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldCoverable: true,
	})
	g.UseDB(db)
	g.ApplyBasic(
		&entity.User{}, &entity.Video{}, &entity.Favorite{}, &entity.Comment{},
	)

	g.Execute()

}

func main() {
	genCode()
}
