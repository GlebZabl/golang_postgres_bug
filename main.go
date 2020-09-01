package main

import (
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/lib/pq"
)

func main()  {
	err := selectFromDb(65536)
	if err != nil{
		panic(err)
	}

	println("where was no errors!")
}

func selectFromDb(argsNumber int) (err error){
	//open connect to postgres database, put your database connection string below, instead of second argument
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/test?sslmode=disable")
	if err != nil {
		return
	}

	//preparing query, you can call this func with argsNumber = 10(or some not huge number)
	//to check that it is working correct
	var args []interface{}
	query := "SELECT id FROM test_tbl WHERE id IN ("
	for i := 0; i < argsNumber; i++ {
		args = append(args, strconv.Itoa(i))
		query = fmt.Sprintf("%s $%d, ", query, i+1)
		fmt.Printf("preparing query, %d/%d \n", i, argsNumber)
	}
	query = fmt.Sprintf("%s);", query[:len(query)-2])
	fmt.Println("preparing finished:")

	//uncomment string bellow to see result query and check it
	//fmt.Println(query)


	//preparing query
	stmt, err := db.Prepare(query)
	if err != nil {
		return
	}

	//selecting
	rows, err := stmt.Query(args...)
	if err != nil{
		return
	}

	//closing db connection
	err = rows.Close()
	if err != nil{
		return
	}

	return
}
