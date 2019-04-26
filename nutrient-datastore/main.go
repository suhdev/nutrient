package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gocraft/dbr"
	"github.com/micro/go-micro"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/suhdev/nutrient-datastore/db"
	"github.com/suhdev/nutrient-datastore/handler"
	"github.com/suhdev/nutrient/proton"
)

func main() {
	service := micro.NewService(
		micro.Name("DS"),
	)
	service.Init()
	conn, err := dbr.Open("mysql", "root:1234@/recipes?parseTime=true", nil)
	if err != nil {
		log.Fatal(err)
	}
	conn.SetMaxOpenConns(20)
	migrate.Exec(conn.DB, "mysql", db.Migrations, migrate.Down)
	migrate.Exec(conn.DB, "mysql", db.Migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
	proton.RegisterDataStoreHandler(service.Server(),
		handler.NewDataStore(conn))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
