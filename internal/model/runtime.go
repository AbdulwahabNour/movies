package model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)


type Runtime int32

var ErrInvalidRuntimeFormat = errors.New("invalid runtime format")

func (r *Runtime) MarshalJSON()([]byte, error){
     
    jsonVal := fmt.Sprintf("%d mins", r)
    quotedJsonVal := strconv.Quote(jsonVal) 
    return []byte(quotedJsonVal), nil
}


func (r *Runtime) UnmarshalJSON(jsonVal []byte) error{
  
    //"50 mins"
    unquotedJson, err := strconv.Unquote(string(jsonVal))

    if err != nil{
        return ErrInvalidRuntimeFormat
    }
    parts := strings.Split(unquotedJson, " ")

    if len(parts[0]) < 1 ||  len(parts[0]) > 3 || parts[1] != "mins"{
        return ErrInvalidRuntimeFormat
    }

    val, err := strconv.ParseInt(parts[0], 10, 32)
    if err != nil{
        return ErrInvalidRuntimeFormat
    }

    *r = Runtime(val)

    return nil
}