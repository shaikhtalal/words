// Copyright © 2022 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"net/http"
	"words/param"
	"words/router"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	cors "github.com/rs/cors/wrapper/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	listenAddr  string
	corsHosts   []string
	logrusEntry *logrus.Entry
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve command runs the http server",
	RunE:  run,
}

func init() {
	serveCmd.Flags().StringVarP(&listenAddr, "listen-addr", "", ":9112", "listen address")
	serveCmd.Flags().StringSliceVar(&corsHosts, "cors-host", []string{"*"}, "CORS allowed hosts, comma separated")
	RootCmd.AddCommand(serveCmd)
}

func run(cmd *cobra.Command, args []string) error {
	logrusEntry = logrus.NewEntry(logrus.StandardLogger())

	p := &param.Param{
		HTTPListenAddr: listenAddr,
		LogrusEntry:    logrusEntry,
		CorsHosts:      corsHosts,
	}
	return serve(p)
}

func serve(p *param.Param) error {
	chErr := make(chan error)

	go func() {
		logrus.Infof("listening on %s", p.HTTPListenAddr)
		chErr <- serveHTTP(p)
	}()
	return <-chErr
}

func serveHTTP(p *param.Param) error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Use(cors.New(cors.Options{
		AllowedOrigins: p.CorsHosts,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders:   []string{"authorization", "*"},
		ExposedHeaders:   []string{"authorization", "*"},
		AllowCredentials: true,
	}))

	param.Inject(r, p)
	router.HandleHTTP(r)
	return errors.Wrap(r.Run(p.HTTPListenAddr), "unable to start server")
}
