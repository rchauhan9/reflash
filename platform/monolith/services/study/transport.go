package study

import (
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/rchauhan9/reflash/monolith/common/middlewares/logging"
	"net/http"
)

func RegisterRoutes(svc Service, router *gin.Engine, logger log.Logger) error {
	endpoints := MakeEndpoints(svc)

	endpoints.CreateStudyProjectEndpoint = logging.Middleware(logger)(endpoints.CreateStudyProjectEndpoint)
	endpoints.CreateOrReplaceStudyProjectCardsEndpoint = logging.Middleware(logger)(endpoints.CreateOrReplaceStudyProjectCardsEndpoint)

	studyGroup := router.Group("/study")
	studyGroup.POST("/projects", CreateStudyProjectHandler(endpoints.CreateStudyProjectEndpoint))
	studyGroup.POST("/projects/cards", CreateOrReplaceStudyProjectCardsHandler(endpoints.CreateOrReplaceStudyProjectCardsEndpoint))

	return nil
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

func decodeCreateOrReplaceStudyProjectCardsRequest(c *gin.Context) (interface{}, error) {
	var req createOrReplaceStudyProjectCardsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, err
	}
	return req, nil
}

func encodeCreateOrReplaceStudyProjectCardsResponse(c *gin.Context, response interface{}) {
	resp := response.(createOrReplaceStudyProjectCardsResponse)
	c.JSON(http.StatusOK, gin.H{"cards": resp.Cards})
}
