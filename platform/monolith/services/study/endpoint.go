package study

import (
	"context"
	"github.com/go-kit/kit/endpoint"
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
		studyProject, err := svc.CreateStudyProject(ctx, req.Name, req.Icon)
		if err != nil {
			return nil, err
		}
		return createStudyProjectResponse{StudyProject: studyProject}, nil
	}
}

func MakeCreateOrReplaceStudyProjectCardsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createOrReplaceStudyProjectCardsRequest)
		cards, err := svc.CreateOrReplaceStudyProjectCards(ctx, req.StudyProjectID)
		if err != nil {
			return nil, err
		}
		return createOrReplaceStudyProjectCardsResponse{Cards: cards}, nil
	}
}
