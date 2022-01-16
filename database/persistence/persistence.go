package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	OUTPUT = "OUTPUT"
	INPUT  = "INPUT"
)

type Persistence interface {
	Insert(data int, table string) error
	Update(id string, data int, table string) error
	Delete(id string) error
	CreateTable(tableName string) error
	Rollback() error
	Init() error
}

type DbImpl struct {
	Host     string
	Port     int
	User     string
	Password string
	pers     *sql.DB
}

var pers Persistence

func GetPers() Persistence {
	return pers
}

//Init get called on service startup
func (db *DbImpl) Init() error {
	//todo: error handle
	db.GetConnect()
	db.CreateTable(OUTPUT)
	db.CreateTable(OUTPUT)

	return nil
}

//
//Connect use the crediential get from config to connect
func (db *DbImpl) GetConnect() bool {

	if !db.IsConnected() {
		psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", db.Host, db.Port, db.User, db.Password)
		for i := 0; i < 10; i++ {
			if myPers, err := sql.Open("postgres", psqlconn); err != nil {
				log.Println("Fail to connect to db")
				time.Sleep(10 * time.Second)
			} else {
				db.pers = myPers
				break
			}
		}
		return false
	}
	return true
}

//IsConnected check if db if connected
func (db *DbImpl) IsConnected() bool {
	return true
}

//Insert not familar with the postgres, use emulated query
func (db *DbImpl) Insert(data int, table string) error {
	if !db.GetConnect() {
		log.Print("Error connecting to DB")
	}
	insertDynStmt := `Not familar with postgres `
	if _, err := db.pers.Exec(insertDynStmt, data, table); err != nil {
		db.Rollback()
	}
	return nil
}
func (db *DbImpl) Update(id string, data int, table string) error {
	return nil
}
func (db *DbImpl) Delete(id string) error {
	return nil
}
func (db *DbImpl) CreateTable(tableName string) error {
	return nil
}
func (db *DbImpl) Rollback() error {
	return nil
}
