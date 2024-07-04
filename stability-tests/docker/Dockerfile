ARG SAD_IMAGE
ARG SADMINER_IMAGE

FROM ${SAD_IMAGE} as sad
FROM ${SADMINER_IMAGE} as sadminer

FROM golang:1.19-alpine

RUN mkdir -p /go/src/github.com/sadnetwork/sad

WORKDIR /go/src/github.com/sadnetwork/sad

RUN apk add bash build-base git

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

COPY --from=sad /app/ /app/
COPY --from=sadminer /app/ /app/
ENV PATH="/app:${PATH}"

WORKDIR /go/src/github.com/sadnetwork/sad/stability-tests

RUN go install ./...

ENTRYPOINT ["./run/run.sh"]
