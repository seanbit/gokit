package foundation

type Error interface {
	error
	Code() int
	Msg() string
}

type cerr struct {
	error
	code 	int
	msg  	string
}

func (this *cerr) Error() string {
	if this.error == nil {
		return ""
	}
	return this.error.Error()
}

func (this *cerr) Code() int {
	return this.code
}

func (this *cerr) Msg() string {
	return this.msg
}

func NewError(err error, code int, msg string) Error {
	return &cerr{
		error: err,
		code:  code,
		msg:  msg,
	}
}
