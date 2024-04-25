// SPDX-License-Identifier: Apache-2.0
// Copyright 2023 Cloudbase Solutions SRL
//
//    Licensed under the Apache License, Version 2.0 (the "License"); you may
//    not use this file except in compliance with the License. You may obtain
//    a copy of the License at
//
//         http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
//    WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
//    License for the specific language governing permissions and limitations
//    under the License.

package provider

import (
	"testing"

	commonParams "github.com/cloudbase/garm-provider-common/params"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtraSpecs(t *testing.T) {
	tests := []struct {
		name           string
		specs          commonParams.BootstrapInstance
		expectedOutput extraSpecs
		errString      string
	}{
		{
			name: "Empty specs",
			specs: commonParams.BootstrapInstance{
				ExtraSpecs: nil,
			},
			expectedOutput: extraSpecs{},
			errString:      "",
		},
		{
			name: "Valid specs",
			specs: commonParams.BootstrapInstance{
				ExtraSpecs: []byte(`{"disable_updates": true, "extra_packages": ["package1", "package2"], "enable_boot_debug": false}`),
			},
			expectedOutput: extraSpecs{
				DisableUpdates:  true,
				ExtraPackages:   []string{"package1", "package2"},
				EnableBootDebug: false,
			},
			errString: "",
		},
		{
			name: "Invalid specs",
			specs: commonParams.BootstrapInstance{
				ExtraSpecs: []byte(`{"disable_updates": true, "extra_packages": ["package1", "package2", "package3], "enable_boot_debug": false}`),
			},
			expectedOutput: extraSpecs{},
			errString:      "unmarshaling extra specs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := parseExtraSpecsFromBootstrapParams(tt.specs)
			if tt.errString != "" {
				assert.ErrorContains(t, err, tt.errString)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.expectedOutput, output)
		})
	}
}
