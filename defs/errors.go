package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
	HttpSC int
	Error  Err
}

var (
	ErrorRequestParamParseFailed = ErrResponse{HttpSC: 400, Error: Err{Error: "Request param is not correct", ErrorCode: "002"}}
	ErrorResouceNotFound         = ErrResponse{HttpSC: 404, Error: Err{Error: "Resource not found", ErrorCode: "404"}}
	ErrorInternalFaults          = ErrResponse{HttpSC: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
	ErrorTooManyRequests         = ErrResponse{HttpSC: 429, Error: Err{Error: "Too many Request", ErrorCode: "005"}}
)
