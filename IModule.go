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

import "net/http"

// interface encapsulating the method(s) for a module running on echogogo server
type IModule interface {
	// method to return the config model for a REST api path / endPoint(s)
	GetRestConfig() RestConfigModel

	// method to perform the actual action for the module and return a model (interface{})
	// for the final response rendering
	DoAction(request http.Request, response http.Response, endPoint string, optionalMap ...map[string]interface{}) interface{}
}



