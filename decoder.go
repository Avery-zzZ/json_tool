package json_tool

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
)

type JsonStringMap map[string]interface{}

func Handle_json_byte(body io.ReadCloser) (*JsonStringMap, error) {
	json_str, err := ioutil.ReadAll(body) //->[]byte
	if err != nil {
		return nil, errors.New("json_handle/json_handle.go - Handle_json_byte:\n"+err.Error())
	}

	var json_itf interface{} //->interface{}
	err = json.Unmarshal(json_str, &json_itf)
	if err != nil {
		return nil, err
	}

	json_map, ok := json_itf.(map[string]interface{}) //->map[string]interface{}
	if !ok {
		return nil, errors.New("json_handle/json_handle.go - Handle_json_byte:\ninterface{} -> map[string]interface{} fail")
	}

	return (*JsonStringMap)(&json_map), nil
}

func Handle_json(json_itf interface{}) (*JsonStringMap, error) {

	json_map, ok := json_itf.(map[string]interface{}) //->map[string]interface{}
	if !ok {
		return nil, errors.New("json_handle/json_handle.go - Handle_json_byte:\ninterface{} -> map[string]interface{} fail")
	}

	return (*JsonStringMap)(&json_map), nil
}

func (jsonstringmap *JsonStringMap) Find(route []string) (interface{}, error) {
	if route == nil {
		return nil, errors.New("json_handle/json_handle.go - Find:\nempty route")
	}

	temp_map := *jsonstringmap
	var ind int

	for ind = 0; ind < len(route)-1; ind++ {
		if temp_itf, ok := temp_map[route[ind]]; !ok {
			return nil, errors.New("json_handle/json_handle.go - Find:\nno matching element \""+route[ind]+"\"")
		} else {
			// fmt.Println(temp_itf)
			temp_map = temp_itf.(map[string]interface{})
		}
	}

	if temp_itf, ok := temp_map[route[ind]]; !ok {
		return nil, errors.New("json_handle/json_handle.go - Find:\nno matching element \""+route[ind]+"\"")
	} else {
		return temp_itf, nil
	}
}


