package utils

const (
	//StatusFail represents status fail of API eventhough the code is 200 (StatusOK), it might be fail in terms of API.
	StatusFail string = "fail"
	// StatusSuccess implies that API successfully responded
	StatusSuccess string = "success"
)

const (
	// MessageOk is a responses message of status success
	MessageOk string = "OK"
)

// BaseSuccessResponse is an object for the application success response
type BaseSuccessResponse struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"OK"`
	Data    interface{} `json:"data"`
}

// NewSuccessResponse creates a success response object
func NewSuccessResponse(data interface{}) BaseSuccessResponse {
	r := BaseSuccessResponse{
		Status:  StatusSuccess,
		Message: MessageOk,
		Data:    data,
	}

	return r
}

// ErrorResponse is an object for the application error response
type ErrorResponse struct {
	Status  string `json:"status" example:"fail"`
	Message string `json:"message" example:"Error message will be show here"`
}

// NewErrorResponse creates an error response
func NewErrorResponse(message string) ErrorResponse {
	errorResponse := ErrorResponse{}
	errorResponse.Status = StatusFail
	errorResponse.Message = message
	return errorResponse
}
