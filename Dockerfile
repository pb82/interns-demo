FROM registry.access.redhat.com/ubi8/ubi-minimal:8.5 AS builder

RUN microdnf install -y tar gzip make which

# install go 1.17.8
RUN curl -O -J https://dl.google.com/go/go1.17.8.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.17.8.linux-amd64.tar.gz
RUN ln -s /usr/local/go/bin/go /usr/local/bin/go

WORKDIR /workspace

COPY . ./
RUN go mod tidy
RUN go mod vendor
RUN go build main.go

FROM registry.access.redhat.com/ubi8/ubi-minimal:8.5
COPY --from=builder /workspace/main /usr/local/bin/
EXPOSE 8000
ENTRYPOINT ["/usr/local/bin/main"]
