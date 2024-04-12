package response

import "github.com/IsThatASkyline/fiberGo/order/internal/application/common"

type ExceptionResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func SetExceptionPayload(response *ExceptionResponse, exc common.CustomException) {
	response.Message = exc.Message
	response.Data = exc.Ctx
}
