all: test build

.PHONY: test
test:
	@gofmt -l . > fmt.result
	@if [[ -s fmt.result ]]; then\
		echo "Code not formatted:";\
		echo fmt.result;\
		exit 1;\
	fi
	@rm -f ./fmt.result
	@go test -coverprofile=coverage.out ./...

.PHONY:build
build:
	@CGO_ENABLED=0 go build -o bin/retagger retagger.go

install:./bin/retagger
	@mv ./bin/retagger /usr/local/bin/retagger
	@chmod +x /usr/local/bin/retagger

uninstall:/usr/local/bin/retagger
	@rm -f /usr/local/bin/retagger

.PHONY:build_artifacts
build_artifacts: build_linux build_windows build_darwin

.PHONY: build_linux
build_linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o bin/retagger retagger.go
	@cd bin;tar -czf retagger_linux_amd64.tar.gz retagger;rm -f retagger

.PHONY: build_windows
build_windows:
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -a -o bin/retagger retagger.go
	@cd bin;tar -czf retagger_windows_amd64.tar.gz retagger;rm -f retagger

.PHONY: build_darwin
build_darwin:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -o bin/retagger retagger.go
	@cd bin;tar -czf retagger_darwin_amd64.tar.gz retagger;rm -f retagger