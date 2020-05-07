package fields

import (
	//"context"

	"github.com/graphql-go/graphql"
	//"go.mongodb.org/mongo-driver/bson"
	//
	//"golang-graphql-service/data"
	types "golang-graphql-service/types"
)

type todoStruct struct {
	NAME        string `json:"name"`
	DESCRIPTION string `json:"description"`
}

var GetNotTodos = &graphql.Field{
	Type:        graphql.NewList(types.NotTodo),
	Description: "Get all not todos",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {

		//notTodoCollection := mongo.Client.GetInternalClient().Database("medium-app").Collection("Not_Todos")
		var todosList []todoStruct
		//todos, err := notTodoCollection.Find(context.Background(), nil)
		//if err != nil { panic(err) }
		//for todos.Next(context.Background()) {
		//	doc := bson.D{}
		//	err := todos.Decode(doc)
		//	if err != nil { panic(err) }
		//	keys, err := doc.M.Keys(false)
		//	if err != nil { panic(err) }
		//	// convert BSON to struct
		//	todo := todoStruct{}
		//	for _, key := range keys {
		//		keyString := key.String()
		//		elm, err := doc.M.Lookup(keyString)
		//		if err != nil { panic(err) }
		//		switch (keyString) {
		//		case "name":
		//			todo.NAME = elm.Value().StringValue()
		//		case "description":
		//			todo.DESCRIPTION = elm.Value().StringValue()
		//		default:
		//		}
		//	}
		//	todosList = append(todosList, todo)
		//}

		return todosList, nil
	},
}
