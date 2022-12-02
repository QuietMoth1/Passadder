build:
	go build -o passadder passadder.go

run:
	go run passadder.go

compile:
    # 32-Bit Systems
    # FreeBDS
    GOOS=freebsd GOARCH=386 go build -o passadder passadder.go
    # MacOS
    GOOS=darwin GOARCH=386 go build -o passadder passadder.go
    # Linux
    GOOS=linux GOARCH=386 go build -o passadder passadder.go
    # Windows
    GOOS=windows GOARCH=386 go build -o passadder.exe passadder.go
        # 64-Bit
    # FreeBDS
    GOOS=freebsd GOARCH=amd64 go build -o passadder passadder.go
    # MacOS
    GOOS=darwin GOARCH=amd64 go build -o passadder passadder.go
    # Linux
    GOOS=linux GOARCH=amd64 go build -o passadder passadder.go
    # Windows
    GOOS=windows GOARCH=amd64 go build -o passadder.exe passadder.go