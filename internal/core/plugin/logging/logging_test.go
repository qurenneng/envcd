/*
 * Licensed to the AcmeStack under one or more contributor license
 * agreements. See the NOTICE file distributed with this work for
 * additional information regarding copyright ownership.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package logging

import (
	"net/http"
	"testing"

	"github.com/acmestack/envcd/internal/pkg/context"
)

func TestPrintLog(t *testing.T) {
	header := http.Header{}
	header["name"] = []string{"envcd"}
	body := make(map[string]interface{})
	body["id"] = "1"
	body["name"] = "envcd"
	c := context.Context{
		Uri:         "/test/uri",
		Method:      "POST",
		Headers:     header,
		ContentType: "application/json",
		Body:        body,
	}
	printLog(&c)
}
