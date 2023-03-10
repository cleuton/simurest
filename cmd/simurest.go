/*
Copyright 2023 Cleuton Sampaio de Melo Junior

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"encoding/json"
	"log"
	"net/http"

	flag "github.com/spf13/pflag"
)

/*
	"fmt"
	"net/http"
	"flag"
	"github.com/gorilla/mux"
*/

const disclaimer = `
SimuRest - Service simulator - V. 0.1
Usage: 
	simurest [
		[--quiet | -q <do not print this message (default true)>]
		[--method | -m <http method (default=GET): GET | POST | HEAD | DELETE | PUT>]
		[--port | -p <tcp port (default=8080)> ]
		[--uri | -u <http URI (default="/")]
		[--status | -s <http Status of the response (default=200)]
		[--body | -b <http response body (default: '{"status": "ok"}')]
	]
`

var (
	flagQuiet  *bool
	flagMethod *string
	flagPort   *int
	flagUri    *string
	flagStatus *int
	flagBody   *string
)

func init() {
	flagQuiet = flag.BoolP("quiet", "q", false, "to supress the disclaimer header")
	flagMethod = flag.StringP("method", "m", "GET", "http method to accept")
	flagPort = flag.IntP("port", "p", 8080, "tcp port to listen")
	flagUri = flag.StringP("uri", "u", "/", "uri to handle")
	flagStatus = flag.IntP("status", "s", 200, "http status to return")
	flagBody = flag.StringP("body", "b", `{"status": "ok"}`, "response body to return")
}

func setupLog() log.Logger {
	logger := log.Default()
	return *logger
}

func WriteResponse(status int, body interface{}, w http.ResponseWriter) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	if body != nil {
		payload, _ := json.Marshal(body)
		w.Write(payload)
	}
}

func dynHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			WriteResponse(200, "OK", w)
		},
	)
}

func main() {
	logger := setupLog()
	logger.Println(disclaimer)
	flag.Parse()
	http.Handle("/", dynHandler())

	http.ListenAndServe(":8080", nil)
}
