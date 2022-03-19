package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"log"

	"github.com/khihadysucahyo/orderaja/database"
	"github.com/khihadysucahyo/orderaja/graph/generated"
	"github.com/khihadysucahyo/orderaja/graph/model"
)

var db = database.Connect()

func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (res *model.Item, err error) {
	if res, err = db.Store(&input); err != nil {
		log.Printf("error storing item: %v", err)
	}
	return
}

func (r *queryResolver) Item(ctx context.Context, id string) (res *model.Item, err error) {
	if res, err = db.GetByID(id); err != nil {
		log.Printf("error getting item: %v", err)
	}
	return
}

func (r *queryResolver) Items(ctx context.Context) (res []*model.Item, err error) {
	if res, err = db.Fetch(); err != nil {
		log.Printf("error fetching items: %v", err)
	}
	return
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
