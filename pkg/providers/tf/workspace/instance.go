// Copyright 2018 the Service Broker Project Authors.
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

package workspace

import (
	"encoding/json"
	"fmt"
)

// ModuleInstance represents the configuration of a single instance of a module.
type ModuleInstance struct {
	ModuleName    string                 `json:"module_name"`
	InstanceName  string                 `json:"instance_name"`
	Configuration map[string]interface{} `json:"configuration"`
}

// MarshalDefinition converts the module instance definition into a JSON
// definition that can be fed to Terraform to be created/destroyed.
func (instance *ModuleInstance) MarshalDefinition(outputs []string) (json.RawMessage, error) {
	instanceConfig := make(map[string]interface{})
	for k, v := range instance.Configuration {
		instanceConfig[k] = v
	}

	instanceConfig["source"] = fmt.Sprintf("./%s", instance.ModuleName)

	outputMap := make(map[string]interface{})
	for _, variable := range outputs {
		outputMap[variable] = map[string]string{"value": fmt.Sprintf("${module.%s.%s}", instance.InstanceName, variable)}
	}

	defn := map[string]interface{}{
		"module": map[string]interface{}{
			instance.InstanceName: instanceConfig,
		},
	}

	if len(outputMap) > 0 {
		defn["output"] = outputMap
	}

	return json.Marshal(defn)
}
