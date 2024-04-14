package main

import (
	"database/sql"
	"fmt"
	"go-sample/controller"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const maxRetry = 5 // 最大リトライ回数

func open(path string, count uint) *sql.DB {
	db, err := sql.Open("mysql", path)
	if err != nil {
		log.Fatal("open error:", err)
	}

	if err = db.Ping(); err != nil {
		if count == 0 {
			log.Fatalf("max retries exceeded: %v", err)
		}
		time.Sleep(time.Second * 2)
		fmt.Printf("retry... count:%v\n", count)
		return open(path, count-1) // リトライ回数をデクリメントして再帰呼び出し
	}

	fmt.Println("db connected!!")
	return db
}

func connectDB() *sql.DB {
	var path string = fmt.Sprintf("%s:%s@tcp(db:3306)/%s?charset=utf8&parseTime=true",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DATABASE"))

	return open(path, maxRetry)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	db := connectDB()
	defer db.Close()
	controller.GetAllUsersHandler(w, r, db)
}

func main() {
	http.HandleFunc("/", handleRequest)
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	http.ListenAndServe(":8080", nil)
}
