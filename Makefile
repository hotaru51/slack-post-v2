EXECUTABLE_FILE = slack-post

.PHONY: build
build:
	go build -o ${EXECUTABLE_FILE}

.PHONY: linux_amd64
linux_amd64:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o ${EXECUTABLE_FILE} -ldflags '-extldflags=-static'

.PHONY: linux_arm
linux_arm:
	GOOS=linux GOARCH=arm CGO_ENABLED=0 go build -o ${EXECUTABLE_FILE} -ldflags '-extldflags=-static'

.PHONY: clean
clean:
	rm ${EXECUTABLE_FILE}
