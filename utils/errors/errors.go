package errors

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(message string, status int, error string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  status,
		Error:   error,
	}
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  400,
		Error:   "bad_request",
	}
}
