BINARY := calm
PREFIX ?= $(HOME)/.local

.PHONY: build
build:
	go build -ldflags="-s -w" -o $(BINARY) .

.PHONY: run
run: build
	./$(BINARY)

.PHONY: vet
vet:
	go vet ./...

.PHONY: install
install: build
	install -d $(PREFIX)/bin
	install -m 0755 $(BINARY) $(PREFIX)/bin/$(BINARY)

.PHONY: uninstall
uninstall:
	rm -f $(PREFIX)/bin/$(BINARY)

.PHONY: clean
clean:
	rm -f $(BINARY)
