default:
	env CGO_ENABLED=0 GOOS=linux go build -o post-process-linux main.go
	env CGO_ENABLED=0 GOOS=darwin go build -o post-process-darwin main.go
	env CGO_ENABLED=0 GOOS=windows go build -o post-process.exe main.go
