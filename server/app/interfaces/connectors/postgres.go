package connectors

import (
	"fmt"
	"log"
	"math"

	"github.com/golang-migrate/migrate"
	"github.com/jinzhu/gorm"
	"github.com/timPrachasri/app-chat/server/app/environments"
)

var (
	// postgresDBInstance is an instance for of postgres db in gorm format
	postgresDBInstance *gorm.DB
)

// ConnectPostgresDB is a connector function for connecting postgres
func ConnectPostgresDB() *gorm.DB {
	return connect(
		environments.PostgresHost,
		environments.PostgresPort,
		environments.PostgresUser,
		environments.PostgresPassword,
		environments.PostgresDB,
	)
}

// RunMigration will run migration script in postgres
func RunMigration() {
	runMigration(
		environments.PostgresUser,
		environments.PostgresPassword,
		environments.PostgresHost,
		environments.PostgresPort,
		environments.PostgresDB,
		"migrations",
	)
}

func connect(host, port, user, password, dbName string) *gorm.DB {
	var err error
	connection := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName)
	mut.Lock()
	defer mut.Unlock()
	if postgresDBInstance == nil {
		postgresDBInstance, err = gorm.Open("postgres", connection)
		if err != nil {
			log.Println("ERROR :(((((")
			panic(err.Error())
		}
		log.Println("database is connected")
	}

	return postgresDBInstance
}

func runMigration(user, password, host, port, dbName, migrationFolder string) {
	databaseURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbName)

	migrationEngine, err := migrate.New("file://../app/"+migrationFolder, databaseURL)
	if err != nil {
		panic(err)
	}
	_ = migrationEngine.Steps(math.MaxInt32)
}
