package fields
///EXAMPLE CUSTOM FIELD
import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Runtime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r *Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJSONValue := strconv.Quote(jsonValue)
	return []byte(quotedJSONValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error{
	//Incoming JSON value has to be string format "<runtime> mins", therefore remove double quotes
	unquotedJsonValue, err := strconv.Unquote(string(jsonValue))
	if err != nil{
		//return Invalid format if unquotable
		return ErrInvalidRuntimeFormat
	}

	//Split string to isolate the part containg the number
	parts := strings.Split(unquotedJsonValue, " ")
	if len(parts) != 2 || parts [1] != "mins"{
		return ErrInvalidRuntimeFormat
	}

	//Parse number into a base 10, int3
	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil{
		return ErrInvalidRuntimeFormat
	}

	//Convert int32 to a Runtime type and assign to the receiver.
	// use the * operator to deference the receiver (which is a pointer to a Runtime
	// type) in order to set the underlying value of the pointer.
	
	*r = Runtime(i)

	return nil
}
