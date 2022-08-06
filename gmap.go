package gmap

import (
	"strconv"
	"strings"
)

type Map map[string]interface{}

// Get get value for the specified path
func (m Map) Get(path string) Value {
	return get(m, path, false)
}

// Get get value for the specified path, but case-insensitive
func (m Map) GetIgnoreCase(path string) Value {
	return get(m, path, true)
}

type Slice []interface{}

// Get get value for the specified path
func (s Slice) Get(path string) Value {
	return get(s, path, false)
}

// Get get value for the specified path, but case-insensitive
func (s Slice) GetIgnoreCase(path string) Value {
	return get(s, path, true)
}

type nextResult struct {
	Index string
	Path  string
	More  bool
}

func nextIndex(path string) (r nextResult) {
	for i := 0; i < len(path); i++ {
		if path[i] == '\\' {
			epart := []byte(path[:i])
			i++
			if i < len(path) {
				epart = append(epart, path[i])
				i++
				for ; i < len(path); i++ {
					if path[i] == '\\' {
						i++
						if i < len(path) {
							epart = append(epart, path[i])
						}
						continue
					} else if path[i] == '.' {
						r.Index = string(epart)
						r.Path = path[i+1:]
						r.More = true
						return
					}
					epart = append(epart, path[i])
				}
			}
			r.Index = string(epart)
			return
		}
		if path[i] == '.' {
			r.Index = path[:i]
			r.Path = path[i+1:]
			r.More = true
			return
		}
	}
	r.Index = path
	return
}

func get(v interface{}, path string, ignoreCase bool) Value {
	next := nextResult{Path: path}

	for {
		next = nextIndex(next.Path)

		var m map[string]interface{}
		var s []interface{}
		var smode bool

		switch t := v.(type) {
		case Map:
			m = t
		case map[string]interface{}:
			m = t
		case Slice:
			s = t
			smode = true
		case []interface{}:
			s = t
			smode = true
		default:
			return Value{}
		}

		if smode {
			if idx, err := strconv.Atoi(next.Index); err != nil || idx < 0 || idx >= len(s) {
				return Value{}
			} else {
				v = s[idx]
			}
		} else {
			var fnd bool
			v, fnd = m[next.Index]
			if !fnd && ignoreCase {
				for k, vv := range m {
					if strings.EqualFold(k, next.Index) {
						v = vv
						break
					}
				}
			}
		}
		if !next.More {
			return Value{Has: true, Raw: v}
		}
	}
}
