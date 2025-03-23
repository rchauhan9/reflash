package hello

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"github.com/rchauhan9/reflash/monolith/common/middlewares/logging"
	"github.com/rchauhan9/reflash/monolith/utils"
	"net/http"
)

func RegisterRoutes(appContext *utils.AppContext) {
	rep := NewRepository()
	svc := NewService(rep)
	endpoints := MakeEndpoints(svc)

	helloLogger := log.WithSuffix(appContext.Logger, "svc", "hello")

	endpoints.SayHello = logging.Middleware(helloLogger)(endpoints.SayHello)

	helloGroup := appContext.Router.Group("/hello")
	helloGroup.GET("/greeting/:name", SayHelloHandler(endpoints.SayHello))

}

func SayHelloHandler(endpoint endpoint.Endpoint) func(c *gin.Context) {
	return func(c *gin.Context) {
		req, err := decodeSayHelloRequest(c)
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
		encodeSayHelloResponse(c, response)
	}
}

func decodeSayHelloRequest(c *gin.Context) (interface{}, error) {
	name := c.Param("name")
	if name == "" {
		return nil, errors.New("missing path parameter: name")
	}
	req := sayHelloRequest{
		Name: name,
	}
	return req, nil
}

func encodeSayHelloResponse(c *gin.Context, response interface{}) {
	resp := response.(sayHelloResponse)
	c.JSON(http.StatusOK, gin.H{"greeting": resp.Greeting})
}
