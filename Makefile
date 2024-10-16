test:
	if [ -x "$(command -v docker)" ]; then \
		docker build -t test .; \
	else \
		finch build -t test .; \
	fi
