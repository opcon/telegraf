package fieldsystem

import "math"
import "strings"

// Return the new or updated values in "new"
func mapdiff(old map[string]interface{}, new map[string]interface{}) map[string]interface{} {
	diff := make(map[string]interface{})
	for k, nval := range new {
		if oval, ok := old[k]; !ok || oval != nval {
			diff[k] = nval
		}
	}
	return diff
}

func IsNan32(f float32) bool {
	return f != f
}

func IsInf32(f float32) bool {
	return f < -math.MaxFloat32 || f > math.MaxFloat32
}

// Convert a null-terminated byte array to a native Go string
func cstr(str []byte) string {
	n := 0
	for ; str[n] != 0; n++ {
	}
	if n == 0 {
		return ""
	}
	return string(str[:n-1])
}

// Convert a space padded byte array to a native Go string
func fsstr(s []byte) string {
	return strings.TrimSpace(string(s))
}
