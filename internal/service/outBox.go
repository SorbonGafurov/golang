package service

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func OutBoxMessage() {
	connStr := "user=postgres password=Sorbon260494 dbname=Outbox sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("sds")
	}
	defer db.Close()

	//result, err := db.Exec("INSERT INTO Products (model, company, price) VALUES ('iPhone X', 'Apple', 72000)")

	rows, err := db.Query("select o.message from outboxmessages o")
	if err != nil {
		panic("sds")
	}
	defer rows.Close()

	var message string
	for rows.Next() {
		rows.Scan(&message)
	}

	println(&message)
}
