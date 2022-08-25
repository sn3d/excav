package cast

import "strconv"

func ToStr(in interface{}) string {
	if in == nil {
		return ""
	}
	return in.(string)
}

func ToData(in interface{}) map[interface{}]interface{} {
	return in.(map[interface{}]interface{})
}

func ToStrData(in interface{}) map[string]interface{} {
	out := make(map[string]interface{})
	inData := ToData(in)
	for k, v := range inData {
		out[k.(string)] = v
	}
	return out
}

func ToStrArr(in interface{}) []string {
	inArr := in.([]interface{})
	out := make([]string, len(inArr))
	for i, val := range inArr {
		out[i] = val.(string)
	}
	return out
}

func ToUint(in interface{}) uint64 {
	val := ToStr(in)
	num, _ := strconv.ParseUint(val, 10, 0)
	return num
}
