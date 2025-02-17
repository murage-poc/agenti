FROM golang:1.24.0-alpine AS build

WORKDIR /app

RUN apk add make

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN BUILD_OUTPUT_FILENAME=agenti && make build


FROM scratch AS release

WORKDIR /app

COPY --from=build /app/dist/api/agenti ./agenti

EXPOSE 8080

ENTRYPOINT ["/app/agenti"]