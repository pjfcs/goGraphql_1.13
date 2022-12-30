package src

import (
	"context"

	graphql1 "github.com/pjfcs/goGraphQL_1.13/graphql"
	"github.com/pjfcs/goGraphQL_1.13/models"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Meetup() graphql1.MeetupResolver {
	return &meetupResolver{r}
}
func (r *Resolver) Mutation() graphql1.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() graphql1.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() graphql1.UserResolver {
	return &userResolver{r}
}

type meetupResolver struct{ *Resolver }

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	panic("not implemented")
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateMeetup(ctx context.Context, input graphql1.NewMeetup) (*models.Meetup, error) {
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	panic("not implemented")
}

type userResolver struct{ *Resolver }

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	panic("not implemented")
}
