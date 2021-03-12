fNameSrc = main
fNameOut = main

# Windows settings
.windows:
	go env -w GOARCH=amd64
	go env -w GOOS=windows

# Linux settings
.linux:
	go env -w GOARCH=amd64
	go env -w GOOS=linux

# Build for windows (default)
windows: .windows build

# Build for linux
linux: .linux build .windows

#  Builder command
build:
	go build -o ./${fNameSrc}.exe ./${fNameOut}.go

b: windows

# Runs go.main with windows as default
run: .windows
	go run ./${fNameSrc}.go

r: run

# Runs executable
execute:
	./bin/${fNameOut}
