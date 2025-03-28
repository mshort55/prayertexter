VERSION=$(shell git describe --tags --long --dirty 2>/dev/null)

### we must have tagged the repo at least once for VERSION to work
ifeq ($(VERSION),)
	VERSION = UNKNOWN
endif

buildcmd = GOARCH=amd64 GOOS=linux go build -ldflags "-X main.version=${VERSION}" -o bootstrap && mv bootstrap ../../

build-announcer:
	(cd cmd/announcer && $(buildcmd))

build-prayertexter:
	(cd cmd/prayertexter && $(buildcmd))

build-stateresolver:
	(cd cmd/stateresolver && $(buildcmd))