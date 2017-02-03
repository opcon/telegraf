package fieldsystem

import "math"
import "strings"

func IsNan32(f float32) bool {
	return f != f
}

func IsInf32(f float32) bool {
	return f < -math.MaxFloat32 || f > math.MaxFloat32
}

// Convert a space padded byte array to a native Go string
func fsstr(s []byte) string {
	return strings.TrimSpace(string(s))
}
