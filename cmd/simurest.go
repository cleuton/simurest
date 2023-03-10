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
	"fmt"
	"log"
	"net/http"
	"strings"

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

	uri can be exact or finish with wildcard (*):
	-u /api/user/1 : Only accepts exact uri
	-u "/api/user*" : accept any uri that begins with "/api/user" (if using * don't forget to enclose in double quotes)
`

var (
	flagQuiet  *bool
	flagMethod *string
	flagPort   *int
	flagUri    *string
	flagStatus *int
	flagBody   *string
)

func setFlags() {
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

type JSONString string

func (j JSONString) MarshalJSON() ([]byte, error) {
	return []byte(j), nil
}

func WriteResponse(status int, body string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if len(body) > 0 {
		content, _ := json.Marshal(JSONString(body))
		w.Write(content)
	}
}

func checkUrl(uri string) bool {
	if (*flagUri)[len((*flagUri))-1] == '*' {
		if strings.Contains((*flagUri), uri) {
			return true
		} else if uri == *flagUri {
			return true
		}
	}
	return false
}

func checkMethod(method string) bool {
	if method == *flagMethod {
		return true
	}
	return false
}

func dynHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			logger.Println(
				fmt.Sprintf("%v -> %v %v%v", strings.Split(req.Proto, "/")[0], req.Method, req.Host, req.URL.Path))
			if !checkUrl(req.URL.Path) {
				WriteResponse(404, "Not found", w)
				return
			}
			if !checkMethod(req.Method) {
				WriteResponse(405, "Method not allowed", w)
				return
			}
			if len(*flagBody) == 0 || *flagStatus == 201 {
				WriteResponse(*flagStatus, "", w)
			}
			WriteResponse(*flagStatus, *flagBody, w)
		},
	)
}

var logger = setupLog()

func main() {
	setFlags()
	flag.Parse()
	if !*flagQuiet {
		logger.Println(disclaimer)
	}
	hostPort := fmt.Sprintf(":%d", *flagPort)
	http.Handle("/", dynHandler())
	http.ListenAndServe(hostPort, nil)

}
