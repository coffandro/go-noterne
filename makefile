.PHONY: watch-web

build-web:
	mkdir -p build
	cd src-web && \
	GOOS=js GOARCH=wasm go build -o ../build/main.wasm .

watch-web:
	@echo "Watching ./internal for changes..."
	@while true; do \
		inotifywait -r -e modify,create,delete ./internal; \
		make build-web; \
	done

build-desktop:
	mkdir -p build
	cd src && \
	go build -o ../build/main .

run-desktop: build-desktop
	build/main
