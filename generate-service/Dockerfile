# Compile stage
FROM golang:1.14.3

ENV GOPROXY=https://goproxy.cn
ENV GO111MODULE=on

WORKDIR /generate-service

ADD ../go.mod .
ADD ../go.sum .
RUN go mod download

ADD generate-service /generate-service

WORKDIR /generate-service/cmd/main

RUN GOOS=linux CGO_ENABLED=0  go build -ldflags="-s -w" -installsuffix cgo -o app main.go

FROM scratch

COPY --from=0 /generate-service/config.yml /
COPY --from=0 /generate-service/cmd/main/app /

EXPOSE 8888

CMD ["/app"]