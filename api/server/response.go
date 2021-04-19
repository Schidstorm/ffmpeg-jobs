package server

type ResponseOk struct {
	Success bool
	Data    interface{}
}

func CreateResponseOk(data interface{}) ResponseOk {
	return ResponseOk{
		Success: true,
		Data:    data,
	}
}

type ResponseError struct {
	Success bool
}

func CreateResponseError() ResponseError {
	return ResponseError{Success: false}
}
