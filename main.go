package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5"
)

type Dbcreds struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	DbName   string `json:"dbname"`
	Password string `json:"password"`
}

func main() {
	dbcreds := readDbcreds("dbcreds.json")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbcreds.Host, dbcreds.Port, dbcreds.User, dbcreds.Password, dbcreds.DbName)

	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
}

func normalize(phone string) string {
	digits := "1234567890"
	var ret string
	for _, ch := range phone {
		if strings.Contains(digits, string(ch)) {
			ret += string(ch)
		}
	}
	return ret
}

func readDbcreds(fileName string) Dbcreds {
	configFile, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading %v:\n", fileName)
		panic(err)
	}
	var dbcreds Dbcreds
	if err := json.Unmarshal(configFile, &dbcreds); err != nil {
		fmt.Printf("Error unmarshaling %v:\n", fileName)
		panic(err)
	}
	return dbcreds
}
