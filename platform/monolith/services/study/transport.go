package study

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/rchauhan9/reflash/monolith/common/middlewares/logging"
	"net/http"
)

func RegisterRoutes(svc Service, router *gin.Engine, logger log.Logger) error {
	endpoints := MakeEndpoints(svc)

	endpoints.ListStudyProjectsEndpoint = logging.Middleware(logger)(endpoints.ListStudyProjectsEndpoint)
	endpoints.CreateStudyProjectEndpoint = logging.Middleware(logger)(endpoints.CreateStudyProjectEndpoint)
	endpoints.ListCardsEndpoint = logging.Middleware(logger)(endpoints.ListCardsEndpoint)
	endpoints.CreateOrReplaceStudyProjectCardsEndpoint = logging.Middleware(logger)(endpoints.CreateOrReplaceStudyProjectCardsEndpoint)

	studyGroup := router.Group("/study")
	studyGroup.GET("/projects", ListStudyProjectsHandler(endpoints.ListStudyProjectsEndpoint))
	studyGroup.POST("/projects", CreateStudyProjectHandler(endpoints.CreateStudyProjectEndpoint))
	studyGroup.GET("/projects/:projectID/cards", ListCardsHandler(endpoints.ListCardsEndpoint))
	studyGroup.POST("/projects/:projectID/cards", CreateOrReplaceStudyProjectCardsHandler(endpoints.CreateOrReplaceStudyProjectCardsEndpoint))

	return nil
}

func ListStudyProjectsHandler(endpoint endpoint.Endpoint) func(c *gin.Context) {
	return func(c *gin.Context) {
		req, err := decodeListStudyProjectsRequest(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the endpoint with the request.
		response, err := endpoint(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Encode and send the response.
		encodeListStudyProjectsResponse(c, response)
	}
}

func CreateStudyProjectHandler(endpoint endpoint.Endpoint) func(c *gin.Context) {
	return func(c *gin.Context) {
		req, err := decodeCreateStudyProjectRequest(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the endpoint with the request.
		response, err := endpoint(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Encode and send the response.
		encodeCreateStudyProjectResponse(c, response)
	}
}

func ListCardsHandler(endpoint endpoint.Endpoint) func(c *gin.Context) {
	return func(c *gin.Context) {
		req, err := decodeListCardsRequest(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the endpoint with the request.
		response, err := endpoint(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Encode and send the response.
		encodeListCardsResponse(c, response)
	}
}

func CreateOrReplaceStudyProjectCardsHandler(endpoint endpoint.Endpoint) func(c *gin.Context) {
	return func(c *gin.Context) {
		req, err := decodeCreateOrReplaceStudyProjectCardsRequest(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Call the endpoint with the request.
		response, err := endpoint(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Encode and send the response.
		encodeCreateOrReplaceStudyProjectCardsResponse(c, response)
	}
}

func decodeListStudyProjectsRequest(c *gin.Context) (interface{}, error) {
	var req listStudyProjectsRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeListStudyProjectsResponse(c *gin.Context, response interface{}) {
	resp := response.(listStudyProjectsResponse)
	c.JSON(http.StatusOK, gin.H{"study_projects": resp.StudyProjects})
}

func decodeCreateStudyProjectRequest(c *gin.Context) (interface{}, error) {
	var req createStudyProjectRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateStudyProjectResponse(c *gin.Context, response interface{}) {
	resp := response.(createStudyProjectResponse)
	c.JSON(http.StatusOK, gin.H{"study_project": resp.StudyProject})
}

func decodeListCardsRequest(c *gin.Context) (interface{}, error) {
	var req listCardsRequest

	// Extract the project ID from the URL parameters
	projectID := c.Param("projectID")
	if projectID == "" {
		return nil, fmt.Errorf("projectID is required")
	}
	req.StudyProjectID = projectID

	return req, nil
}

func encodeListCardsResponse(c *gin.Context, response interface{}) {
	resp := response.(listCardsResponse)
	c.JSON(http.StatusOK, gin.H{"cards": resp.Cards})
}

func decodeCreateOrReplaceStudyProjectCardsRequest(c *gin.Context) (interface{}, error) {
	var req createOrReplaceStudyProjectCardsRequest

	// Extract the project ID from the URL parameters
	projectID := c.Param("projectID")
	if projectID == "" {
		return nil, fmt.Errorf("projectID is required")
	}
	req.StudyProjectID = projectID

	return req, nil
}

func encodeCreateOrReplaceStudyProjectCardsResponse(c *gin.Context, response interface{}) {
	resp := response.(createOrReplaceStudyProjectCardsResponse)
	c.JSON(http.StatusOK, gin.H{"cards": resp.Cards})
}
