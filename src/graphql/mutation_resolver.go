package graphql

import (
	"context"
	"errors"

	"github.com/pjfcs/goGraphQL_1.13/models"
)

type mutationResolver struct{ *Resolver }

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

func (m *mutationResolver) CreateMeetup(ctx context.Context, input NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("Name not long enough")
	}

	if len(input.Description) < 3 {
		return nil, errors.New("Description not long enough")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}

	return m.MeetupsRepo.CreateMeetup(meetup)

}
