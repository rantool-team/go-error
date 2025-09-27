package error

import "encoding/json"

var PREFIX_JSON = ""
var IDENT_JSON = "  "

type JSONResponse struct {
	Status      int    `json:"status"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

func (e Error) getJsonResponse() JSONResponse {
	return JSONResponse{
		Status:      e.GetStatusCode(),
		Message:     e.getErrorMessage(),
		Description: e.getDescription(),
	}
}

func (e Error) GetJSONFormatOfError() string {
	res := e.getJsonResponse()
	resStr, _ := json.MarshalIndent(res, PREFIX_JSON, IDENT_JSON)

	return string(resStr)
}
