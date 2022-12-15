package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"go_graphql/graph/generated"
	"go_graphql/graph/model"
)

// CreateAuthor is the resolver for the createAuthor field.
func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.NewAuthor) (*model.User, error) {
	author := &model.User{ID: input.ID, Name: input.Name}
	r.authors = append(r.authors, author)
	return author, nil
}

// CreateVideo is the resolver for the createVideo field.
func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {
	video := &model.Video{ID: input.UserID, Title: input.Title, URL: input.URL}
	r.videos = append(r.videos, video)
	return video, nil
}

// Videos is the resolver for the videos field.
func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	return r.videos, nil
}

// Authors is the resolver for the authors field.
func (r *queryResolver) Authors(ctx context.Context) ([]*model.User, error) {
	return r.authors, nil
}

// Author is the resolver for the author field.
func (r *queryResolver) Author(ctx context.Context, id string) (*model.User, error) {
	var author *model.User
	for i := 0; i < len(r.authors); i++ {
		if r.authors[i].ID == id {
			fmt.Print(r.authors[i])
			author = r.authors[i]
		}
	}
	return author, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
