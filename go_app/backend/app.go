//package main is needed to  make it runable 

package main


//imports database and should work if we decide to move over to SQL
import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main(){
	
	//opens whoknows.db if null creates whoknows.db
	db, err :=sql.Open("sqlite", "file:whoknows.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//error handling
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}


	//print so we know if database is connected 
	fmt.Println("SQLite connected!")
}
