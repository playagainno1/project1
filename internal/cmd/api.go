package cmd

import (
	"context"
	"taylor-ai-server/internal/api"
	"taylor-ai-server/internal/config"
	"taylor-ai-server/internal/pkg/connection"
	"taylor-ai-server/internal/pkg/migration"
	"taylor-ai-server/internal/router"

	"github.com/braintree/manners"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewAPICommand() *cobra.Command {
	flags := &APIOptions{}
	cmd := &cobra.Command{
		Use: "api",
		Run: func(cmd *cobra.Command, args []string) {
			checkError(flags.Complete(cmd, args))
			checkError(flags.Run(cmd.Context()))
		},
	}
	flags.AddFlags(cmd)
	return cmd
}

type APIOptions struct {
	File string
}

func NewAPIOptions() (*APIOptions, error) {
	return &APIOptions{}, nil
}

func (o *APIOptions) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&o.File, "config-file", "c", "", "Configuration file path")
}

func (o *APIOptions) Complete(cmd *cobra.Command, args []string) error {
	return nil
}

func (o *APIOptions) Run(ctx context.Context) error {
	err := config.LoadConfig(o.File)
	if err != nil {
		return errors.WithStack(err)
	}

	c := config.Config

	connection.Init()
	err = migration.Migrate(connection.DB())
	if err != nil {
		return errors.WithStack(err)
	}

	h := api.NewHTTP()
	r := router.NewRouter(h).Router()
	done := ctx.Done()
	o.serve(r, c.Addr, done)

	return nil
}

func (o *APIOptions) serve(r *gin.Engine, addr string, done <-chan struct{}) {
	logrus.Infof("listening on %s", addr)

	go func() {
		err := manners.ListenAndServe(addr, r)
		checkError(err)
	}()
	<-done
	manners.Close()
}
