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
	"errors"
	"github.com/SmartEnergyPlatform/jwt-http-router"
	"sort"
	"strings"

	"log"
)

func GetConnectionFilteredDevicesOrder(jwt jwt_http_router.Jwt, value string, sortAsc bool) (result []map[string]interface{}, err error) {
	result, err = GetConnectionFilteredDevices(jwt,value)
	if err != nil {
		log.Println("ERROR GetConnectionFilteredDevices", err)
		return result, err
	}

	result = sortByName(result, sortAsc)

	return
}

func GetConnectionFilteredDevicesSearchOrder(jwt jwt_http_router.Jwt, value string, searchText string, sortAsc bool) (result []map[string]interface{}, err error) {
	result, err = GetConnectionFilteredDevices(jwt,value)

	if err != nil {
		log.Println("ERROR GetConnectionFilteredDevices", err)
		return result, err

	}

	result = filter(result, "name", searchText)
	result = sortByName(result, sortAsc)

	return
}

func sortByName(input []map[string]interface{}, sortAsc bool) (output []map[string]interface{})  {
	output = input
	if sortAsc == true {
		sort.Slice(output, func(i, j int) bool {
			return output[i]["name"].(string) < output[j]["name"].(string)
		})
	} else {
		sort.Slice(output, func(i, j int) bool {
			return output[i]["name"].(string) > output[j]["name"].(string)

		})
	}
	return
}

func filter(list []map[string]interface{}, key string, value string)(result []map[string]interface{}) {
	for _, element := range list{
		str, ok := element[key].(string)
		if ok && strings.Contains(str, value) {
			result = append(result, element)
		}
	}
	return
}

func GetConnectionFilteredDevices(jwt jwt_http_router.Jwt, value string) (result []map[string]interface{}, err error) {
	devices, err := PermListAllDevices(jwt, "r")
	if err != nil {
		log.Println("ERROR GetConnectionFilteredDevices.PermListAllDevices()", err)
		return result, err
	}
	return FilterDevicesByState(jwt, devices, value)
}

func FilterDevicesByState(jwt jwt_http_router.Jwt, devices []map[string]interface{}, state string)(result []map[string]interface{}, err error){
	devicesWithOnlineState, err := completeDeviceList(jwt, devices)
	if err != nil {
		log.Println("ERROR GetConnectionFilteredDevices.completeDeviceList()", err)
		return devices, err
	}

	for _, device := range devicesWithOnlineState {
		devicestate, ok := device["log_state"]
		if state == "connected" && ok && devicestate.(bool) {
			result = append(result, device)
		}
		if state == "disconnected" && ok && !devicestate.(bool) {
			result = append(result, device)
		}
		if state == "unknown" && !ok {
			result = append(result, device)
		}
	}
	return
}

func ListDevices(jwt jwt_http_router.Jwt, limit string, offset string) (result []map[string]interface{}, err error) {
	devices, err := PermListDevices(jwt, "r", limit, offset)
	if err != nil {
		log.Println("ERROR ListDevices.PermListDevices()", err)
		return result, err
	}
	return completeDeviceList(jwt, devices)
}

func ListDevicesOrdered(jwt jwt_http_router.Jwt, limit string, offset string, orderfeature string, direction string) (result []map[string]interface{}, err error) {
	devices, err := PermListDevicesOrdered(jwt, "r", limit, offset, orderfeature, direction)
	if err != nil {
		log.Println("ERROR ListDevices.PermListDevices()", err)
		return result, err
	}
	return completeDeviceList(jwt, devices)
}

func SearchDevices(jwt jwt_http_router.Jwt, query string, limit string, offset string) (result []map[string]interface{}, err error) {
	devices, err := PermSearchDevices(jwt, query, "r", limit, offset)
	if err != nil {
		return result, err
	}
	return completeDeviceList(jwt, devices)
}

func SearchDevicesOrdered(jwt jwt_http_router.Jwt, query string, limit string, offset string, orderfeature string, direction string) (result []map[string]interface{}, err error) {
	devices, err := PermSearchDevicesOrdered(jwt, query, "r", limit, offset, orderfeature, direction)
	if err != nil {
		return result, err
	}
	return completeDeviceList(jwt, devices)
}

func ListDevicesByTag(jwt jwt_http_router.Jwt, value string) (result []map[string]interface{}, err error) {
	devices, err := PermSelectTagDevices(jwt, value, "r")
	if err != nil {
		return result, err
	}
	return completeDeviceList(jwt, devices)
}

func ListOrderdDevicesByTag(jwt jwt_http_router.Jwt, value string, limit string, offset string, orderfeature string, direction string) (result []map[string]interface{}, err error) {
	devices, err := PermSelectTagDevicesOrdered(jwt, value, "r", limit, offset, orderfeature, direction)
	if err != nil {
		return result, err
	}
	return completeDeviceList(jwt, devices)
}

func ListDevicesByUserTag(jwt jwt_http_router.Jwt, value string) (result []map[string]interface{}, err error) {
	devices, err := PermSelectUserTagDevices(jwt, value, "r")
	if err != nil {
		return result, err
	}
	return completeDeviceList(jwt, devices)
}

func ListOrderedDevicesByUserTag(jwt jwt_http_router.Jwt, value string, limit string, offset string, orderfeature string, direction string) (result []map[string]interface{}, err error) {
	devices, err := PermSelectUserTagDevicesOrdered(jwt, value, "r", limit, offset, orderfeature, direction)
	if err != nil {
		return result, err
	}
	return completeDeviceList(jwt, devices)
}

func CompleteDevices(jwt jwt_http_router.Jwt, ids []string) (result []map[string]interface{}, err error) {
	devices, err := PermDeviceIdList(jwt, ids, "r")
	if err != nil {
		return result, err
	}
	return completeDeviceList(jwt, devices)
}

func CompleteDevicesOrdered(jwt jwt_http_router.Jwt, ids []string, limit string, offset string, orderfeature string, direction string) (result []map[string]interface{}, err error) {
	devices, err := PermDeviceIdListOrdered(jwt, ids, "r", limit, offset, orderfeature, direction)
	if err != nil {
		return result, err
	}
	return completeDeviceList(jwt, devices)
}

func completeDeviceList(jwt jwt_http_router.Jwt, devices []map[string]interface{}) (result []map[string]interface{}, err error) {
	ids := []string{}
	deviceMap := map[string]map[string]interface{}{}
	for _, device := range devices {
		id, ok := device["id"]
		if !ok {
			err = errors.New("unable to get device id")
			return
		}
		idStr, ok := id.(string)
		if !ok {
			err = errors.New("unable to cast device id to string")
			return
		}
		ids = append(ids, idStr)
		deviceMap[idStr] = device
	}
	logStates, err := GetDeviceLogStates(jwt, ids)
	if err != nil {
		log.Println("ERROR completeDeviceList.GetDeviceLogStates()", err)
		return result, err
	}
	/*
		gateways, err := GatewayNames(jwt, ids)
		if err != nil {
			log.Println("ERROR completeDeviceList.GatewayNames()", err)
			return result, err
		}
	*/
	for _, id := range ids {
		device := deviceMap[id]
		logState, logExists := logStates[id]
		if logExists {
			device["log_state"] = logState
		}
		//device["gateway_name"] = gateways[id]
		result = append(result, device)
	}
	return
}

func GetDevicesHistory(jwt jwt_http_router.Jwt, duration string) (result []map[string]interface{}, err error) {
	result, err = PermListAllDevices(jwt, "r")
	if err != nil {
		log.Println("ERROR PermListAllDevices()", err)
		return result, err
	}
	result, err = completeDeviceHistory(jwt, duration, result)
	return
}

func completeDeviceHistory(jwt jwt_http_router.Jwt, duration string, devices []map[string]interface{}) (result []map[string]interface{}, err error) {
	ids := []string{}
	deviceMap := map[string]map[string]interface{}{}
	for _, device := range devices {
		id, ok := device["id"]
		if !ok {
			err = errors.New("unable to get device id")
			return
		}
		idStr, ok := id.(string)
		if !ok {
			err = errors.New("unable to cast device id to string")
			return
		}
		ids = append(ids, idStr)
		deviceMap[idStr] = device
	}
	logStates, err := GetDeviceLogStates(jwt, ids)
	if err != nil {
		log.Println("ERROR completeDeviceList.GetDeviceLogStates()", err)
		return result, err
	}
	logHistory, err := GetDeviceLogHistory(jwt, ids, duration)
	if err != nil {
		log.Println("ERROR completeDeviceList.GetDeviceLogHistory()", err)
		return result, err
	}
	logEdges, err := GetLogedges(jwt, "device", ids, duration)
	if err != nil {
		log.Println("ERROR completeDeviceList.GetLogedges()", err)
		return result, err
	}
	for _, id := range ids {
		device := deviceMap[id]
		logState, logExists := logStates[id]
		if !logExists {
			device["log_state"] = "unknown"
		} else {
			if logState {
				device["log_state"] = "connected"
			} else {
				device["log_state"] = "disconnected"
			}
		}
		device["log_history"] = logHistory[id]
		device["log_edge"] = logEdges[id]
		result = append(result, device)
	}
	return
}
