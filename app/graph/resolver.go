package graph

import (
	"golang/graphql/app/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	User repository.UserRepo
}
