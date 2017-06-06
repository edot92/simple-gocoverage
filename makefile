
all: build

build: app

app:
	env  GOOS=darwin GOARCH=amd64 go build  -o ./build/darwin_amd64/appedy
	env  GOOS=linux GOARCH=386 go build  -o ./build/linux_386/appedy
	env  GOOS=linux GOARCH=amd64 go build  -o  ./build/linux_amd64/appedy
	env  GOOS=linux GOARCH=arm go build  -o ./build/linux_arm/appedy
	env  GOOS=windows GOARCH=386 go build  -o ./build/windows_386/appedy.exe
	env  GOOS=windows GOARCH=amd64 go build  -o ./build/windows_amd64/appedy.exe
	env  GOOS=linux GOARCH=mips64 go build  -o ./build/linux_mips64/appedy
	env  GOOS=linux GOARCH=mips64le go build  -o ./build/linux_mips64le/appedy
	env  GOOS=linux GOARCH=mips go build  -o ./build/linux_mips/appedy
	env  GOOS=linux GOARCH=mipsle go build  -o ./build/linux_mipsle/appedy
