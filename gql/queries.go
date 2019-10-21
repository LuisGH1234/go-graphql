package gql

import "github.com/graphql-go/graphql"

// Root ...
type Root struct {
	Query *graphql.Object
}

// QueryRoot ...
func QueryRoot() *Root {
	return &Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "QueryRoot",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						Type: graphql.NewList(UserType),
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							_, ok := p.Args["name"].(string)
							if ok {
								var users []User
								for i := 0; i < 3; i++ {
									n := User{ID: i, Name: "sd", Address: "as"}
									users = append(users, n)
								}
								return users, nil
							}
							return nil, nil
						},
					},
				},
			},
		),
	}
}
