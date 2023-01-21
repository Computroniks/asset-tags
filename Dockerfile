# SPDX-FileCopyrightText: 2023 Matthew Nickson <mnickson@sidingsmedia.com>
# SPDX-License-Identifier: MIT

FROM golang:1.16-alpine as build

## Build
WORKDIR /build

COPY go.mod /build
COPY go.sum /build

# Download go modules
RUN go mod download

# Copy all files
COPY . /build

# Compile binary
RUN CGO_ENABLED=0 go build -a -o asset-tags

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /build/asset-tags /asset-tags

EXPOSE 3000/tcp

USER nonroot:nonroot

ENTRYPOINT ["/asset-tags"]
