package response

import (
	"encoding/json"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data  interface{} `json:"data"`
	Error []string    `json:"error"`
}

func ReadResponse(recorder *httptest.ResponseRecorder) (*Response, error) {
	resp := new(Response)
	bytes := recorder.Body.Bytes()
	if len(bytes) == 0 {
		return resp, nil
	}
	err := json.Unmarshal(bytes, resp)
	return resp, err
}

func Respond(ctx *gin.Context, statusCode int, data interface{}) {
	ctx.JSON(statusCode, &Response{Data: data, Error: ctx.Errors.Errors()})
}

func RespondIfError(ctx *gin.Context, statusCode int) bool {
	if allErr := ctx.Errors.Errors(); len(allErr) > 0 {
		ctx.JSON(statusCode, &Response{Error: allErr})
		return true
	}
	return false
}
