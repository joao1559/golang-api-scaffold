package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

//Server ...
type Server struct {
	DBConn *sql.DB
	Env    string
}

//factory
var (
	DBConn *sql.DB
)

//InitDb represent a factory of database
func InitDb() {
	a := Server{}
	a.Env = os.Getenv("ENV")
	connectionString := fmt.Sprintf("%s", a.GetDNS())
	var err error

	DBConn, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("[db/init] - Erro ao tentar abrir conexão (%s). Erro: %s", a.Env, err.Error())
	}
	DBConn.SetConnMaxLifetime(time.Minute * 5)
	DBConn.SetMaxIdleConns(0)
	DBConn.SetMaxOpenConns(15)
	DBConn.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Printf("[db/init] - Erro ao tentar abrir conexão (%s). Erro: %s", a.Env, err.Error())
	}
}

//GetDNS representa a recuperação do acesso ao banco
func (a *Server) GetDNS() string {
	var (
		dbUser     string
		dbPassword string
		dbName     string
		dbHost     string
		dbPort     int
	)
	file, err := ioutil.ReadFile("./config/env.json")
	if err == nil {
		jsonMap := make(map[string]interface{})
		json.Unmarshal(file, &jsonMap)

		env := os.Getenv("ENV")
		if env == "" {
			env = "development"
		}

		database := jsonMap[env].(map[string]interface{})

		dbUser = fmt.Sprintf("%v", database["DBUSER"])
		dbPassword = fmt.Sprintf("%v", database["DBPASSWORD"])
		dbName = fmt.Sprintf("%v", database["DBNAME"])
		dbHost = fmt.Sprintf("%v", database["DBHOST"])
		dbPort, _ = strconv.Atoi(fmt.Sprintf("%v", database["DBPORT"]))
	} else {
		dbUser = os.Getenv("DBUSER")
		dbPassword = os.Getenv("DBPASSWORD")
		dbName = os.Getenv("DBNAME")
		dbHost = os.Getenv("DBHOST")
		dbPort, _ = strconv.Atoi(os.Getenv("DBPORT"))
	}

	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)
	return connectionString
}
