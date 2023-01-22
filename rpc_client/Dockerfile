FROM structs AS builder

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o client

FROM scratch

COPY --from=builder /src/client .

CMD ["./client"]