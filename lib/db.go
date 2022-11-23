package lib

import (
	"context"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Database struct {
	Driver neo4j.DriverWithContext
}

func NewDatabase(ctx context.Context, env Env, logger Logger) Database {
	username := env.DBUsername
	password := env.DBPassword
	host := env.DBHost
	port := env.DBPort
	uri := "neo4j://" + host + ":" + port
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		logger.Fatal(err)
	}
	defer driver.Close(ctx)

	return Database{Driver: driver}
}
