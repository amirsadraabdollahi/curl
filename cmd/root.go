/*
Copyright © 2024 Amirsadra Abdollahi
*/
package cmd

import (
	"context"
	"github.com/amirsadraabdollahi/curl/internal/request"
	"github.com/amirsadraabdollahi/curl/util"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/spf13/cobra"
)

type InvalidArgsError struct {
	msg string
}

func (i InvalidArgsError) Error() string {
	return i.msg
}

var requester request.Requester = request.NewHttpRequester(util.NewBasePrinter())

var rootCmd = &cobra.Command{
	Use:   "curl",
	Short: "A curl application in Go",
	Run:   rootCmdFunc,
}

func rootCmdFunc(cmd *cobra.Command, args []string) {
	if len(args) > 1 {
		log.WithField("args", len(args)).WithError(InvalidArgsError{msg: "invalid number of arguments"}).Errorln()
	}
	url := args[0]
	ctx := context.Background()
	method := cmd.Flag("method").Value.String()
	if verbose := cmd.Flag("verbose").Changed; verbose {
		requester.SetPrinter(util.NewVerbosePrinterDecorator(util.NewBasePrinter()))
	}
	switch method {
	case "GET":
		_ = requester.Get(&ctx, url)
	}
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("method", "X", "GET", "specifying request method")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "verbose")
}
