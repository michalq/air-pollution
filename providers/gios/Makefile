.PHONY: all clean build view

all: clean build

build:
	swagger generate server -f ./swagger.json -A gios-api-client

clean:
	rm -rf client
	rm -rf models

view:
	swagger serve ./swagger.json