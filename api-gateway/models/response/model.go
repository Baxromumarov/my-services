package response

type RegisterSuccessModel struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// SuccessModel ...
type SuccessModel struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorModel ...
type ErrorModel struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

// FileUploadedModel ...
type FileUploadedModel struct {
	Filename string `json:"filename"`
	URL      string `json:"file_url"`
}

type SocketResponse struct {
	StatusCode    int        `json:"status_code"`
	Error         ErrorModel `json:"error"`
	Data          string     `json:"data"`
	Action        string     `json:"action"`
	CorrelationID string     `json:"correlation_id"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ResponseOK struct {
	Message string `json:"message"`
}

type Response struct {
	SessionID     string      `json:"session_id,omitempty"  swaggerignore:"true"`
	StatusCode    int         `json:"status_code,omitempty"`
	ID            string      `json:"id,omitempty"`
	Error         Error       `json:"error,omitempty"`
	Data          interface{} `json:"data,omitempty"`
	NoResponse    bool        `json:"-"`
	Action        string      `json:"Action"`
	CorrelationID string      `json:"correlation_id" swaggerignore:"true"`
}

type ResponseError struct {
	Error Error `json:"error"`
}

type CreateResponse struct {
	ID string `json:"id"`
}
