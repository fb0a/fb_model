package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

// order
func main() {
	//order()
	fb()
}

func order() {
	cfg := gen.Config{
		OutPath: "./order",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	}

	cfg.WithDataTypeMap(map[string]func(columnType gorm.ColumnType) (dataType string){
		//"id": func(columnType gorm.ColumnType) string { return "int64" },
	})
	g := gen.NewGenerator(cfg)

	gormdb, _ := gorm.Open(mysql.Open("dev_user:hK9%bZ8.eC2%@tcp(47.250.115.85:4000)/orders?charset=utf8&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.ApplyBasic(model.User{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}

func fb() {
	cfg := gen.Config{
		OutPath: "./fb",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	}

	cfg.WithDataTypeMap(map[string]func(columnType gorm.ColumnType) (dataType string){
		//"id": func(columnType gorm.ColumnType) string { return "int64" },
	})
	g := gen.NewGenerator(cfg)

	gormdb, _ := gorm.Open(mysql.Open("dev_user:hK9%bZ8.eC2%@tcp(47.250.115.85:4000)/fb?charset=utf8&parseTime=True&loc=Local"))
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	//g.ApplyBasic(model.User{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	//g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
