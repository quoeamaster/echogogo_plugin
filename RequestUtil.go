/*
 * Licensed to Echogogo under one or more contributor
 * license agreements. See the NOTICE file distributed with
 * this work for additional information regarding copyright
 * ownership. Echogogo licenses this file to you under
 * the Apache License, Version 2.0 (the "License"); you may
 * not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package echogogo

import "strings"

// method to extract the "path-parameter" of the given url according to the given endpoint
// e.g. url => /mock/getCityNameByZip AND endpoint => /mock -- path-parameter => /getCityNameByZip
func extractPathParameterFromUrl(url string, endpoint string) string {
	methodName := ""

	if idx := strings.Index(url, endpoint); idx != -1 {
		endpointLen := len(endpoint)
		methodName = url[idx + endpointLen:]
	}
	return methodName
}
