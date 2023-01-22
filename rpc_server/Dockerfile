FROM structs AS builder

COPY . /src

WORKDIR /src

RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM scratch

COPY --from=builder /src/server .

CMD ["./server"]