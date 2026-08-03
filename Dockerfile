FROM public.ecr.aws/docker/library/golang:1.26.5@sha256:3aff6657219a4d9c14e27fb1d8976c49c29fddb70ba835014f477e1c70636647

RUN apt-get update -q && \
  apt-get install -y zip

WORKDIR /go/src/github.com/buildkite/yaml2json
ADD . /go/src/github.com/buildkite/yaml2json

CMD [ "scripts/build" ]
