package middleware

import (
	"github.com/schidstorm/ffmpeg-jobs/api/lib"
	"net/http"
)

type Cors struct {


}

func (c Cors) Handle() lib.MiddlewareHandler {
	return func(writer http.ResponseWriter, request *http.Request, nextHandler lib.MiddlewareNextHandler) error {
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "*")
		return nextHandler(writer, request)
	}
}
