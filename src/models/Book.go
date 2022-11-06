package models

import (
	"errors"
	"strings"
	"time"
)

type Book struct {
	ID         uint64    `json:"id"`
	Titulo     string    `json:"titulo"`
	Autor      string    `json:"autor"`
	QtdPaginas int       `json:"qtd_paginas"`
	Editora    string    `json:"editora"`
	CriadoEm   time.Time `json:"CriadoEm,omitempty"`
}

func (b *Book) Prepare() error {
	if err := b.validate(); err != nil {
		return err
	}

	b.formate()

	return nil
}

func (b Book) validate() error {
	if b.Titulo == "" {
		return errors.New("o campo titulo não pode estar vazio")
	}

	if b.Autor == "" {
		return errors.New("o campo autor não pode estar vazio")
	}

	if b.Editora == "" {
		return errors.New("o campo editora não pode estar vazio")
	}

	if b.QtdPaginas == 0 {
		return errors.New("o campo qtd_paginas não pode estar vazio")
	}

	return nil
}

func (b *Book) formate() {
	b.Autor = strings.TrimSpace(b.Autor)
	b.Titulo = strings.TrimSpace(b.Titulo)
	b.Editora = strings.TrimSpace(b.Editora)
}
