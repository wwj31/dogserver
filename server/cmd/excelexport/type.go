package main

import (
	"fmt"
	"strconv"
	"strings"
)

type fnConvert func(val string) (interface{}, error)

type array_int []int64
type array_str []string
type array_float []float32
type int2int map[int64]int64
type int2str map[int64]string
type str2int map[string]int64
type str2str map[string]string

const (
	INT        = "int"
	STR        = "string"
	FLOAT      = "float"
	ARRAYINT   = "[int]"
	ARRAYSTR   = "[string]"
	ARRAYFLOAT = "[float]"
	INT2INT    = "map[int]int"
	INT2STR    = "map[int]string"
	STR2INT    = "map[string]int"
	STR2STR    = "map[string]string"

	ARRAY_SPLIT = "|"
	MAP_SPLIT   = ":"
)

var TypeIndex = map[string]string{
	INT:        "int64",
	STR:        "string",
	FLOAT:      "float32",
	ARRAYINT:   "array_int",
	ARRAYSTR:   "array_str",
	ARRAYFLOAT: "array_float",
	INT2INT:    "int2int",
	INT2STR:    "int2str",
	STR2INT:    "str2int",
	STR2STR:    "str2str",
}

var KeyIndex = map[string]string{
	INT:        "int64",
	STR:        "string",
	FLOAT:      "float32",
	ARRAYINT:   "int",
	ARRAYSTR:   "int",
	ARRAYFLOAT: "int",
	INT2INT:    "int64",
	INT2STR:    "int64",
	STR2INT:    "string",
	STR2STR:    "string",
}

var ValueIndex = map[string]string{
	INT:        "int64",
	STR:        "string",
	FLOAT:      "float32",
	ARRAYINT:   "int64",
	ARRAYSTR:   "string",
	ARRAYFLOAT: "float32",
	INT2INT:    "int64",
	INT2STR:    "string",
	STR2INT:    "int64",
	STR2STR:    "string",
}

var TypeConvert = map[string]fnConvert{
	INT:        convert_int,
	STR:        convert_string,
	FLOAT:      convert_float,
	ARRAYINT:   convert_intSlice,
	ARRAYSTR:   convert_stringSlice,
	ARRAYFLOAT: convert_floatSlice,
	INT2INT:    convert_int2int,
	INT2STR:    convert_int2string,
	STR2INT:    convert_string2int,
	STR2STR:    convert_string2string,
}

// return int64
func convert_int(val string) (interface{}, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return 0, nil
	}

	i, e := strconv.Atoi(val)
	if e != nil {
		return 0, fmt.Errorf("convert_int格式错误 value=%v err=%v", val, e)
	}
	return int64(i), nil
}

// return string
func convert_string(val string) (interface{}, error) {
	return val, nil
}

// return float32
func convert_float(val string) (interface{}, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return 0, nil
	}

	i, e := strconv.ParseFloat(val, 32)
	if e != nil {
		return 0, fmt.Errorf("convert_float格式错误 value=%v err=%v", val, e)
	}
	return float32(i), nil
}

// return []int64
func convert_intSlice(val string) (interface{}, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil, nil
	}

	strs := strings.Split(val, ARRAY_SPLIT)
	intslice := array_int{}
	for _, v := range strs {
		i, e := strconv.Atoi(v)
		if e != nil {
			return nil, fmt.Errorf("convert_intSlice格式错误 value=%v err=%v", val, e)
		}
		intslice = append(intslice, int64(i))
	}
	return intslice, nil
}

// return []string
func convert_stringSlice(val string) (interface{}, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil, nil
	}

	strslice := array_str{}
	strs := strings.Split(val, ARRAY_SPLIT)
	for _, str := range strs {
		strslice = append(strslice, str)
	}
	return strslice, nil
}

// return []float32
func convert_floatSlice(val string) (interface{}, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil, nil
	}

	floatSlice := array_float{}
	strs := strings.Split(val, ARRAY_SPLIT)
	for _, v := range strs {
		i, e := strconv.ParseFloat(v, 32)
		if e != nil {
			return nil, fmt.Errorf("convert_floatSlice格式错误 value=%v err=%v", val, e)
		}
		floatSlice = append(floatSlice, float32(i))
	}
	return floatSlice, nil
}

// return map[int64]int64
func convert_int2int(val string) (interface{}, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil, nil
	}

	int2int := int2int{}
	strs := strings.Split(val, ARRAY_SPLIT)
	for _, v := range strs {
		subStrs := strings.Split(v, MAP_SPLIT)
		if len(subStrs) != 2 {
			return nil, fmt.Errorf("convert_int2int格式错误1 value=%v err=length", val)
		}
		key, e := strconv.Atoi(subStrs[0])
		if e != nil {
			return nil, fmt.Errorf("convert_int2int格式错误2 value=%v err=%v", val, e)
		}
		v, e := strconv.Atoi(subStrs[1])
		if e != nil {
			return nil, fmt.Errorf("convert_int2int格式错误3 value=%v err=%v", val, e)
		}
		int2int[int64(key)] = int64(v)
	}

	return int2int, nil
}

// return map[int64]string
func convert_int2string(val string) (interface{}, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil, nil
	}

	int2string := int2str{}
	strs := strings.Split(val, ARRAY_SPLIT)
	for _, v := range strs {
		subStrs := strings.Split(v, MAP_SPLIT)
		if len(subStrs) != 2 {
			return nil, fmt.Errorf("convert_int2string格式错误1 value=%v err=length", val)
		}
		key, e := strconv.Atoi(subStrs[0])
		if e != nil {
			return nil, fmt.Errorf("convert_int2string格式错误2 value=%v err=%v", val, e)
		}
		int2string[int64(key)] = subStrs[1]
	}
	return int2string, nil
}

// return map[string]int64
func convert_string2int(val string) (interface{}, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil, nil
	}

	string2int := str2int{}
	strs := strings.Split(val, ARRAY_SPLIT)
	for _, v := range strs {
		subStrs := strings.Split(v, MAP_SPLIT)
		if len(subStrs) != 2 {
			return nil, fmt.Errorf("convert_string2int格式错误1 value=%v err=length", val)
		}
		v, e := strconv.Atoi(subStrs[1])
		if e != nil {
			return nil, fmt.Errorf("convert_string2int格式错误 value=%v err=%v", val, e)
		}
		string2int[subStrs[0]] = int64(v)
	}
	return string2int, nil
}

// return map[string]string
func convert_string2string(val string) (interface{}, error) {
	val = strings.TrimSpace(val)
	if val == "" {
		return nil, nil
	}

	string2string := str2str{}
	strs := strings.Split(val, ARRAY_SPLIT)
	for _, v := range strs {
		subStrs := strings.Split(v, MAP_SPLIT)
		if len(subStrs) != 2 {
			return nil, fmt.Errorf("convert_string2string格式错误1 value=%v err=length", val)
		}
		string2string[subStrs[0]] = subStrs[1]
	}
	return string2string, nil
}
