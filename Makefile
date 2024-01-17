BUILD_ENV = docker build -t speech-recognition-service:dev -f docker/Dockerfile .
RUN_ENV = docker run -it speech-recognition-service:dev /bin/bash

build-env:
	$(BUILD_ENV)

run-env:
	$(RUN_ENV)

start-env:
	$(BUILD_ENV)
	$(RUN_ENV)