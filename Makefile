EXECUTABLE_FILE = slack-post

.PHONY: build
build:
	go build -o ${EXECUTABLE_FILE}

.PHONY: clean
clean:
	rm ${EXECUTABLE_FILE}
