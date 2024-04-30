FROM golang:1.22-bookworm AS builder

RUN apt-get update && apt-get install -y git
RUN git clone https://github.com/mrusme/journalist /app
WORKDIR /app
# disable CGO, which also prevents the use of sqlite3
ENV CGO_ENABLED=0
RUN go build -o /usr/bin/journalist .


FROM gcr.io/distroless/static-debian12:nonroot-amd64

WORKDIR /home/nonroot
COPY --from=builder /usr/bin/journalist /usr/bin/journalist
ENTRYPOINT ["/usr/bin/journalist"]
