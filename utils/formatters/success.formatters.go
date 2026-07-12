package formatters

import (
	"Coderx/utils/json"
	"net/http"
)

func SuccessResponse(w http.ResponseWriter, status int, message string, data any) error {

	response := map[string]any{}
	response["status"] = status
	response["message"] = message
	response["data"] = data

	return json.CovertTOJSON(w,status,response)

}