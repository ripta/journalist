FROM golang:1.22-bookworm AS builder

RUN apt-get update && apt-get install -y git
RUN git clone https://github.com/mrusme/journalist /app
WORKDIR /app
RUN go build -o /usr/bin/journalist .

FROM debian:bookworm-slim

RUN groupadd --gid 1683 journalist \
  && useradd --uid 1683 --gid 1683 -m journalist

USER journalist
WORKDIR /home/journalist
COPY --from=builder /usr/bin/journalist /usr/bin/journalist
CMD ["journalist"]
