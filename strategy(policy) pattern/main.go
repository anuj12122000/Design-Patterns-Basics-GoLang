package main

import "fmt"

type IDBConnection interface { // this can be assigned to any type of Db at the time of runtime
	connect()
}

type DbConnection struct {
	db IDBConnection
}

func (dbc *DbConnection) DBConnect() {
	dbc.db.connect()
}

type mysql struct {
}

func (sql *mysql) connect() {
	fmt.Println("MY SQL DB CONNECTED")
}

type mongoDb struct {
}

func (mongoDb *mongoDb) connect() {
	fmt.Println("Mongo Db Connected")
}

func main() {
	mysqlConnection := &DbConnection{db: &mysql{}}
	mysqlConnection.db.connect()

	mongoConnection := &DbConnection{db: &mongoDb{}}
	mongoConnection.db.connect()

}
