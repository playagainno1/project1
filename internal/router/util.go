package router

import (
	"context"
	"net/http"
	"reflect"

	"taylor-ai-server/internal/domain"
	"taylor-ai-server/internal/router/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type panicWriter struct {
}

func (w panicWriter) Write(p []byte) (n int, err error) {
	logrus.WithField("panic", string(p)).Error("request panic")
	return len(p), nil
}

func WrapHandler[R any, W any](fn func(context.Context, *R) (*W, error)) func(*gin.Context) {
	return func(c *gin.Context) {
		r := new(R)
		if err := c.ShouldBind(r); err != nil {
			responseError(c, err, domain.ErrBadRequest)
			return
		}
		if parser, ok := any(r).(RequestParser); ok {
			if err := parser.Parse(c); err != nil {
				responseError(c, err, domain.ErrBadRequest)
				return
			}
		}

		resp, err := fn(c, r)
		if err != nil {
			responseError(c, err, domain.ErrInternalError)
			return
		}

		if w, ok := any(resp).(ResponseWriter); ok {
			w.Write(c, c.Writer, c.Request)
			return
		}

		c.JSON(http.StatusOK, resp)
		middlewares.SaveResponse(c, resp)
	}
}

func responseError(c *gin.Context, err error, defaultErr domain.Error) {
	middlewares.SaveRequestError(c, err)

	var finalErr domain.Error
	if e, ok := errors.Cause(err).(domain.Error); ok {
		finalErr = e
	} else {
		finalErr = defaultErr
	}
	c.JSON(finalErr.Status(), finalErr)
}

type RequestParser interface {
	Parse(ctx context.Context) error
}

type ResponseWriter interface {
	Write(ctx context.Context, w http.ResponseWriter, r *http.Request)
}

func wrapGoHandler(f http.HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		f(c.Writer, c.Request)
	}
}

func Call(fn interface{}, params ...interface{}) ([]reflect.Value, error) {
	f := reflect.ValueOf(fn)
	if f.Kind() != reflect.Func {
		return nil, errors.New("fn is not a function")
	}
	if len(params) != f.Type().NumIn() {
		return nil, errors.New("parameters of function are error")
	}

	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result := f.Call(in)
	return result, nil
}

func Call2(fn interface{}, params ...interface{}) (interface{}, error) {
	result, err := Call(fn, params...)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if len(result) != 2 {
		return nil, errors.New("returns of function are error")
	}
	if result[1].IsNil() {
		return result[0].Interface(), nil
	}
	err, ok := result[1].Interface().(error)
	if !ok {
		return nil, errors.New("should return a error for function")
	}
	return nil, errors.WithStack(err)
}

func NewHandlerRequest(fn interface{}) (interface{}, error) {
	f := reflect.ValueOf(fn)
	if f.Kind() != reflect.Func {
		return nil, errors.New("fn is not a function")
	}
	if f.Type().NumIn() != 2 {
		return nil, errors.New("parameters number should be two")
	}
	reqPtr := f.Type().In(1)
	if reqPtr.Kind() != reflect.Ptr {
		return nil, errors.New("request parameter must be a pointer")
	}
	req := reflect.New(reqPtr.Elem()).Interface()
	return req, nil
}
