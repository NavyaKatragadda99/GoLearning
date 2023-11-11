package main

import (
	"database/sql"
	"example/myproject/api"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)
func main() {

    dsn := "root:24021199@tcp(127.0.0.1:3306)/sys?parseTime=true"
    db, err := sql.Open("mysql", dsn)
    if err != nil{
        log.Fatal((err))
    }
    defer db.Close()

    if err := db.Ping(); err !=nil{
        log.Fatal(err)
    }

    api.RegisterRoutes(db)

    log.Println("server starting on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
    
    a := 87
    day := 3
    fmt.Println("Hello, World!")
}