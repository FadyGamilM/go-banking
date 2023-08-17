package domain

type ErrCode string

const (
	InvalidRequest      ErrCode = "INVALID-REQUEST-ERROR"
	InternalServerError ErrCode = "INTERNAL-SERVER-ERROR"
)

type ErrMsg string

const (
	InvalidRequestErrMsg ErrMsg = "Invalid request data"
	InternalServerErrMsg ErrMsg = "something went wrong, internal server error"
)