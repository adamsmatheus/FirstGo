package banco

import (
	"GoTest/src/config"
	"database/sql"
	_ "github.com/go-sql-driver/mysql" //Driver
)

// Metodo para conectar no banco
func Conectar() (*sql.DB, error) {

	db, erro := sql.Open("mysql", config.StringConexao)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
