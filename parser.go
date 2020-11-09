package main

import (
	"errors"
)

func fromJSONToMap(data interface{}) (map[string]interface{}, error) {
	if data == nil {
		return nil, errors.New("Cannot parse an empty interface{}")
	}

	if content, ok := data.(map[interface{}]interface{}); ok {
		normalizedmap := make(map[string]interface{})

		for k, v := range content {
			switch k.(type) {
			case string:
				normalizedmap[k.(string)] = v
			default:
				continue
			}
		}

		return normalizedmap, nil
	}

	return nil, nil
}

func fromJSONToArrayMap(data interface{}) ([]map[string]interface{}, error) {
	if data == nil {
		return nil, errors.New("Cannot parse an empty interface{}")
	}

	if contents, ok := data.([]interface{}); ok {
		normalizedmap := make([]map[string]interface{}, 0)

		for _, content := range contents {
			result, err := fromJSONToMap(content)

			if err != nil || result == nil {
				continue
			}

			normalizedmap = append(normalizedmap, result)
		}

		return normalizedmap, nil
	}

	return []map[string]interface{}{}, nil
}
