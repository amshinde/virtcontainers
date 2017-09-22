//
// Copyright (c) 2017 Intel Corporation
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
//

package virtcontainers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBDF(t *testing.T) {
	type testData struct {
		deviceStr   string
		expectedBDF string
	}

	data := []testData{
		{"0000:02:10.0", "02:10.0"},
		{"0000:0210.0", ""},
		{"test", ""},
		{"", ""},
	}

	for _, d := range data {
		deviceBDF, err := bdf(d.deviceStr)
		assert.Equal(t, d.expectedBDF, deviceBDF)
		if d.expectedBDF == "" {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}

func TestIsVFIODeviceGroup(t *testing.T) {
	type testData struct {
		path     string
		expected bool
	}

	data := []testData{
		{"/dev/vfio/16", true},
		{"/dev/vfio/1", true},
		{"/dev/vfio/", false},
		{"/dev/vfio", false},
		{"/dev/vf", false},
		{"/dev", false},
	}

	for _, d := range data {
		device := Device{
			Path: d.path,
		}
		isVFIO := device.isVFIODeviceGroup()
		assert.Equal(t, d.expected, isVFIO)
	}
}
