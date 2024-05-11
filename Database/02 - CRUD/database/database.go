package database

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // MySql Driver
)

type userJson struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Connect  :: Function to connect to the Database.
func Connect() (*sql.DB, error) {
	database, err := sql.Open("mysql", "root:toortoor@/golang?charset=utf8")

	if err != nil {
		return nil, err
	}

	if err = database.Ping(); err != nil {
		return nil, err
	}

	return database, nil
}

// Disconnect :: Function to disconnect/close the database connection.
func Disconnect(database *sql.DB) {
	database.Close()
}

/* Create ::
Here I could pass the complete request as a function parameter, but I won't because there is no need (\n)
since we only want the request body.
*/

// Create :: Function to create users in the database, it receives the database connection as a parameter.
func Create(database *sql.DB, response http.ResponseWriter, body []byte) (int64, error) {
	var user userJson

	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
		response.Write([]byte("Error converting user to Struct"))
		return 0, err
	}

	statement, err := database.Prepare("insert into users (name, email) values (?, ?)")

	if err != nil {
		log.Println(err)
		response.Write([]byte("Error preparing query."))
		return 0, err
	}

	defer database.Close()
	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)

	if err != nil {
		log.Println(err)
		response.Write([]byte("Error inserting user into the database!"))
		return 0, err
	}

	id, err := insert.LastInsertId()

	if err != nil {
		log.Println(err)
		response.Write([]byte("Error getting ID."))
		return 0, err
	}

	return id, nil
}

// Delete :: Function to delete users by ID.
func Delete(database *sql.DB, response http.ResponseWriter, ID uint64) error {
	statement, err := database.Prepare("delete from users where id = ?")

	if err != nil {
		log.Println(err)
		response.Write([]byte("Error deleting user by ID"))
		return err
	}

	defer statement.Close()
	defer database.Close()

	if _, err := statement.Exec(ID); err != nil {
		log.Println(err)
		response.Write([]byte("Error deleting user!"))
		return err
	}

	return nil
}

// Get :: This function is used to fetch users by ID and print them on the screen.
func Get(database *sql.DB, response http.ResponseWriter, ID uint64) error {
	var user userJson

	query, err := database.Query("select * from users where id = ?", ID)

	if err != nil {
		log.Println(err)
		response.Write([]byte("Error fetching user by ID"))
		return err
	}

	defer query.Close()
	defer database.Close()

	if query.Next() {
		if err := query.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Println(err)
			response.Write([]byte("Error scanning users."))
			return err
		}
	}

	if err := json.NewEncoder(response).Encode(user); err != nil { // This is where the request response is shown. And if there is an error, it is returned.
		log.Println(err)
		response.Write([]byte("Error converting users to Json"))
		return err
	}

	return nil
}

// GetAll :: This function is used to fetch all "users" from the database and print them on the screen.
func GetAll(database *sql.DB, response http.ResponseWriter) error {
	var users []userJson

	query, err := database.Query("select * from users")

	if err != nil {
		log.Println(err)
		response.Write([]byte("Error executing query"))
		return err
	}

	defer query.Close()
	defer database.Close()

	for query.Next() {
		var user userJson

		if err := query.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Println(err)
			response.Write([]byte("Error scanning users."))
			return err
		}

		users = append(users, user)
	}

	if err := json.NewEncoder(response).Encode(users); err != nil { // This is where the request response is shown. And if there is an error, it is returned.
		log.Println(err)
		response.Write([]byte("Error converting users to Json"))
		return err
	}

	return nil
}

// Update :: This function updates user data in the database; name and email.
func Update(database *sql.DB, response http.ResponseWriter, body []byte, ID uint64) error {
	var user userJson

	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
		response.Write([]byte("Error converting users to Json"))
		return err
	}

	statement, err := database.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		log.Println(err)
		response.Write([]byte("Error preparing query"))
		return err
	}

	defer statement.Close()
	defer database.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		log.Println(err)
		response.Write([]byte("Error performing update"))
		return err
	}

	return nil
}
