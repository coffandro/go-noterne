build-web:
	mkdir -p build
	cd src-web && \
	GOOS=js GOARCH=wasm go build -o ../build/main.wasm .

build-desktop:
	mkdir -p build
	cd src && \
	go build -o ../build/main .
