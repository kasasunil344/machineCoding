package cerror

type CError interface {
	GetErrorCode() string
}

type Error struct {
	code string
}

func New(code string) CError {
	return Error{
		code: code,
	}
}

func (e Error) GetErrorCode() string {
	return e.code
}
