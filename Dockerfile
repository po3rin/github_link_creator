FROM golang:1.11 as builder

WORKDIR /api

COPY . .

ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /api .

ARG pip_installer="https://bootstrap.pypa.io/get-pip.py"

RUN apk --update add \
    python \
    curl \
    groff \
    bash \
    dpkg \
    gcc \
    git \
    musl-dev \
    openssh \
    bash \
    sudo

RUN curl ${pip_installer} | python && \
    pip install awscli

RUN sudo chmod +x ./docker-entrypoint.sh
ENTRYPOINT ["./docker-entrypoint.sh"]
CMD ["./github_link_creator"]
