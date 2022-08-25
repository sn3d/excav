package excav

import "strings"

// parameters they're passed into every task. They're used
// for parametrized tasks like 'append-template' etc.
// The parameters can be defined as arguments during apply
// (see excav apply --h)
type Params map[string]string

// create new empty collection of parameters
func NewParams() Params {
	return make(map[string]string)
}

// ParamsAsStringSlice consume parameters as array of  'KEY=VALUE' strings
// and parse it to Params structure - which is map btw. The 'KEY=VALUE' is used
// mainly for parameters given as cli argument.
func StringSliceToParams(params []string) Params {
	res := make(Params)
	for _, p := range params {
		idx := strings.Index(p, "=")
		if idx > 0 {
			res[p[0:idx]] = p[idx+1:]
		}
	}
	return res
}

// ToMap convert parameters into map of interfaces.
// This is needed mainly for patch, because patch is expecting
// parameters as map of interfaces.
func (p Params) ToMap() map[string]interface{} {
	res := make(map[string]interface{}, len(p))
	for k, v := range p {
		res[k] = v
	}
	return res
}

// merges multiple parameter collections into
// one. If the parameter is obtained in multiple collections,
// the value of last collection in function is used. Collections
// have priorrity from lowest to high.
func MergeParams(params ...Params) Params {
	res := NewParams()
	for _, p := range params {
		for key, val := range p {
			res[key] = val
		}
	}
	return res
}
