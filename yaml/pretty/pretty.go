// Copyright 2019 Drone IO, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pretty

import (
	"io"

	"github.com/oliashish/drone-yaml/yaml"
)

// Print pretty prints the manifest.
func Print(w io.Writer, v *yaml.Manifest) {
	state := new(baseWriter)
	for _, r := range v.Resources {
		switch t := r.(type) {
		case *yaml.Cron:
			printCron(state, t)
		case *yaml.Secret:
			printSecret(state, t)
		case *yaml.Signature:
			printSignature(state, t)
		case *yaml.Pipeline:
			printPipeline(state, t)
		}
	}
	state.WriteString("...")
	state.WriteByte('\n')
	w.Write(state.Bytes())
}
