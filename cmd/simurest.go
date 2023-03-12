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
	"fmt"
	"net/http"

	flag "github.com/spf13/pflag"

	backend "network.golang/simurest/internal"
	util "network.golang/simurest/util"
)

var (
	flags util.Flags
)

func init() {
	flags = util.Flags{}
	flags.FlagQuiet = flag.BoolP("quiet", "q", false, "to supress the disclaimer header")
	flags.FlagMethod = flag.StringP("method", "m", "GET", "http method to accept")
	flags.FlagPort = flag.IntP("port", "p", 8080, "tcp port to listen")
	flags.FlagUri = flag.StringP("uri", "u", "/", "uri to handle")
	flags.FlagStatus = flag.IntP("status", "s", 200, "http status to return")
	flags.FlagBody = flag.StringP("body", "b", `{"status": "ok"}`, "response body to return")

}

func main() {
	flag.Parse()
	backend.Logger = backend.SetupLog()
	backend.FlagsList = []util.Flags{flags}
	backend.DisclaimerPrint()
	hostPort := fmt.Sprintf(":%d", *flags.FlagPort)
	http.Handle("/", backend.DynHandler())
	http.ListenAndServe(hostPort, nil)
}
