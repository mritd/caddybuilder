BUILD_VERSION   := $(shell cat version)
BUILD_DATE      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD)

all:
	gox -osarch="darwin/amd64 linux/386 linux/amd64 windows/386 windows/amd64" \
		-output="dist/{{.Dir}}_{{.OS}}_{{.Arch}}" \
		-ldflags	"-X 'main.Version=${BUILD_VERSION}' \
					-X 'main.BuildDate=${BUILD_DATE}' \
					-X 'main.CommitID=${COMMIT_SHA1}'"

docker:
	docker build -t mritd/caddybuilder:${BUILD_VERSION} .

clean:
	rm -rf dist

install:
	go install -ldflags	"-X 'main.Version=${BUILD_VERSION}' \
						-X 'main.BuildDate=${BUILD_DATE}' \
						-X 'main.CommitID=${COMMIT_SHA1}'"

release: all
	ghr -u mritd -t $(GITHUB_RELEASE_TOKEN) -replace -recreate --debug ${BUILD_VERSION} dist

pre-release: all
	ghr -u mritd -t $(GITHUB_RELEASE_TOKEN) -replace -recreate -prerelease --debug ${BUILD_VERSION} dist


.PHONY : all release docker clean install

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
