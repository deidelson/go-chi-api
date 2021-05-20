FROM golang:1.15 as build
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main cmd/main.go

FROM ubuntu
COPY --from=build /app/main /main
EXPOSE 8080
CMD ["./main"]
