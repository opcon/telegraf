#!/bin/bash

#Generate a Go file with all fieldsystem types defined in terms of C types

# Fragile
eval $(head -4 /usr2/fs/Makefile | sed 's/\s*=\s*/=/g' | tail -n+2 -)

read -r -d '' header <<EOF
// +build ignore
// generated with types.sh --- do not edit

package fieldsystem

/*
#include "/usr2/fs/include/params.h"
#include "/usr2/fs/include/fs_types.h"
#include "/usr2/fs/include/fscom.h"
*/
import "C"
EOF

f() {
cat <<EOF
$header
const (
    FieldSystemVersion = "$VERSION.$SUBLEVEL.$PATCHLEVEL"
)
type (
    $(echo "$header"\
        | grep include \
        | cpp \
        | grep -o 'struct[[:space:]]\+\w\+' \
        | cut -d" "  -f2 \
        | sort -u \
        | sed -r 's/^(.+)$/\t\u\1\tC.struct_\1/' -)
)
EOF
};
f | gofmt > types.go

