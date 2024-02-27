package cmd

import (
	"context"
	. "taylor-ai-server/internal/domain"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

type testingT struct {
}

func (t testingT) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (t testingT) FailNow() {
	panic("fail now")
}

func NewTestCommand() *cobra.Command {
	flags := &TestOptions{}
	cmd := &cobra.Command{
		Use: "test",
		Run: func(cmd *cobra.Command, args []string) {
			checkError(flags.Complete(cmd, args))
			checkError(flags.Run(cmd.Context()))
		},
	}
	flags.AddFlags(cmd)
	return cmd
}

type TestOptions struct {
	File    string
	BaseURL string
	client  *http.Client
	T       testingT
}

func NewTestOptions() (*TestOptions, error) {
	return &TestOptions{}, nil
}

func (o *TestOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&o.File, "config-file", "c", "", "Configuration file path")
	cmd.Flags().StringVarP(&o.BaseURL, "base-url", "b", "http://localhost:8000", "Base URL of the API")
}

func (o *TestOptions) Complete(cmd *cobra.Command, args []string) error {
	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	o.client = &http.Client{Jar: cookieJar}
	o.T = testingT{}

	return nil
}

func (o *TestOptions) Run(ctx context.Context) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %+v", r)
		}
	}()

	o.login(ctx)
	o.profile(ctx)

	return nil
}

func (o *TestOptions) request(ctx context.Context, method, path string, request url.Values, response interface{}) error {
	uri := o.BaseURL + path
	var reqBody io.Reader = http.NoBody

	if method == http.MethodGet || method == http.MethodDelete {
		uri = uri + "?" + request.Encode()
	} else {
		reqBody = strings.NewReader(request.Encode())
	}

	r, err := http.NewRequestWithContext(ctx, method, uri, reqBody)
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", "drrr.ai")

	resp, err := o.client.Do(r)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"method":   method,
		"path":     path,
		"request":  request,
		"response": string(respData),
	}).Debug("test request")

	err = json.Unmarshal(respData, response)
	return err
}

func (o *TestOptions) login(ctx context.Context) {
	request := url.Values{}
	request.Set("username", "test")
	request.Set("email", "test@drrr.ai")
	request.Set("deviceId", "506df560-5349-4bc1-b062-0405bcff4be5")

	response := struct {
		User User `json:"user"`
	}{}

	err := o.request(ctx, http.MethodPost, "/api/login", request, &response)
	require.NoError(o.T, err)
}

func (o *TestOptions) sendMail(ctx context.Context) {
	request := url.Values{}
	request.Set("username", "test")
	request.Set("email", "test@drrr.ai")
	request.Set("deviceId", "506df560-5349-4bc1-b062-0405bcff4be5")
	request.Set("type", "verifyLogin")

	response := struct{}{}

	err := o.request(ctx, http.MethodPost, "/api/send/mail", request, &response)
	require.NoError(o.T, err)
}

func (o *TestOptions) profile(ctx context.Context) {
	request := url.Values{}
	response := struct {
		User User `json:"user"`
	}{}
	err := o.request(ctx, http.MethodGet, "/api/profile", request, &response)
	require.NoError(o.T, err)
}
