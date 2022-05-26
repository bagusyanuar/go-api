# FROM golang:1.17-alpine as build-env
 
# ENV APP_NAME go-api
# ENV CMD_PATH main.go
 
# COPY . $GOPATH/src/$APP_NAME
# WORKDIR $GOPATH/src/$APP_NAME
 
# RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
 
# FROM alpine:3.14
 
# ENV APP_NAME go-api
 
# COPY --from=build-env /$APP_NAME .
 
# EXPOSE 8000
 
# CMD ./$APP_NAME

FROM golang:1.17-alpine as build-env
 
ENV APP_NAME /go/src/go-api
ENV CMD_PATH main.go
 
WORKDIR $APP_NAME

COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o go-api
 
FROM alpine:3.14
 
ENV APP_NAME /go/src/go-api

RUN mkdir -p $APP_NAME
WORKDIR $APP_NAME

COPY views views/
COPY .env ./
 
COPY --from=build-env $APP_NAME/go-api $APP_NAME
 
EXPOSE 8000
 
CMD ./go-api