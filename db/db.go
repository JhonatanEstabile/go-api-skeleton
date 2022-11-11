package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/qustavo/dotsql"
)

var client *sqlx.DB

func Connect() {
	var err error
	client, err = sqlx.Connect("mysql", getConnectionUrl())
	if err != nil {
		log.Fatalf("Erro ao conectar no banco de dados: %s", err.Error())
	}

	client.SetConnMaxIdleTime(time.Minute * 5)
	client.SetMaxIdleConns(5)
	client.SetMaxOpenConns(10)

	runMigrations()
}

func Close() {
	client.Close()
}

func GetClient() *sqlx.DB {
	return client
}

func runMigrations() {
	jsonFile, err := os.Open("./queries/migrations/migrations.json")
	if err != nil {
		log.Fatalf("Erro ao ler json de migrações: %v", err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var tables []string

	err = json.Unmarshal(byteValue, &tables)
	if err != nil {
		log.Fatalf("Erro ao fazer o parse do json de migrações: %v", err)
	}

	dot, err := dotsql.LoadFromFile("./queries/migrations/tables.sql")
	if err != nil {
		log.Fatalf("Erro ao fazer o parse do json de migrações: %v", err)
	}

	for i := 0; i < len(tables); i++ {
		_, err = dot.Exec(
			client,
			tables[i],
		)

		if err != nil {
			log.Fatalf("Erro criar tabela: %s, Erro: %v", tables[i], err)
		}
	}
}

func getConnectionUrl() string {
	connectionUrl := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)

	return connectionUrl
}
