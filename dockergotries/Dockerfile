# Stage 1: Compile the binary in a containerized Golang environment
#
FROM golang:1.19 AS builder

# Copy the source files from the host
COPY . /src

# Set the working directory to the same place we copied the source files
WORKDIR /src

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -o test

# Stage 2: Build the test image proper
#
# Use a "scratch" image, which contains no distribution files
FROM scratch

# Copy the binary from the builder image
COPY --from=builder /src/test .

# Tell docker to execute the binary when the container starts
CMD ["./test"]



