FROM  golang:latest as build-env
ENV GOPROXY=https://goproxy.cn,direct CGO_ENABLED=0
WORKDIR /home/docker
COPY . .
RUN go build -o api .

FROM alpine:latest
LABEL MAINTAINER="kcoewoys"
EXPOSE 8080
COPY --from=build-env /home/docker/api /home/docker/
CMD ["/home/docker/api"]