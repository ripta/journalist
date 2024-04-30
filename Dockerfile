FROM golang:1.22-bookworm AS builder

RUN apt-get update && apt-get install -y git
RUN git clone https://github.com/mrusme/journalist /app
WORKDIR /app
RUN go build -o /usr/bin/journalist .


FROM gcr.io/distroless/base-debian12:debug-nonroot-amd64

WORKDIR /journalist
COPY --from=builder /usr/bin/journalist /usr/bin/journalist
ENTRYPOINT ["journalist"]
