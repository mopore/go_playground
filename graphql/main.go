package main

import (
	"encoding/json"
	"log"

	"github.com/graphql-go/graphql"
)

type Comment struct {
        Body string
}

type Author struct {
    Name string
    Books []int
}

type Book struct {
	Title    string
	Author   string
	Comments []Comment
}

func populate() []Book {
    author := &Author{
        Name: "Author Name",
        Books: []int{1},
    }
    book := Book{
        ID:   1,
        Title: "Book Title",
        Author: *author,
        Comments: []Comment{
            Comment{Body: "Comment 1"},
            Comment{Body: "Comment 2"},
        },
    }

    var books []Book
    books = append(books, book)

    return books
    }



func main() {
	// Call the function
	log.Println("Hello, World!")

	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return "world", nil
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	query := `
        {
            hello
        }
    `

	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, err := json.Marshal(r)
	if err != nil {
		log.Fatalf("failed to marshal result: %v", err)
	}

	log.Printf("Result from graphql: %s", rJSON)

}
