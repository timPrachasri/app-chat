package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/timPrachasri/app-chat/server/app/environments"
	"github.com/timPrachasri/app-chat/server/app/interfaces/connectors"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	environments.Init()
	connectors.ConnectPostgresDB()
	connectors.RunMigration()

	go func() {
		app := gin.Default()
		app.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowMethods:    []string{"PUT", "GET", "POST", "DELETE", "PATCH"},
			AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
			ExposeHeaders:   []string{"Content-Length"},
		}))
		app.Run()
	}()

	select {}
}
