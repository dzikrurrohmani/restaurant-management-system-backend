package response

const (
	SuccessCode         = "00"
	SuccessMessage      = "Success"
	DefaultErrorCode    = "XX"
	DefaultErrorMessage = "Something went wrong"
)

type AppHttpResponse interface {
	Send()
}
