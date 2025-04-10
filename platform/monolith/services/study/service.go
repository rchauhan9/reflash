package study

import (
	"context"
	"fmt"
	"github.com/go-kit/log"
	"github.com/pkg/errors"
	"github.com/rchauhan9/reflash/monolith/common/clients/card_creator"
	"mime/multipart"
)

type Service interface {
	ListStudyProjects(ctx context.Context, userID string) ([]StudyProject, error)
	CreateStudyProject(ctx context.Context, userID string, name string, icon *string) (StudyProject, error)
	CreateProjectFile(ctx context.Context, userID string, studyProjectID string, filename string, file *multipart.FileHeader) (*string, error)
	ListCards(ctx context.Context, userID string, studyProjectID string) ([]StudyProjectCard, error)
	CreateOrReplaceStudyProjectCards(ctx context.Context, userID string, studyProjectID string) ([]StudyProjectCard, error)
}

func NewService(repository Repository, cardCreatorClient card_creator.Client, logger log.Logger) Service {
	return &service{
		repository:        repository,
		cardCreatorClient: cardCreatorClient,
		logger:            logger,
	}
}

type service struct {
	repository        Repository
	cardCreatorClient card_creator.Client
	logger            log.Logger
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

func (s *service) CreateProjectFile(ctx context.Context, userID string, studyProjectID string, filename string, file *multipart.FileHeader) (*string, error) {
	openedFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer openedFile.Close()

	bucket := "./output/" + userID
	key := fmt.Sprintf("/%v/%v", studyProjectID, filename)
	err = UploadFile(ctx, bucket, key, openedFile)
	if err != nil {
		return nil, err
	}

	path := bucket + key
	fileReferenceID, err := s.repository.CreateProjectFile(ctx, userID, studyProjectID, path)
	if err != nil {
		return nil, err
	}

	return &fileReferenceID, nil
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
	s.logger.Log("msg", "created cards via client", "cards", fmt.Sprintf("%+v", cards))
	if err != nil {
		return nil, errors.Wrapf(err, "could not create cards via client")
	}

	// in a transaction
	err = s.repository.DeleteCards(ctx, userID, studyProjectID)

	createCards := make([]CreateCard, len(cards))
	for i, card := range cards {
		createCards[i] = CreateCard{
			UserID:         userID,
			StudyProjectID: studyProjectID,
			Question:       card.Question,
			Answer:         card.Answer,
		}
	}

	newCards, err := s.repository.CreateCards(ctx, createCards)
	if err != nil {
		return nil, errors.Wrapf(err, "could not create cards in database")
	}

	return newCards, nil
}
