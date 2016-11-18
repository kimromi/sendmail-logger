#!/bin/bash
rm -rf pkg/
gox -output "pkg/{{.Dir}}_{{.OS}}_{{.Arch}}"
