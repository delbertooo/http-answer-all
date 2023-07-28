FROM golang:1.20-alpine3.18 as build

WORKDIR /app
COPY go.* /app/
COPY cmd/ /app/cmd/

RUN GOARCH="$TARGETARCH" go build -o /app/http-answer-all ./cmd


FROM alpine:3.18

COPY --from=build /app/http-answer-all /usr/local/bin/http-answer-all

USER nobody

CMD [ "http-answer-all" ]
