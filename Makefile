PREFIX    ?= /usr
BINPREFIX ?= $(PREFIX)/bin

default: build

install_deps:
	dep ensure
clean_build:
	rm -rf build
build: clean_build install_deps
	mkdir build
	go build -o build/powerd main.go battery_info.go config.go daemon.go
install:
	mkdir -p "$(DESTDIR)$(BINPREFIX)"
	install -m 755 "build/powerd" "$(BINPREFIX)"
uninstall:
	rm -f "$(BINPREFIX)/powerd"
clean: clean_build
