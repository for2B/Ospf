package TreadPit

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/lib/pq"
)

var TestDb *sql.DB

const dbDriverName = "postgres"
const dbStartDataBase = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
const dbHost = "139.199.196.31"
const dbPort = 5432
const dbUser = "ag"
const dbName = "agdb"
const dbPassword = "ag123456"

func InitTestDb(){
	DbInfo := fmt.Sprintf(dbStartDataBase, dbHost, dbPort, dbUser, dbPassword, dbName)
	var err error
	TestDb, err = sql.Open(dbDriverName, DbInfo)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("TestDb Open success!\nbegin pind db serve %v:%v \n",dbHost,dbPassword)
	err = TestDb.Ping()
	if err!=nil{
		fmt.Println("ping fail")
		TestDb.Close()
		log.Fatal(err)
	}
	fmt.Println("ping success!")
}

func GetTestDb() *sql.DB{
	if TestDb == nil{
		InitTestDb()
	}
	return TestDb
}

func ClearTestData(){
	//DoSomeClear
	_, err := TestDb.Exec("delete from project  where project_name like 'T%';")
	if err!=nil {
		fmt.Println("clear mock_data fail:",err)
		return
	}
}
