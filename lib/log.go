/*
 * Copyright 2018 InfAI (CC SES)
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package lib

import (
	"github.com/SmartEnergyPlatform/jwt-http-router"
)

func (this *Lib) GetDeviceLogStates(jwt jwt_http_router.Jwt, deviceIds []string) (result map[string]bool, err error) {
	result = map[string]bool{}
	err = jwt.Impersonate.PostJSON(this.config.ConnectionLogUrl+"/intern/state/device/check", deviceIds, &result)
	return
}

func (this *Lib) GetGatewayLogStates(jwt jwt_http_router.Jwt, deviceIds []string) (result map[string]bool, err error) {
	result = map[string]bool{}
	err = jwt.Impersonate.PostJSON(this.config.ConnectionLogUrl+"/intern/state/gateway/check", deviceIds, &result)
	return
}

func (this *Lib) GetDeviceLogHistory(jwt jwt_http_router.Jwt, deviceIds []string, duration string) (result map[string]HistorySeries, err error) {
	return this.GetLogHistory(jwt, "device", deviceIds, duration)
}

func (this *Lib) GetGatewayLogHistory(jwt jwt_http_router.Jwt, ids []string, duration string) (result map[string]HistorySeries, err error) {
	return this.GetLogHistory(jwt, "gateway", ids, duration)
}

type HistoryResult struct {
	Series []HistorySeries `json:"Series"`
}

type HistorySeries struct {
	Name    string            `json:"name"`
	Tags    map[string]string `json:"tags"`
	Columns []string          `json:"columns"`
	Values  [][]interface{}   `json:"values"`
}

func (this *Lib) GetLogHistory(jwt jwt_http_router.Jwt, kind string, ids []string, duration string) (result map[string]HistorySeries, err error) {
	result = map[string]HistorySeries{}
	temp := []HistoryResult{}
	err = jwt.Impersonate.PostJSON(this.config.ConnectionLogUrl+"/intern/history/"+kind+"/"+duration, ids, &temp)
	if err != nil {
		return result, err
	}
	for _, series := range temp[0].Series {
		result[series.Tags[kind]] = series
	}
	return
}

func (this *Lib) GetLogstarts(jwt jwt_http_router.Jwt, kind string, ids []string) (result map[string]interface{}, err error) {
	result = map[string]interface{}{}
	err = jwt.Impersonate.PostJSON(this.config.ConnectionLogUrl+"/intern/logstarts/"+kind, ids, &result)
	return
}

func (this *Lib) GetLogedges(jwt jwt_http_router.Jwt, kind string, ids []string, duration string) (result map[string]interface{}, err error) {
	result = map[string]interface{}{}
	err = jwt.Impersonate.PostJSON(this.config.ConnectionLogUrl+"/intern/logedge/"+kind+"/"+duration, ids, &result)
	return
}
