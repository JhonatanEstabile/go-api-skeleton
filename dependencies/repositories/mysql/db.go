package mysql

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

type Repository struct {
	Client   *sqlx.DB
	log      *log.Logger
	user     string
	password string
	host     string
	name     string
}

func NewMysql(log *log.Logger, user, password, host, name string) *Repository {
	return &Repository{
		log:      log,
		user:     user,
		password: password,
		host:     host,
		name:     name,
	}
}

func (m *Repository) Connect() {
	var err error
	m.Client, err = sqlx.Connect("mysql", m.getConnectionUrl())
	if err != nil {
		//IMPLEMENTAR LOG
		log.Fatalf("Erro ao conectar no banco de dados: %s", err.Error())
	}

	m.Client.SetConnMaxIdleTime(time.Minute * 5)
	m.Client.SetMaxIdleConns(5)
	m.Client.SetMaxOpenConns(10)

	m.runMigrations()
}

func (m *Repository) Close() {
	m.Client.Close()
}

func (m *Repository) GetClient() *sqlx.DB {
	return m.Client
}

func (m *Repository) runMigrations() {
	file, err := os.Open("./queries/migrations/migrations.json")
	if err != nil {
		log.Fatalf("Erro ao ler json de migrações: %v", err)
	}
	//defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)
	var tables []string

	err = json.Unmarshal(byteValue, &tables)
	if err != nil {
		//IMPLEMENTAR LOG
		log.Fatalf("Erro ao fazer o parse do json de migrações: %v", err)
	}

	dot, err := dotsql.LoadFromFile("./queries/migrations/tables.sql")
	if err != nil {
		//IMPLEMENTAR LOG
		log.Fatalf("Erro ao fazer o parse do json de migrações: %v", err)
	}

	for i := 0; i < len(tables); i++ {
		_, err = dot.Exec(
			m.Client,
			tables[i],
		)

		if err != nil {
			//IMPLEMENTAR LOG
			log.Fatalf("Erro criar tabela: %s, Erro: %v", tables[i], err)
		}
	}
}

func (m *Repository) getConnectionUrl() string {
	connectionUrl := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s",
		m.user,
		m.password,
		m.host,
		m.name,
	)

	return connectionUrl
}
