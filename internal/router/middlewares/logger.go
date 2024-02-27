package middlewares

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-multierror"
	"github.com/sirupsen/logrus"
)

var (
	logger = &logrus.Logger{
		Out: os.Stdout,
		Formatter: &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05.000-07:00",
		},
		Hooks:        make(logrus.LevelHooks),
		Level:        logrus.InfoLevel,
		ExitFunc:     os.Exit,
		ReportCaller: false,
	}
	logResponseBody = true
	emptyJSONObject = []byte("{}")
)

const requestErrorKey = "requestError"

func getRequestError(c *gin.Context) error {
	value, ok := c.Get(requestErrorKey)
	if !ok {
		return nil
	}
	return value.(error)
}

func SaveRequestError(c *gin.Context, err error) {
	errs := getRequestError(c)
	if errs == nil {
		errs = err
	} else {
		errs = multierror.Append(errs, err)
	}
	c.Set(requestErrorKey, errs)
}

func RequestLogger(c *gin.Context) {
	t := time.Now()
	reqBody := getRequestRawBody(c)
	c.Next()
	logRequest(c, c.Request, t, reqBody, logResponseBody)
}

func logRequest(c *gin.Context, req *http.Request, start time.Time, reqBody []byte, logResponseBody bool) {
	respWriter := c.Writer

	executionTime := time.Now().Sub(start).Milliseconds()
	hostname, _ := os.Hostname()

	requestFields := logrus.Fields{
		"method":     req.Method,
		"uri":        req.RequestURI,
		"remoteAddr": req.RemoteAddr,
		"proto":      req.Proto,
		"headers":    req.Header,
		"form":       json.RawMessage(getRequestBody(c, reqBody)),
	}

	responseFields := logrus.Fields{
		"headers":    respWriter.Header(),
		"statusCode": respWriter.Status(),
	}
	if logResponseBody {
		responseFields["body"] = json.RawMessage(getResponseBody(c))
	}

	enviromentFields := logrus.Fields{
		"hostname": hostname,
		"lang":     runtime.Version(),
		"pid":      os.Getpid(),
	}

	fields := logrus.Fields{
		"request":       requestFields,
		"response":      responseFields,
		"enviroment":    enviromentFields,
		"user":          GetCurrentUser(c),
		"executionTime": executionTime,
	}

	reqErr := getRequestError(c)
	if reqErr != nil {
		fields["error"] = reqErr
	}

	logger.WithFields(fields).Info("api request")
}

const responseBodyKey = "responseBody"

func SaveResponse(c *gin.Context, resp interface{}) {
	if !logResponseBody {
		return
	}
	data, err := json.Marshal(resp)
	if err != nil {
		data = emptyJSONObject
	}
	c.Set(responseBodyKey, data)
}

func getRequestRawBody(c *gin.Context) []byte {
	if c.Request.Body == nil {
		return nil
	}
	contentType := strings.ToLower(c.ContentType())
	ok := strings.Contains(contentType, "application/json")
	if !ok {
		return nil
	}
	data, _ := io.ReadAll(c.Request.Body)
	c.Request.Body = io.NopCloser(bytes.NewBuffer(data))
	return data
}

func getRequestBody(c *gin.Context, rawBody []byte) []byte {
	m := make(map[string]string)
	if len(rawBody) > 0 {
		if json.Valid(rawBody) {
			return rawBody
		}
		m["body"] = string(rawBody)
	}
	form := c.Request.Form
	for k, v := range form {
		m[k] = strings.Join(v, ",")
	}
	data, _ := json.Marshal(m)
	if len(data) == 0 {
		return emptyJSONObject
	}
	return data
}

func getResponseBody(c *gin.Context) []byte {
	value, ok := c.Get(responseBodyKey)
	if !ok {
		return emptyJSONObject
	}
	data := value.([]byte)
	if len(data) == 0 {
		return emptyJSONObject
	}
	if json.Valid(data) {
		return data
	}
	return emptyJSONObject
}
