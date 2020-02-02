package api

type Response struct {
	Status int8        `json:"staus"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"message"`
}

type ErrorResponse struct {
	Msg  string `json:"message"`
	Type string `json:"type"`
	Code int
}

