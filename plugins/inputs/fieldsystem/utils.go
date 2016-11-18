package fieldsystem

import "math"

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
