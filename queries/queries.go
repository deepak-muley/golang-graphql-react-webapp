package queries

import (
	"github.com/graphql-go/graphql"
	fields "golang-graphql-service/queries/fields"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getNotTodos": fields.GetNotTodos,
	},
})
