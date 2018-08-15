all: milmap deps jsx

milmap: main.go
	CGO_ENABLED=0 go build -a -installsuffix cgo -ldflags '-s' .

deps: js/package.json 
	cd js; yarn install; cd ..

SOURCES := $(shell find js/src -name '*')

jsx: $(SOURCES)
	cd js; webpack; cd ..
