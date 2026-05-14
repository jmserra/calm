BINARY := calm
BINDIR ?= $(HOME)/doc/bin

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
	install -d $(BINDIR)
	install -m 0755 $(BINARY) $(BINDIR)/$(BINARY)

.PHONY: uninstall
uninstall:
	rm -f $(BINDIR)/$(BINARY)

.PHONY: clean
clean:
	rm -f $(BINARY)
