FROM golang:1.20

RUN apt-get update -y

RUN apt-get install -y \
    vim

RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s v1.51.2

ARG SENTRY_RELEASE
ENV SENTRY_RELEASE=${SENTRY_RELEASE:-}

WORKDIR /app
COPY . .
RUN go install ./cmd/uffizzi/...

CMD ["bash", "-c", "/go/bin/uffizzi"]
