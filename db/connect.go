package db

import (
	"fmt"
	"grpc-ent/configs"
	"grpc-ent/ent"

	_ "github.com/go-sql-driver/mysql"
)

func Open() (*ent.Client, error) {
	configs.LoadConfigs()
	dsn := fmt.Sprint(configs.Values.User,
		":",
		configs.Values.Password,
		"@tcp(",
		configs.Values.Host,
		":",
		configs.Values.Port,
		")/",
		configs.Values.DBName, "?parseTime=True")

	client, err := ent.Open("mysql", dsn)
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}
	return client, nil
}
