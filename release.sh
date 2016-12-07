#!/bin/bash
rm -rf pkg/
gox -output "pkg/{{.Dir}}_{{.OS}}_{{.Arch}}"
ghr v0.0.3 pkg/
