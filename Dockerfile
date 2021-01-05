FROM golang:1.16beta1

RUN apt-get update -q && \
  apt-get install -y zip && \
  go get github.com/buildkite/github-release

WORKDIR /go/src/github.com/buildkite/yaml2json
ADD . /go/src/github.com/buildkite/yaml2json

ENV GO111MODULE=on

CMD [ "scripts/build" ]
