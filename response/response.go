package response

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

var (
	ErrorCodeNone      = 0
	ErrorCodeGeneral   = 1
	ErrorStatusNone    = "Success"
	ErrorStatusGeneral = "General Error"
)

// SuccessStatus - just retrun a success status
func SuccessStatus() *Response {
	rs := &Response{}
	rs.Code = ErrorCodeNone
	rs.Status = ErrorStatusNone
	return rs
}

// ErrorStatus - just retrun a general status
func ErrorStatus() *Response {
	rs := &Response{}
	rs.Code = ErrorCodeGeneral
	rs.Status = ErrorStatusGeneral
	return rs
}

// MakeSuccessStatus - todo
func MakeSuccessStatus(status string) *Response {
	rs := &Response{}
	rs.Code = ErrorCodeNone
	rs.Status = status
	return rs
}

// MakeErrorStatus - todo
func MakeErrorStatus(status string) *Response {
	rs := &Response{}
	rs.Code = ErrorCodeGeneral
	rs.Status = status
	return rs
}

// MakeSuccessResponse - todo
func MakeSuccessResponse(data interface{}) *Response {
	rs := &Response{}
	rs.Code = ErrorCodeNone
	rs.Status = ErrorStatusNone
	rs.Data = data
	return rs
}

// MakeErrorResponse - todo
func MakeErrorResponse(data interface{}) *Response {
	rs := &Response{}
	rs.Code = ErrorCodeGeneral
	rs.Status = ErrorStatusGeneral
	rs.Data = data
	return rs
}

// MakeResponse - retrun a Response interface
func MakeResponse(data interface{}, code int, status string) *Response {
	rs := &Response{}
	rs.Code = code
	rs.Status = status
	rs.Data = data
	return rs
}
