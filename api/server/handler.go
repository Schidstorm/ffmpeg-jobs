package server

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/schidstorm/ffmpeg-jobs/api/lib"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

type HttpHandler struct {
	globalRouter *httprouter.Router
}

func NewHttpHandler() *HttpHandler {
	return &HttpHandler{globalRouter: httprouter.New()}
}

func (h *HttpHandler) Handler() http.Handler {
	return h.globalRouter
}

func (h *HttpHandler) AddController(controller lib.Controller) {

	name := controller.Name()

	listHandler := controller.ListHandler()
	if listHandler != nil {
		h.globalRouter.GET("/"+name, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			data, err := listHandler(request.URL.Query())
			if err != nil {
				handleError(writer, err)
				return
			}

			handleSuccess(writer, data)
		})
	}

	getHandler := controller.GetHandler()
	if getHandler != nil {
		h.globalRouter.GET("/"+name+"/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			idString := params.ByName("id")
			id, err := strconv.ParseInt(idString, 10, 64)
			if err != nil {
				handleError(writer, err)
				return
			}

			data, err := getHandler(id, request.URL.Query())
			if err != nil {
				handleError(writer, err)
				return
			}

			handleSuccess(writer, data)
		})
	}

	postHandler, postDataStructure := controller.PostHandler()
	if postHandler != nil {
		h.globalRouter.POST("/"+name, func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			if postDataStructure != nil {
				bodyData, err := ioutil.ReadAll(request.Body)
				if err != nil {
					handleError(writer, err)
					return
				}

				clear(postDataStructure)
				err = json.Unmarshal(bodyData, postDataStructure)
				if err != nil {
					handleError(writer, err)
					return
				}
			}

			data, err := postHandler(request.URL.Query(), postDataStructure)
			if err != nil {
				handleError(writer, err)
				return
			}

			handleSuccess(writer, data)
		})
	}

	putHandler, putDataStructure := controller.PutHandler()
	if putHandler != nil {
		h.globalRouter.PUT("/"+name+"/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
			idString := params.ByName("id")
			id, err := strconv.ParseInt(idString, 10, 64)
			if err != nil {
				handleError(writer, err)
				return
			}

			if putDataStructure != nil {
				bodyData, err := ioutil.ReadAll(request.Body)
				if err != nil {
					handleError(writer, err)
					return
				}

				clear(putDataStructure)
				err = json.Unmarshal(bodyData, putDataStructure)
				if err != nil {
					handleError(writer, err)
					return
				}
			}

			data, err := putHandler(id, request.URL.Query(), putDataStructure)
			if err != nil {
				handleError(writer, err)
				return
			}

			handleSuccess(writer, data)
		})
	}
}

func handleError(w http.ResponseWriter, err error) {
	logrus.Error(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	data, err := json.Marshal(CreateResponseError())
	if err != nil {
		logrus.Error(err)
	}
	_, err = w.Write(data)
	if err != nil {
		logrus.Error(err)
	}
}

func handleSuccess(w http.ResponseWriter, data interface{}) {
	jsonData, err := json.Marshal(CreateResponseOk(data))
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	_, err = w.Write(jsonData)
	if err != nil {
		logrus.Error(err)
	}
}

func clear(v interface{}) {
	if v == nil {
		return
	}

	if reflect.TypeOf(v).Kind() == reflect.Ptr {
		p := reflect.ValueOf(v).Elem()
		p.Set(reflect.Zero(p.Type()))
	}
}
