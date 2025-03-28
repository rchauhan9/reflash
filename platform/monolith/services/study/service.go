package study

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rchauhan9/reflash/monolith/common/clients/card_creator"
)

type Service interface {
	CreateStudyProject(ctx context.Context) (StudyProject, error)
	CreateOrReplaceStudyProjectCards(ctx context.Context, studyProjectID string) error
}

func NewService(repository Repository, cardCreatorClient card_creator.Client) Service {
	return &service{
		repository:        repository,
		cardCreatorClient: cardCreatorClient,
	}
}

type service struct {
	repository        Repository
	cardCreatorClient card_creator.Client
}

func (s *service) CreateStudyProject(ctx context.Context) (StudyProject, error) {
	return s.repository.CreateStudy(ctx)
}

func (s *service) CreateOrReplaceStudyProjectCards(ctx context.Context, studyProjectID string) error {

	cards, err := s.cardCreatorClient.CreateCards(ctx)
	if err != nil {
		return errors.Wrapf(err, "could not create cards via client")
	}

	// in a transaction
	err = s.repository.DeleteCards(ctx, studyProjectID)
	err = s.repository.CreateCards(ctx, studyProjectID, cards)

	return nil
}
