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

// format in json
const FORMAT_JSON 		= "json"
// format in xml
const FORMAT_XML 		= "xml"
// format in xml OR json
const FORMAT_XML_JSON 	= "xml,json"

// Model encapsulating REST api configuration(s)
type RestConfigModel struct {
	// main path for the REST api service e.g. "/echo"
	Path string					`json:"path" description:"main path for the REST api"`

	// api endpoints to be mapped against the main path e.g. [ "/get", "/get/{user_id}" ]; then the actual path
	// would be "/echo/get" and "/echo/get/{user-id}"
	EndPoints []string			`json:"endPoints" description:"api endPoints under this main path"`

	// default request in what format?
	ConsumeFormat string		`json:"consumeFormat" description:"default format for request"`

	// default response in what format?
	ProduceFormat string		`json:"produceFormat" description:"default format for response"`

	optionalConfigs map[string]interface{}
}

// method to init the optinalConfigs map
func (m *RestConfigModel) initOptionalConfigs() {
	if m.optionalConfigs == nil {
		m.optionalConfigs = make(map[string]interface{})
	}
}

// method to add/update the value under the key
func (m *RestConfigModel) PutConfigByKey(key string, value interface{}) {
	m.initOptionalConfigs()
	m.optionalConfigs[key] = value
}

// method to get the value under the key
func (m *RestConfigModel) GetConfigByKey(key string, defaultValue interface{}) interface{} {
	m.initOptionalConfigs()
	value, found := m.optionalConfigs[key]
	if !found {
		value = defaultValue
	}
	return value
}

// method to remove the value under the key
func (m *RestConfigModel) RemoveConfigByKey(key string) {
	m.initOptionalConfigs()
	delete(m.optionalConfigs, key)
}

