build:
	go build -o build/git-lab

clean:
	rm -rf build/
test:
	go test ./...
