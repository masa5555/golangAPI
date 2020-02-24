package main

import (
  "fmt"
  "net/http"
  "os"
  "time"
  "crypto/sha1"
  "database/sql"
  
  "github.com/go-sql-driver/mysql"
  "./handlers"
)
func main(){
 http.HandleFunc("/user/", handlers.UsersRouter)
 arr3 := [...]string{"One", "Two", "Three"}
 //sha1
 //http.HandleFunc("/", handlers.RootHandler)
 
 server := &http.Server{
		Addr:              ":11111",
		ReadHeaderTimeout: 40 * time.Second,
		ReadTimeout:       40 * time.Second,
	}
	
 err := server.ListenAndServe()
 if err != nil {
   fmt.Println(err)
   os.Exit(1)
 }
}