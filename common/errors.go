package common

type AppError struct {
	Msg string
}

func (e AppError) Error() string {
	return e.Msg
}

var (
	NotFound            AppError = AppError{Msg: "not found"}
	BadRequest          AppError = AppError{Msg: "bad request"}
	InternalServerError AppError = AppError{Msg: "internal server error"}
)
