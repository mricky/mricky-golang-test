package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code int `json:"code"`
	Status string `json:"status"`
}

func APIResponse (Message string, Code int, Status string, data interface{}) Response {
	meta := Meta {
		Message: Message,
		Code: Code,
		Status: Status,
	}

	jsonResponse := Response{
		Meta: meta,
		Data: data,
	}
	return jsonResponse
}

func FormatError(err error)[]string {
	var errors []string

	// mengubah error menjadi string
	for _, e :=range err.(validator.ValidationErrors){
			errors = append(errors, e.Error())
	}

	return errors;
}