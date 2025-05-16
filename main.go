package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	logger2 "gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

// order
func main() {
	order()
	fb()
}

func order() {
	cfg := gen.Config{
		OutPath: "./order",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	}

	cfg.WithDataTypeMap(map[string]func(columnType gorm.ColumnType) (dataType string){
		"bigint": func(columnType gorm.ColumnType) string {
			fmt.Println(columnType.ColumnType())
			val, _ := columnType.ColumnType()
			if strings.Contains(val, "bigint unsigned") {
				return "uint64"
			}
			return "int64"
		},
	})
	g := gen.NewGenerator(cfg)

	newLogger := logger2.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 打印到标准输出
		logger2.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger2.Info, // 这里设置为 Info 才能看到 SQL
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	gormdb, _ := gorm.Open(mysql.Open("dev_user:hK9%bZ8.eC2%@tcp(47.250.115.85:4000)/orders?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
		Logger: newLogger,
	})
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
		"bigint": func(columnType gorm.ColumnType) string {
			fmt.Println(columnType.ColumnType())
			val, _ := columnType.ColumnType()
			if strings.Contains(val, "bigint unsigned") {
				return "uint64"
			}
			return "int64"
		},
	})
	g := gen.NewGenerator(cfg)

	newLogger := logger2.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 打印到标准输出
		logger2.Config{
			SlowThreshold:             time.Second,  // 慢 SQL 阈值
			LogLevel:                  logger2.Info, // 这里设置为 Info 才能看到 SQL
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	gormdb, _ := gorm.Open(mysql.Open("dev_user:hK9%bZ8.eC2%@tcp(47.250.115.85:4000)/fb?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{
		Logger: newLogger,
	})
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
