package environment

import "os"

type Environment struct {
	Port string
}

func NewEnvironment() *Environment {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Environment{
		Port: port,
	}
}
