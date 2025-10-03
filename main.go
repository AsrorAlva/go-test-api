package main
 
import (
    "example.com/go-test-api/db"
    "example.com/go-test-api/router"
)

 
func main() {
   db.InitPostgresDB()
   router.InitRouter().Run()
}