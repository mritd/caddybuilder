FROM golang:1.12.5-alpine3.9 AS builder

COPY . /go/src/github.com/mritd/caddybuilder

WORKDIR /go/src/github.com/mritd/caddybuilder

ENV GO111MODULE on
ENV GOPROXY https://athens.azurefd.net

RUN apk upgrade \
    && apk add git \
    && BUILD_VERSION=$(cat version) \
    && BUILD_DATE=$(date "+%F %T") \
    && COMMIT_SHA1=$(git rev-parse HEAD) \
    && go install -ldflags  "-X 'main.Version=${BUILD_VERSION}' \
                            -X 'main.BuildDate=${BUILD_DATE}' \
                            -X 'main.CommitID=${COMMIT_SHA1}'"

FROM golang:1.12.5-alpine3.9 AS dist

LABEL maintainer="mritd <mritd@linux.com>"

RUN apk upgrade \
    && apk add bash git --no-cache

COPY --from=builder /go/bin/caddybuilder /usr/bin/caddybuilder

ENTRYPOINT ["caddybuilder"]

CMD ["--help"]