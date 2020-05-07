package main

import (
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"golang-graphql-service/mutations"
	"golang-graphql-service/queries"
	"log"
	"net/http"
)

var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query:    queries.RootQuery,
	Mutation: mutations.RootMutation,
})

func main() {
	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
	})
	http.Handle("/graphql", disableCors(h))
	log.Println("Now server is running on port 8009")
	http.ListenAndServe(":8009", nil)
}
func disableCors(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Max-Age", "86400")
			w.WriteHeader(http.StatusOK)
			return
		}
		h.ServeHTTP(w, r)
	})
}
