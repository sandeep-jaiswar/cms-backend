package responses

import "github.com/gin-gonic/gin"

type Response struct {
	Status string      `json:"status"`          // Status of the response (e.g., success, error)
	Data   interface{} `json:"data,omitempty"`  // Data returned in the response, if any
	Error  string      `json:"error,omitempty"` // Error message, if any
}

func NewSuccessResponse(data interface{}) Response {
	return Response{
		Status: "success",
		Data:   data,
	}
}

func NewErrorResponse(message string) Response {
	return Response{
		Status: "error",
		Error:  message,
	}
}

func WriteResponse(c *gin.Context, httpStatus int, response Response) {
	c.JSON(httpStatus, response)
}
