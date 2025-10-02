package errorstruct

import (
	"encoding/json"

	"github.com/rantool-team/go-error/context"
)

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
		Description: e.GetDescription(),
	}
}

func (e Error) GetJSONFormatOfError() string {
	res := e.getJsonResponse()
	resStr, _ := json.MarshalIndent(res, PREFIX_JSON, IDENT_JSON)

	return string(resStr)
}

func (e Error) GetJsonLocal() string {
	res := e.Context.Local
	resStr, _ := json.MarshalIndent(res, PREFIX_JSON, IDENT_JSON)

	return string(resStr)
}

func (e Error) GetLocal() context.Local {
	return e.Context.Local
}
