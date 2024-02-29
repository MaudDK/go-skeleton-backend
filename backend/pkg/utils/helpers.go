package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
)

type Envelope map[string]interface{}


func ReadJSON(w http.ResponseWriter, r *http.Request, dst any) error{
	//Limit Request Size to 1MB.
	maxBytes := 1_048_576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))
	
	//Intialize decoder
	dec :=json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	
	//Decode Request
	err := dec.Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError 
		var invalidUnmarshalError *json.InvalidUnmarshalError
		var maxBytesError *http.MaxBytesError

		switch{
			//Uses errors.As() function to check if the error is a syntax error
			case errors.As(err, &syntaxError):
				return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)
			
			//Uses erros.Is() function to check if EOF
			case errors.Is(err, io.ErrUnexpectedEOF):
				return errors.New("body contains badly-formed JSON")
			
			//Checks if a unmarshalType error. Occurs when JSON value is wrong type for target
			case errors.As(err, &unmarshalTypeError):
				//If error relates to a field
				if unmarshalTypeError.Field != ""{
					return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
				}
				return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)
			
			//Checks if EOF (request body is empty)
			case errors.Is(err, io.EOF):
				return errors.New("body must not be empty")
			
			//Decode will return unknown field, extract unknown field name and return it
			case strings.HasPrefix(err.Error(), "json: unknown field"):
				fieldName := strings.TrimPrefix(err.Error(), "json: unknown field")
				return fmt.Errorf("body contains unknown key %s", fieldName)
			
			//Check if exceeds size limit of 1MB and returns and error msg
			case errors.As(err, &maxBytesError):
				return fmt.Errorf("body must not be larger than %d bytes", maxBytesError.Limit)

			//Invalid UnmarshalError is returned when we pass a non nil pointer to decode then we catch and panic
			case errors.As(err, &invalidUnmarshalError):
				panic(err)
			
			//For everything else return err as is
			default:
				return err
		}
	}

	// Call Decode() again, using a pointer to an empty anonymous struct as the
	// destination. If the request body only contained a single JSON value this will
	// return an io.EOF error. So if we get anything else, we know that there is
	// additional data in the request body and we return our own custom error message.
	err = dec.Decode(&struct{}{})
	if err != io.EOF{
		return errors.New("body must only contain a single JSON value")
	}
	return nil
}

func WriteJson(w http.ResponseWriter, status int, data Envelope, headers http.Header) error{
	//Encode data to JSON, returning error if error exists
	js, err := json.Marshal(data)
	if err != nil{
		return err
	}

	js = append(js, '\n')

	//Add each header to the http.ResponseWriter header map
	for key, value := range headers{
		w.Header()[key] = value
	}

	//Add the json content type header
	w.Header().Set("Content-Type", "application/json")
	//Write status code
	w.WriteHeader(status)
	//Write JSON response
	w.Write(js)

	return nil
}

//Retrieve URL parameter from current request context
func ReadParam(param string, r *http.Request)(int64, error){
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName(param), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

	return id, nil

}
