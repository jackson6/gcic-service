/*
 *
 * Copyright 2017 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// This file contains wrappers for logger functions.
// The transport package only logs to verbose level 2 by default.

package transport

import "github.com/micro/grpc-go/logger"

const logLevel = 2

func infof(format string, args ...interface{}) {
	if logger.V(logLevel) {
		logger.Infof(format, args...)
	}
}

func warningf(format string, args ...interface{}) {
	if logger.V(logLevel) {
		logger.Warningf(format, args...)
	}
}

func errorf(format string, args ...interface{}) {
	if logger.V(logLevel) {
		logger.Errorf(format, args...)
	}
}
