FROM golang:1.20-alpine3.18 as build

WORKDIR /app
COPY main.go /app/

RUN GOARCH="$TARGETARCH" go build -o /app/http-answer-all main.go

FROM alpine:3.18

COPY --from=build /app/http-answer-all /usr/local/bin/http-answer-all

CMD [ "http-answer-all" ]
