FROM golang:1.22.2-alpine3.19 as stage_build

WORKDIR /app

ADD . /app
RUN go mod download

RUN go build

FROM alpine:3.19.0

WORKDIR /app

COPY --from=stage_build /app/backendtest-go /app/
COPY --from=stage_build /app/.env /app/

ENTRYPOINT [ "./backendtest-go" ]

EXPOSE 3000:3000