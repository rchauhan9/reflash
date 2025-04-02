package study

import (
	"context"
	"github.com/pkg/errors"
	"github.com/rchauhan9/reflash/monolith/common/clients/card_creator"
)

type Service interface {
	ListStudyProjects(ctx context.Context, userID string) ([]StudyProject, error)
	CreateStudyProject(ctx context.Context, userID string, name string, icon *string) (StudyProject, error)
	ListCards(ctx context.Context, userID string, studyProjectID string) ([]StudyProjectCard, error)
	CreateOrReplaceStudyProjectCards(ctx context.Context, userID string, studyProjectID string) ([]StudyProjectCard, error)
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

func (s *service) ListStudyProjects(ctx context.Context, userID string) ([]StudyProject, error) {
	studyProjects, err := s.repository.ListStudyProjects(ctx, userID)
	if err != nil {
		return nil, errors.Wrapf(err, "could not list study projects")
	}
	return studyProjects, nil
}

func (s *service) CreateStudyProject(ctx context.Context, userID string, name string, icon *string) (StudyProject, error) {
	studyProjectId, err := s.repository.CreateStudyProject(ctx, userID, name, icon)
	if err != nil {
		return StudyProject{}, errors.Wrapf(err, "could not create study project")
	}

	return StudyProject{
		ID:   studyProjectId,
		Name: name,
		Icon: icon,
	}, nil
}

func (s *service) ListCards(ctx context.Context, userID string, studyProjectID string) ([]StudyProjectCard, error) {
	cards, err := s.repository.ListCards(ctx, userID, studyProjectID)
	if err != nil {
		return nil, errors.Wrapf(err, "could not list study project cards")
	}
	return cards, nil
}

func (s *service) CreateOrReplaceStudyProjectCards(ctx context.Context, userID string, studyProjectID string) ([]StudyProjectCard, error) {

	cards, err := s.cardCreatorClient.CreateCards(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create cards via client")
	}

	// in a transaction
	err = s.repository.DeleteCards(ctx, userID, studyProjectID)

	createCards := make([]CreateCard, len(cards))
	for _, card := range cards {
		createCards = append(createCards, CreateCard{
			StudyProjectID: studyProjectID,
			Question:       card.Question,
			Answer:         card.Answer,
		})
	}

	newCards, err := s.repository.CreateCards(ctx, userID, createCards)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create cards in database")
	}

	return newCards, nil
}
