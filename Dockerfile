FROM golang:1.15 as build
WORKDIR /krew-museum
COPY src/ .
RUN apt-get update && apt-get upgrade -y \
  && mkdir bin/ \
  && go get github.com/gorilla/mux \
  && CGO_ENABLED=0 GOOS=linux go build -tags netgo -a -v -o bin/krew-museum src/*

FROM alpine:latest
COPY --from=build /krew-museum/bin/krew-museum krew-museum
RUN chmod +x /krew-museum
EXPOSE 8080
ENTRYPOINT ["/krew-museum"]