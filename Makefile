DB="user=musicbrainz dbname=musicbrainz host=127.0.0.1 sslmode=disable"

APP_NAME=microphone
PORT=7000
HOST=127.0.0.1
GO_CMD=go
GO_TEST=DB=$(DB) $(GO_CMD) test -v
GO_BUILD=$(GO_CMD) build -v
RUN=DB=$(DB) PORT=$(PORT) HOST=$(HOST) ./$(APP_NAME)
BUILD_AND_RUN=$(GO_BUILD) && $(RUN)

all: build
test: RunTests
build: BuildApp
run: RunApp

BuildApp:RunTests
	$(GO_BUILD)

RunTests:
	$(GO_TEST)

RunApp:
	$(BUILD_AND_RUN)
