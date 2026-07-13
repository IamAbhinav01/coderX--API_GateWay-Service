package json

import (
	"encoding/json"
	"net/http"
)

func CovertTOJSON(w http.ResponseWriter,status int, data any) error{

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	
	return json.NewEncoder(w).Encode(data)
}


func DecodeFROMJSON(r *http.Request,result any) error{

	decoder:=json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	
	return decoder.Decode(result)
}


func Marshall(data any) ([]byte,error){
	return json.Marshal(data)
}

func UnMarshall(data []byte,v any) error{
	return json.Unmarshal(data,v)
}