/*
Licensed to the Apache Software Foundation (ASF) under one or more
contributor license agreements.  See the NOTICE file distributed with
this work for additional information regarding copyright ownership.
The ASF licenses this file to You under the Apache License, Version 2.0
(the "License"); you may not use this file except in compliance with
the License.  You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package api

import (
	"testing"

	"github.com/apache/incubator-devlake/core/errors"
	"github.com/apache/incubator-devlake/impls/logruslog"
	"github.com/stretchr/testify/assert"
)

func TestApiClientBlackList(t *testing.T) {
	for _, tc := range []struct {
		Name      string
		Pattern   string
		Endpoints []string
		Err       errors.Error
	}{
		{
			Name:    "Internal IP Addresses",
			Pattern: "10.0.0.1/16",
			Endpoints: []string{
				"https://10.0.0.1",
				"http://10.0.0.254",
				"http://10.0.254.1",
				"https://10.0.254.254",
			},
			Err: ErrHostNotAllowed,
		},
		{
			Name:    "Internal IP Addresses",
			Pattern: "10.0.0.1/16",
			Endpoints: []string{
				"http://10.1.0.1",
				"http://10.1.0.254",
				"http://10.1.254.1",
				"http://10.1.254.254",
			},
			Err: nil,
		},
	} {
		t.Run(tc.Name, func(t *testing.T) {
			for _, endpoint := range tc.Endpoints {
				err := checkCidrBlacklist(tc.Pattern, endpoint, logruslog.Global)
				assert.Equal(t, tc.Err, err, "pattern %s and endpoint %s should return %v, but got %v", tc.Pattern, endpoint, tc.Err, err)
			}
		})
	}
}
