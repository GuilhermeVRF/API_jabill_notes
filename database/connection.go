package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	host = "localhost"
	port = 3306
	user = "root"
	password = "root"
	name = "jabill_notes"
)

func Connect() (*sql.DB, error){
	mysqlConnectionInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s",
		user, password, host, port, name,
	)
	
	connection, err := sql.Open("mysql", mysqlConnectionInfo)
	if err != nil {
		log.Fatal("Erro ao abrir a conexão com o banco de dados:", err)
		return nil, err
	}

	// Verificando a conexão
	err = connection.Ping()
	if err != nil {
		log.Fatal("Erro ao tentar se conectar ao banco de dados:", err)
		return nil, err
	}

	// Exibindo mensagem de sucesso
	fmt.Println("Conectado ao banco de dados " + name)

	return connection, nil
}