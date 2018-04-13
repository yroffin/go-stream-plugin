all: bin/cmd bin/plugin-producer.so
	cd bin && ./cmd

bin/plugin-producer.so: build
	cd plugin-producer && go build -buildmode=plugin && ls -lrt
	mv -f plugin-producer/plugin-producer.so bin

bin/cmd: build
	cd cmd && go build && ls -lrt
	mv -f cmd/cmd bin

build:

