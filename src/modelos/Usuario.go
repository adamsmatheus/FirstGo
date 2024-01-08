package modelos

import (
	"errors"
	"strings"
	"time"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O campo nome e obrigatorio!")
	}

	if usuario.Nick == "" {
		return errors.New("O campo nick e obrigatorio!")
	}

	if usuario.Email == "" {
		return errors.New("O campo email e obrigatorio!")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O campo senha e obrigatorio!")
	}

	return nil
}

func (usuario *Usuario) RemoveEspaco() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Nick = strings.TrimSpace(usuario.Nick)

}

// Método que será chamado pelo controller.
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	usuario.RemoveEspaco()
	return nil

}
