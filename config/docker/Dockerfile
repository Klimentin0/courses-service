# Go Binary
FROM golang:1.23.4 as build_courses-api
ENV CGO_ENABLED 0
ARG BUILD_REF

# Copy the source code into container.
COPY . /courses-service

# Build the service binary
WORKDIR /courses-service
RUN go build -ldflags "-X main.build=${BUILD_REF}"

# Run the go Binary in Alpine.
FROM alpine:3.14
ARG BUILD_DATE
ARG BUILD_REF
COPY --from=build_courses-api /courses-service /courses-service/courses-service
WORKDIR /courses-service
CMD ["./courses-service"]

LABEL org.opencountainers.image.created="${BUILD_DATE}" \
    org.opencountainers.image.title="courses-api" \
    org.opencountainers.image.authors="Klimentii Meshkov <meshkov.biz@yandex.ru" \
    org.opencountainers.image.source="https://github.com/Klimentin0/courses-service" \
    org.opencountainers.image.revision="${BUILD_REF}}" \
    org.opencountainers.image.vendor="Klimentii" \
