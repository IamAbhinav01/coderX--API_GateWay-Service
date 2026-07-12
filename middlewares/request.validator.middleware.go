package middlewares

import (
	"Coderx/dtos"
	"Coderx/utils/formatters"
	"Coderx/utils/json"
	"Coderx/utils/validators"
	"context"
	"net/http"
)

func SingUpRequestValidation(next http.Handler) http.Handler{

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			var payload dtos.SignupRequestDTO
			err:=json.DecodeFROMJSON(r,&payload)

			if err != nil{
				formatters.ErrorResponse(w,http.StatusBadRequest,"Error occured while decoding the json",err)
				return 
			}

			validationErr:=validators.Validate.Struct(payload)

			if validationErr!= nil{
				formatters.ErrorResponse(w,http.StatusBadRequest,"Invalid Request Payload",validationErr)
				return 
			}
			reqContext := r.Context()
			ctx:=context.WithValue(reqContext,"payload",payload)
			
			
			
			next.ServeHTTP(w,r.WithContext(ctx))
			
	})

	
}