package repositorios

import (
	"GoTest/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?,?,?,?)")
	if erro != nil {
		return 0, erro
	}

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil

	return 0, nil
}

// Busca pelos registros por filtro
func (repositorio Usuarios) Buscar(nomeouNick string) ([]modelos.Usuario, error) {
	nomeouNick = fmt.Sprintf("%%%s%%", nomeouNick)
	linhas, erro := repositorio.db.Query(""+
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?", nomeouNick, nomeouNick)
	if erro != nil {
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			usuario.Email,
			usuario.CriadoEm,
		); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorID(id uint64) (modelos.Usuario, error) {
	linhas, erro := repositorio.db.Query(
		"select * from usuarios where id = ?", id)
	if erro != nil {
		return modelos.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario modelos.Usuario
	if linhas.Next() {
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil {
			return modelos.Usuario{}, erro
		}
	}

	return usuario, nil
}

func (repositorio Usuarios) Atualizar(id uint64, usuarios modelos.Usuario) error {
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(usuarios.Nome, usuarios.Nick, usuarios.Email, id); erro != nil {
		return erro
	}

	return nil
}

func (repositorio Usuarios) Deletar(id uint64) error {
	statement, erro := repositorio.db.Prepare(
		"delete from usuarios where id = ?")
	if erro != nil {
		return erro
	}

	defer statement.Close()

	if _, erro = statement.Exec(id); erro != nil {
		return erro
	}

	return nil
}
