fNameSrc = main
fNameOut = main

# 64 bit x86 setting
.64Bitx86:
	go env -w GOARCH=amd64

# Windows settings
.windows:
	go env -w GOOS=windows

# Linux settings
.linux:
	go env -w GOOS=linux

# Mac OS X 10.8 and above
.mac:
	go env -w GOOS=darwin

# Default value
.default: .64Bitx86 .windows


# Build for 64 bit windows (default)
windows: .default build

# Build for 64 bit linux
linux: .64Bitx86 .linux build .default

# Build for 64 bit mac
mac: .64Bitx86 .mac build .default


# Builder command
build:
	go build -o ./${fNameSrc}.exe ./${fNameOut}.go

b: build

# Runs go.main
run:
	go run ./${fNameSrc}.go

r: run