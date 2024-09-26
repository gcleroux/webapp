FROM golang:1.22.7-alpine AS build

WORKDIR /build

COPY . .

RUN go mod download && \
    CGO_ENABLED=0 go build -o webapp ./main.go

FROM scratch

COPY --from=build /build/webapp /

ENTRYPOINT [ "/webapp" ]
