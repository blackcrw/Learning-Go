package database

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Driver do MySql
)

type userJson struct {
	ID    uint64 `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Connect  :: Função de conexão ao Banco de Dados.
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

// Disconnect :: Função que irá desconectar/fechar conexão com o banco de dados.
func Disconnect(database *sql.DB) {
	database.Close()
}

/* Create ::
Aqui eu poderia passar como paramêtro da função o request completo, mas não passarei pôs não á necessidade (\n)
Visto que o que apenas queremos é o corpo da requisição (Body).
*/

// Create :: Função de criação de usuários no banco de dados, ela recebe o connect via paramêtro.
func Create(database *sql.DB, response http.ResponseWriter, body []byte) (int64, error) {
	var user userJson

	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
		response.Write([]byte("Erro ao converter o usuário para Struct"))
		return 0, err
	}

	statement, err := database.Prepare("insert into users (name, email) values (?, ?)")

	if err != nil {
		log.Println(err)
		response.Write([]byte("Ocorreu um erro na hora de preparar a query."))
		return 0, err
	}

	defer database.Close()
	defer statement.Close()

	insert, err := statement.Exec(user.Name, user.Email)

	if err != nil {
		log.Println(err)
		response.Write([]byte("Ocorreu um erro ao inserir este usuário no banco de dados!"))
		return 0, err
	}

	id, err := insert.LastInsertId()

	if err != nil {
		log.Println(err)
		response.Write([]byte("Ocorreu um erro ao tentar pegar o ID."))
		return 0, err
	}

	return id, nil
}

// Delete :: Função para apagar os usuários por ID.
func Delete(database *sql.DB, response http.ResponseWriter, ID uint64) error {
	statement, err := database.Prepare("delete from users where id = ?")

	if err != nil {
		log.Println(err)
		response.Write([]byte("Erro ao deletar usuário por ID"))
		return err
	}

	defer statement.Close()
	defer database.Close()

	if _, err := statement.Exec(ID); err != nil {
		log.Println(err)
		response.Write([]byte("Erro ao deletar usuário!"))
		return err
	}

	return nil
}

// Get :: Está função serve para puxarmos os usuários por ID, e imprimir na tela.
func Get(database *sql.DB, response http.ResponseWriter, ID uint64) error {
	var user userJson

	query, err := database.Query("select * from users where id = ?", ID)

	if err != nil {
		log.Println(err)
		response.Write([]byte("Erro ao buscar usuário por ID"))
		return err
	}

	defer query.Close()
	defer database.Close()

	if query.Next() {
		if err := query.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Println(err)
			response.Write([]byte("Erro ao escanear usuários."))
			return err
		}
	}

	if err := json.NewEncoder(response).Encode(user); err != nil { // Aqui é onde o retorno da request é mostrado. E caso haja algum erro ele é retornado.
		log.Println(err)
		response.Write([]byte("Erro ao converter usuários para Json"))
		return err
	}

	return nil
}

// GetAll :: Está função serve para puxarmos todos os "usuários" do banco de dados e imprimir na tela.
func GetAll(database *sql.DB, response http.ResponseWriter) error {
	var users []userJson

	query, err := database.Query("select * from users")

	if err != nil {
		log.Println(err)
		response.Write([]byte("Erro ao fazer a query"))
		return err
	}

	defer query.Close()
	defer database.Close()

	for query.Next() {
		var user userJson

		if err := query.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			log.Println(err)
			response.Write([]byte("Erro ao escanear usuários."))
			return err
		}

		users = append(users, user)
	}

	if err := json.NewEncoder(response).Encode(users); err != nil { // Aqui é onde o retorno da request é mostrado. E caso haja algum erro ele é retornado.
		log.Println(err)
		response.Write([]byte("Erro ao converter usuários para Json"))
		return err
	}

	return nil
}

// Update :: Está função realiza um update dos dados dos usuários no banco de dados; nome e email.
func Update(database *sql.DB, response http.ResponseWriter, body []byte, ID uint64) error {
	var user userJson

	if err := json.Unmarshal(body, &user); err != nil {
		log.Println(err)
		response.Write([]byte("Erro ao converter usuários para Json"))
		return err
	}

	statement, err := database.Prepare("update users set name = ?, email = ? where id = ?")
	if err != nil {
		log.Println(err)
		response.Write([]byte("Erro ao preparar query"))
		return err
	}

	defer statement.Close()
	defer database.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		log.Println(err)
		response.Write([]byte("Erro ao realizar update"))
		return err
	}

	return nil
}
