package main

import (
	"errors"
)

func doMerge(data map[string]any, items []*ConfigItem) (map[string]any, error) {
	for _, item := range items {
		var err error
		data, err = merge(data, item.Data, item.Operation)
		if err != nil {
			return nil, err
		}
	}
	return data, nil
}

func merge(conf1, conf2 map[string]any, operation string) (map[string]any, error) {
	for k,v := range conf2 {
		_,ok := conf1[k]
		if !ok || operation == "replace"{
			conf1[k] = v
			continue
		}

		if operation == "merge"{
			_val, err := mergeVal(conf1[k], v)
			if err != nil {
				return nil, err
		}
			conf1[k] = _val
		}
	}
	return conf1, nil
}


func mergeVal(val1, val2 any) (any, error) {
	switch v1 := val1.(type) {
	case []any:
		switch v2 := val2.(type) {
		case []any:
			return append(v1, v2...), nil
		default:
			return nil, errors.New("append operation not supported for non-list types")
		}
	case map[string]any:
		switch v2 := val2.(type) {
		case map[string]any:
			for k, v := range v2 {
				v1[k] = v
			}
			return v1, nil
		default:
			return nil, errors.New("append operation not supported for non-map types")
		}
	default:
		return val2,nil
	}
}
