package study

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/rchauhan9/reflash/monolith/common/auth"
	"mime/multipart"
)

type listStudyProjectsRequest struct{}

type listStudyProjectsResponse struct {
	StudyProjects []StudyProject `json:"study_projects"`
}

type createStudyProjectRequest struct {
	Name string  `json:"name"`
	Icon *string `json:"icon"`
}

type createStudyProjectResponse struct {
	StudyProject StudyProject `json:"study_project"`
}

type createProjectFileRequest struct {
	StudyProjectID string `json:"study_project_id"`
	Filename       string `json:"filename"`
	File           *multipart.FileHeader
}

type createProjectFileResponse struct {
	ProjectFileID string `json:"project_file_id"`
}

type listCardsRequest struct {
	StudyProjectID string `json:"study_project_id"`
}

type listCardsResponse struct {
	Cards []StudyProjectCard `json:"cards"`
}

type createOrReplaceStudyProjectCardsRequest struct {
	StudyProjectID string `json:"study_project_id"`
}

type createOrReplaceStudyProjectCardsResponse struct {
	Cards []StudyProjectCard `json:"cards"`
}

type Endpoints struct {
	ListStudyProjectsEndpoint                endpoint.Endpoint
	CreateStudyProjectEndpoint               endpoint.Endpoint
	CreateProjectFileEndpoint                endpoint.Endpoint
	ListCardsEndpoint                        endpoint.Endpoint
	CreateOrReplaceStudyProjectCardsEndpoint endpoint.Endpoint
}

func MakeEndpoints(svc Service) Endpoints {
	return Endpoints{
		ListStudyProjectsEndpoint:                MakeListStudyProjectsEndpoint(svc),
		CreateStudyProjectEndpoint:               MakeCreateStudyProjectEndpoint(svc),
		CreateProjectFileEndpoint:                MakeCreateProjectFileEndpoint(svc),
		ListCardsEndpoint:                        MakeListCardsEndpoint(svc),
		CreateOrReplaceStudyProjectCardsEndpoint: MakeCreateOrReplaceStudyProjectCardsEndpoint(svc),
	}
}

func MakeListStudyProjectsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(listStudyProjectsRequest)
		user := ctx.Value("user").(auth.User)
		studyProjects, err := svc.ListStudyProjects(ctx, user.ID)
		if err != nil {
			return nil, err
		}
		return listStudyProjectsResponse{StudyProjects: studyProjects}, nil
	}
}

func MakeCreateStudyProjectEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createStudyProjectRequest)
		user := ctx.Value("user").(auth.User)
		studyProject, err := svc.CreateStudyProject(ctx, user.ID, req.Name, req.Icon)
		if err != nil {
			return nil, err
		}
		return createStudyProjectResponse{StudyProject: studyProject}, nil
	}
}

func MakeCreateProjectFileEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createProjectFileRequest)
		user := ctx.Value("user").(auth.User)
		file, err := svc.CreateProjectFile(ctx, user.ID, req.StudyProjectID, req.Filename, req.File)
		if err != nil {
			return nil, err
		}
		return createProjectFileResponse{ProjectFileID: *file}, nil
	}
}

func MakeListCardsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(listCardsRequest)
		user := ctx.Value("user").(auth.User)
		cards, err := svc.ListCards(ctx, user.ID, req.StudyProjectID)
		if err != nil {
			return nil, err
		}
		return listCardsResponse{Cards: cards}, nil
	}
}

func MakeCreateOrReplaceStudyProjectCardsEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(createOrReplaceStudyProjectCardsRequest)
		user := ctx.Value("user").(auth.User)
		cards, err := svc.CreateOrReplaceStudyProjectCards(ctx, user.ID, req.StudyProjectID)
		if err != nil {
			return nil, err
		}
		return createOrReplaceStudyProjectCardsResponse{Cards: cards}, nil
	}
}
