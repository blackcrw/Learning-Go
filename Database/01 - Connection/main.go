package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

/*
To connect to MySQL, it is necessary to download the library "github.com/go-sql-driver/mysql",
even though Golang already has the option for SQL, databases, etc., it does not come with the driver by default.
When we import the driver, we need to put an _(underscore) before it,
because this file itself does not use/instantiate the library.
And exactly for this reason, we don't instantiate MySQL itself in "sql.Open()",
we provide a string with the name of the SQL driver.

The MySQL URI should be like this: "user:password@/database".

*/

func main() {
	database, err := sql.Open("mysql", "root:toortoor@/golang?charset=utf8")

	if err != nil {
		log.Fatal("An error occurred:", err)
	}

	defer database.Close()

	if err = database.Ping(); err != nil {
		log.Fatal("Couldn't connect to the database!", err)
	}

	fmt.Println("Connected to the database!")

	sql, err := database.Query("select * from users")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(sql)

	defer sql.Close()
}
