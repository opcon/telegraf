//go:generate bash ./types.sh
//go:generate bash -c "cgo -godefs types.go | sed -e 's/Pad_cgo/pad_cgo/' -e 's/]int8/]byte/g' | gofmt > types_$GOARCH.go"
// Note: cgo must be patched to fix Issue #5253 (see patch at https://codereview.appspot.com/122900043)

package fieldsystem
