FROM golang:1.12.7
ENV GO111MODULE on
COPY ./oauth ./oauth
CMD ["./oauth", "--http-port", "8080"]