test:
	if [ -x "$(command -v docker)" ]; then \
		docker build -t test .; \
		docker run test; \
	else \
		finch build -t test .; \
		finch run test; \
	fi
