package apierr

type ApiErr struct { // can be modified to wrap error and log
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	ErrInvalidParam = ApiErr{Code: 1, Message: "invalid params"}
	ErrInternal     = ApiErr{Code: 2, Message: "internal error"}
)

func (e ApiErr) SetErr(err error) ApiErr {
	e.Message = err.Error()
	return e
}

func (e ApiErr) Error() string {
	return e.Message
}
