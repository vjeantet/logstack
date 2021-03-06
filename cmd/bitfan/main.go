//go:generate go generate github.com/vjeantet/bitfan/processors/...
//go:generate go generate github.com/vjeantet/bitfan/codecs/...
// broken - go:generate swagger generate spec -m -b github.com/vjeantet/bitfan/api -o ../../api/swagger.json
// Copyright © 2016 Valere JEANTET <valere.jeantet@gmail.com>
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

package main

import (
	"log"
	"os"
	"runtime"

	"github.com/kardianos/service"
	"github.com/vjeantet/bitfan/cmd/bitfan/commands"
)

var version = "master"
var buildstamp = ""

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	// Service
	if !service.Interactive() {

		// PASS Service
		s := commands.GetService()

		slogger, err := s.Logger(nil)
		if err != nil {
			log.Fatal("EOOR", err)
		}
		err = s.Run()
		if err != nil {
			slogger.Error("service startup error : ", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	//interactive
	commands.Version = version
	commands.Buildstamp = buildstamp
	commands.Execute()

}
