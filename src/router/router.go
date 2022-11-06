package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/leoneville/api-book/src/controllers"
)

type route struct {
	URI      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
}

var booksRoutes = []route{
	{
		URI:      "/books",
		Method:   http.MethodPost,
		Function: controllers.PostBook,
	},
	{
		URI:      "/books",
		Method:   http.MethodGet,
		Function: controllers.GetBooks,
	},
	{
		URI:      "/books/{id}",
		Method:   http.MethodGet,
		Function: controllers.GetBook,
	},
	{
		URI:      "/books/{id}",
		Method:   http.MethodPut,
		Function: controllers.PutBook,
	},
	{
		URI:      "/books/{id}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteBook,
	},
}

func GenerateRoutes() *mux.Router {
	r := mux.NewRouter()

	for _, route := range booksRoutes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
