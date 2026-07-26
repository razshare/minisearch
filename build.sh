#!/usr/bin/env bash
frizzante package
GOOS=windows GOARCH=amd64 frizzante build --output=".gen/bin/windows-amd64" no-package &\
GOOS=linux GOARCH=amd64 frizzante build --output=".gen/bin/linux-amd64" no-package &\
GOOS=linux GOARCH=arm64 frizzante build --output=".gen/bin/linux-arm64" no-package &\
GOOS=darwin GOARCH=amd64 frizzante build --output=".gen/bin/darwin-amd64" no-package &\
GOOS=darwin GOARCH=arm64 frizzante build --output=".gen/bin/darwin-arm64" no-package &\
wait