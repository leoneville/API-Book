package repository

import (
	"database/sql"

	"github.com/leoneville/api-book/src/models"
)

type books struct {
	db *sql.DB
}

func CreateNewRepository(db *sql.DB) *books {
	return &books{db}
}

func (b books) Create(book models.Book) (uint64, error) {
	stmt, err := b.db.Prepare("INSERT INTO books (titulo, autor, qtd_paginas, editora) VALUES (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(book.Titulo, book.Autor, book.QtdPaginas, book.Editora)
	if err != nil {
		return 0, err
	}

	ID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ID), nil
}

func (b books) Search(titleOrAuthorOrComPub string) ([]models.Book, error) {
	rows, err := b.db.Query("SELECT id, titulo, autor, qtd_paginas, editora, criadoEm FROM books WHERE titulo LIKE ? OR autor LIKE ? OR editora LIKE ?", titleOrAuthorOrComPub, titleOrAuthorOrComPub, titleOrAuthorOrComPub)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book

	for rows.Next() {
		var book models.Book
		if err = rows.Scan(
			&book.ID,
			&book.Titulo,
			&book.Autor,
			&book.QtdPaginas,
			&book.Editora,
			&book.CriadoEm,
		); err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (b books) SearchOne(ID uint64) (models.Book, error) {

	row, err := b.db.Query("SELECT id, titulo, autor, qtd_paginas, editora, criadoEm FROM books WHERE id = ?", ID)
	if err != nil {
		return models.Book{}, err
	}
	defer row.Close()

	var book models.Book

	if row.Next() {
		if err = row.Scan(
			&book.ID,
			&book.Titulo,
			&book.Autor,
			&book.QtdPaginas,
			&book.Editora,
			&book.CriadoEm,
		); err != nil {
			return models.Book{}, err
		}
	}

	return book, nil
}

func (b books) Update(ID uint64, book models.Book) error {
	stmt, err := b.db.Prepare(
		"UPDATE books SET titulo = ?, autor = ?, qtd_paginas = ?, editora = ? WHERE id = ?",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(book.Titulo, book.Autor, book.QtdPaginas, book.Editora, ID); err != nil {
		return err
	}

	return nil
}

func (b books) Delete(ID uint64) error {
	stmt, err := b.db.Prepare("DELETE FROM books WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(ID); err != nil {
		return err
	}

	return nil
}
