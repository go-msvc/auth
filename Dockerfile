FROM golang:1.19 AS builder

WORKDIR /opt
COPY service/*.go ./service/
COPY db/*.go ./db/
COPY go.* ./
COPY ./*.go ./
WORKDIR /opt/service/
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -o /opt/service/service

FROM alpine:latest
WORKDIR /root
COPY --from=builder /opt/service/service .
COPY ./service/config.json /root/conf/config.json
CMD [ "/root/service" ]
