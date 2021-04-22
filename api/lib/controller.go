package lib

import (
	"net/url"
)

type GetHandlerFunc = func(id int64, values url.Values) (interface{}, error)
type PostHandlerFunc = func(values url.Values, data interface{}) (interface{}, error)
type PutHandlerFunc = func(id int64, values url.Values, data interface{}) (interface{}, error)
type ListHandlerFunc = func(values url.Values) (interface{}, error)
type DeleteHandlerFunc = func(id int64, values url.Values) (interface{}, error)

type Controller interface {
	Name() string
	GetHandler() GetHandlerFunc
	PutHandler() (PutHandlerFunc, interface{})
	PostHandler() (PostHandlerFunc, interface{})
	ListHandler() ListHandlerFunc
	DeleteHandler() DeleteHandlerFunc
}
