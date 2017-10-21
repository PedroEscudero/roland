package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

type Route struct {
    Method      string
    Path    string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.
            Methods(route.Method).
            Path(route.Path).
            Handler(route.HandlerFunc)
    }

    return router
}

var routes = Routes{
    Route{
        "GET",
        "/bookse",
        GetBooks,
    },
    Route{
        "POST",
        "/books/{id}",
        CreateBook,
    },
    Route{
        "GET",
        "/books/{id}",
        GetBook,
    },
}