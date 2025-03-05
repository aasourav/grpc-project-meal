package graph

import (
	"aas.dev/pkg/handlers"
	"aas.dev/pkg/repository"
	"aas.dev/pkg/services"
	"aas.dev/pkg/utils"
	"github.com/graphql-go/graphql"
)

func getAdminHandler() *handlers.AdminHandler {
	adminRepo := repository.NewAdminRepo(utils.MongoDatabase)
	varificationRepo := repository.NewVerificationRepo(utils.MongoDatabase, true)
	adminService := services.NewAdminService(adminRepo, varificationRepo)
	adminHandler := handlers.NewAdminHandler(adminService)
	return adminHandler
}

var productType = graphql.NewObject(graphql.ObjectConfig{
	Name: "admin",
	Fields: graphql.Fields{
		"id":                      &graphql.Field{Type: graphql.String},
		"userId":                  &graphql.Field{Type: graphql.String},
		"name":                    &graphql.Field{Type: graphql.String},
		"role":                    &graphql.Field{Type: graphql.String},
		"departmentInCharge":      &graphql.Field{Type: graphql.String},
		"createdAt":               &graphql.Field{Type: graphql.String},
		"updatedAt":               &graphql.Field{Type: graphql.String},
		"employeeId":              &graphql.Field{Type: graphql.String},
		"pendingUserApprovalIds":  &graphql.Field{Type: graphql.NewList(graphql.String)},
		"isEmailApproved":         &graphql.Field{Type: graphql.Boolean},
		"isApproved":              &graphql.Field{Type: graphql.Boolean},
		"pendingAdminApprovalIds": &graphql.Field{Type: graphql.NewList(graphql.String)},
		"department":              &graphql.Field{Type: graphql.String},
		"email":                   &graphql.Field{Type: graphql.String},
		"password":                &graphql.Field{Type: graphql.String}},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"getadmins": &graphql.Field{
			Type:        graphql.NewList(productType),
			Description: "Get admins",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.String},
			},
			Resolve: getAdminHandler().GetAdminList,
		},
		"getadmin": &graphql.Field{
			Type:        productType,
			Description: "Get admin by email",
			Args: graphql.FieldConfigArgument{
				"email": &graphql.ArgumentConfig{
					Type: graphql.String},
			},
			Resolve: getAdminHandler().GetAdminByEmail,
		},
	},
})

// var rootMutation = graphql.NewObject(graphql.ObjectConfig{
// 	Name: "RootMutation",
// 	Fields: graphql.Fields{
// 		"registerAdmin": &graphql.Field{
// 			Type:        graphql.String,
// 			Description: "Register a new admin",
// 			Args: graphql.FieldConfigArgument{
// 				"email":    &graphql.ArgumentConfig{Type: graphql.String},
// 				"password": &graphql.ArgumentConfig{Type: graphql.String},
// 				"role":     &graphql.ArgumentConfig{Type: graphql.String},
// 			},
// 			Resolve: getAdminHandler().GetAdminsByEmail,
// 		},
// 		"loginAdmin": &graphql.Field{
// 			Type:        graphql.String,
// 			Description: "Admin login",
// 			Args: graphql.FieldConfigArgument{
// 				"email":    &graphql.ArgumentConfig{Type: graphql.String},
// 				"password": &graphql.ArgumentConfig{Type: graphql.String},
// 			},
// 			Resolve: getAdminHandler().GetAdminsByEmail,
// 		},
// 	},
// })

var Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
	Query: rootQuery,
	// Mutation: rootMutation,
})
