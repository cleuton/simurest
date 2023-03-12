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

package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	util "network.golang/simurest/util"
)

const Disclaimer = `
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
	Logger    log.Logger
	FlagsList []util.Flags
)

func SetupLog() log.Logger {
	Logger := log.Default()
	return *Logger
}

func DisclaimerPrint() {
	if *FlagsList[0].FlagQuiet == false {
		fmt.Println(Disclaimer)
	}
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

func CheckUrl(uri string, currentFlag util.Flags) bool {
	if (*currentFlag.FlagUri)[len((*currentFlag.FlagUri))-1] == '*' {
		if strings.Contains(uri, (*currentFlag.FlagUri)[:len((*currentFlag.FlagUri))-1]) {
			return true
		}
	} else if uri == *currentFlag.FlagUri {
		return true
	}
	return false
}

func CheckMethod(method string, currentFlag util.Flags) bool {
	if method == *currentFlag.FlagMethod {
		return true
	}
	return false
}

func DynHandler() http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			currentFlag := FlagsList[0]
			Logger.Println(
				fmt.Sprintf("%v -> %v %v%v", strings.Split(req.Proto, "/")[0], req.Method, req.Host, req.URL.Path))
			if !CheckUrl(req.URL.Path, currentFlag) {
				WriteResponse(404, "Not found", w)
				return
			}
			if !CheckMethod(req.Method, currentFlag) {
				WriteResponse(405, "Method not allowed", w)
				return
			}
			if len(*currentFlag.FlagBody) == 0 || *currentFlag.FlagStatus == 201 {
				WriteResponse(*currentFlag.FlagStatus, "", w)
			}
			WriteResponse(*currentFlag.FlagStatus, *currentFlag.FlagBody, w)
		},
	)
}
