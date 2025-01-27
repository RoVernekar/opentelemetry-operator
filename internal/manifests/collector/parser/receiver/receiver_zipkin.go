// Copyright The OpenTelemetry Authors
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

package receiver

import (
	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
)

const parserNameZipkin = "__zipkin"

// NewZipkinReceiverParser builds a new parser for Zipkin receivers.
func NewZipkinReceiverParser(logger logr.Logger, name string, config map[interface{}]interface{}) ReceiverParser {
	http := "http"
	return &GenericReceiver{
		logger:             logger,
		name:               name,
		config:             config,
		defaultPort:        9411,
		defaultProtocol:    corev1.ProtocolTCP,
		defaultAppProtocol: &http,
		parserName:         parserNameZipkin,
	}
}

func init() {
	Register("zipkin", NewZipkinReceiverParser)
}
