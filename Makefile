linux:
	GO111MODULE=on GOOS=linux GOARCH=amd64 go build -o build/utils-linux-amd64 ./
	GO111MODULE=on GOOS=linux GOARCH=386 go build -o build/utils-linux-386 ./