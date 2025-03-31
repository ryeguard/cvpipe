.PHONY: build

build:
	if [ -x "$(command -v docker)" ]; then \
		docker buildx create --use --name multi-platform-builder || true; \
		docker buildx build --platform linux/amd64,linux/arm64 -t cvpipe-test .; \
	else \
		finch build --platform linux/arm64 -t cvpipe-test .; \
	fi