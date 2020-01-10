package main

import(
        _"github.com/go-sql-driver/mysql"
        "database/sql"
        "fmt"
        "log"
)

func main() {
        cnn, err := sql.Open("mysql", "docker:docker@tcp(db:3306)/test_db")
        if err != nil {
			log.Fatal(err)
        }

        id := 1
        var name string

        if err := cnn.QueryRow("SELECT name FROM test_tb").Scan(&name); err != nil {
			log.Fatal(err)
        }

	fmt.Println(id, name)
}