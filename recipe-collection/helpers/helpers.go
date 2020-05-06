package helpers

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/joho/godotenv"
)

//LoadENV loads the appropriate env variables with respect
// to the current environment
func LoadENV() {
	var envFile string

	switch env := os.Getenv("RECIPE_ENV"); env {
	case "":
		log.Fatal("Application environment {RECIPE_ENV} not set")
	case "production":
		envFile = ".env.production"
	case "development":
		envFile = ".env.development"
	case "test":
		envFile = ".env.test"
	default:
		log.Fatal("Invalid Application environment value")
	}

	err := godotenv.Load(fmt.Sprintf("%s/%s", AppBasePath(), envFile))

	if err != nil {
		log.Fatalf("Could not load env file: %v, %v", err.Error(), envFile)
	}
}

//AppBasePath returns the base path of the application
func AppBasePath() string {
	_, filename, _, _ := runtime.Caller(0)

	return filepath.Dir(filepath.Dir(filename))
}
