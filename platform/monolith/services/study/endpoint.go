package study

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/google/uuid"
)

type createStudyProjectRequest struct {
	Name string  `json:"name"`
	Icon *string `json:"icon"`
}

type createStudyProjectResponse struct {
	StudyProject StudyProject `json:"study_project"`
}

type createOrReplaceStudyProjectCardsRequest struct {
	StudyProjectID string `json:"study_project_id"`
}

type createOrReplaceStudyProjectCardsResponse struct {
	Cards []StudyProjectCard `json:"cards"`
}

type Endpoints struct {
	CreateStudyProjectEndpoint               endpoint.Endpoint
	CreateOrReplaceStudyProjectCardsEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		CreateStudyProjectEndpoint:               MakeCreateStudyProjectEndpoint(svc),
		CreateOrReplaceStudyProjectCardsEndpoint: MakeCreateOrReplaceStudyProjectCardsEndpoint(svc),
	}
}

func MakeCreateStudyProjectEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createStudyProjectRequest)
		// TODO: get userID from context
		studyProject, err := svc.CreateStudyProject(ctx, uuid.New().String(), req.Name, req.Icon)
		if err != nil {
			return nil, err
		}
		return createStudyProjectResponse{StudyProject: studyProject}, nil
	}
}

func MakeCreateOrReplaceStudyProjectCardsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createOrReplaceStudyProjectCardsRequest)
		// TODO: get userID from context
		cards, err := svc.CreateOrReplaceStudyProjectCards(ctx, uuid.New().String(), req.StudyProjectID)
		if err != nil {
			return nil, err
		}
		return createOrReplaceStudyProjectCardsResponse{Cards: cards}, nil
	}
}
