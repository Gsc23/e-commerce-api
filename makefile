run: build
	@echo "Running the application..."

	@ ./tmp/main

build:
	@echo "Building the application..."

	@ go build \
			-trimpath  \
			-buildvcs=false \
			-o tmp/main \
			./cmd/api

	@echo "Build completed."