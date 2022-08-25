package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"golang/graphql/app/graph/generated"
	"golang/graphql/app/graph/model"
)

// Register is the resolver for the register field.
func (r *mutationResolver) Register(ctx context.Context, input *model.Register) (*model.ResponseRegister, error) {
	return r.Resolver.User.InsertUsert(input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, input *model.Login) (*model.ResponseLogin, error) {
	return r.Resolver.User.SelectUser(input)
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	return nil, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.Resolver.User.SelectUsers(0)
}

// Product is the resolver for the product field.
func (r *queryResolver) Product(ctx context.Context) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

// Products is the resolver for the products field.
func (r *queryResolver) Products(ctx context.Context) ([]*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
