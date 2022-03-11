FROM golang:1.17-alpine
WORKDIR /impl-api/
ADD . /impl-api
RUN cd /impl-api && go build
EXPOSE 8080
ENTRYPOINT ["./impl-api"]