SOURCES := main.go logger.go
BINARY:= bin/app
PLATFORMS := darwin/amd64 darwin/arm64 linux/386 linux/amd64 linux/arm linux/arm64
WIN_PLATFORMS := windows/386 windows/amd64 windows/arm windows/arm64

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

.PHONY : release $(PLATFORMS) $(WIN_PLATFORMS) clean

release: $(PLATFORMS) $(WIN_PLATFORMS)

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o $(BINARY)-$(os)-$(arch) $(SOURCES)

$(WIN_PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o $(BINARY)-$(os)-$(arch).exe $(SOURCES)

build:
	go build -o $(BINARY) $(SOURCES)

clean:
	rm -rf $(BINARY)-*

ci:
	go build -v -o ./app $(SOURCES)

dev:
	go build -o $(BINARY)-dev $(SOURCES)

docker:
	go build -v -o /usr/local/bin/igdb-app $(SOURCES)

run:
	go build -o $(BINARY) $(SOURCES)
	./$(BINARY)-dev