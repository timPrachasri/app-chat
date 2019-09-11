package environments

import (
	"fmt"
	"syscall"
)

var (
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	FirebaseFilename string
)

// Init Initialize env variables
func Init() {
	PostgresHost = requireEnv("POSTGRES_HOST")
	PostgresPort = requireEnv("POSTGRES_PORT")
	PostgresUser = requireEnv("POSTGRES_USER")
	PostgresPassword = requireEnv("POSTGRES_PASSWORD")
	PostgresDB = requireEnv("POSTGRES_DB")
	FirebaseFilename = requireEnv("FIREBASE_FILENAME")
}

func requireEnv(envName string) string {
	value, found := syscall.Getenv(envName)

	if !found {
		panic(fmt.Sprintf("%s env is required", envName))
	}

	return value
}
