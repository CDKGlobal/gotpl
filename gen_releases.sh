#!/bin/bash

OS_S=("linux" "darwin")
ARCH_S=("amd64" "386")

for GOOS in ${OS_S[@]} ; do
  for ARCH in ${ARCH_S[@]} ; do
    GOOS=${GOOS} \
      GOARCH=${ARCH} \
      go build -o releases/gotpl_${GOOS}_${ARCH}
  done
done
