default: build

clean_build:
	@echo "> cleaning build/*"
	@rm -rf build
build: clean_build
	@echo "> building 'powerd' executable in build/"
	@ mkdir build
	@go build -o build/powerd main.go battery_info.go config.go daemon.go
clean: clean_build
