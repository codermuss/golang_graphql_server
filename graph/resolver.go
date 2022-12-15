package graph

import "go_graphql/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	//* Video listesin dönecek
	videos  []*model.Video
	authors []*model.User
	//author  *model.User
}
