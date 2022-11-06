package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/leoneville/api-book/src/database"
	"github.com/leoneville/api-book/src/models"
	"github.com/leoneville/api-book/src/repository"
	"github.com/leoneville/api-book/src/response"
)

func PostBook(w http.ResponseWriter, r *http.Request) {
	req, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var book models.Book

	if err = json.Unmarshal(req, &book); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = book.Prepare(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.CreateNewRepository(db)
	book.ID, err = repo.Create(book)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusCreated, book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	titleOrAuthorOrComPub := fmt.Sprintf("%%%s%%", strings.ToLower(r.URL.Query().Get("book")))

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.CreateNewRepository(db)
	books, err := repo.Search(titleOrAuthorOrComPub)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if len(books) <= 0 {
		response.Error(w, http.StatusNotFound, errors.New("nenhum livro foi encontrado"))
		return
	}

	response.JSON(w, http.StatusOK, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.CreateNewRepository(db)
	book, err := repo.SearchOne(ID)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	if book.ID <= 0 {
		response.Error(w, http.StatusNotFound, errors.New("nenhum livro foi encontrado"))
		return
	}

	response.JSON(w, http.StatusOK, book)

}

func PutBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	req, err := io.ReadAll(r.Body)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	var book models.Book
	if err = json.Unmarshal(req, &book); err != nil {
		response.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	if err = book.Prepare(); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.CreateNewRepository(db)
	if err = repo.Update(ID, book); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repository.CreateNewRepository(db)
	if err = repo.Delete(ID); err != nil {
		response.Error(w, http.StatusBadRequest, err)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)
}
