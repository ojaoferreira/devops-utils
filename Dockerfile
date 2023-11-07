FROM golang:1.19.4 AS build

WORKDIR /app

COPY . ./

RUN go build -o /app/devops-utils

FROM gcr.io/distroless/base-debian11:latest

COPY --from=build /app/devops-utils /app/devops-utils

USER nonroot

WORKDIR /app

ENTRYPOINT [ "./devops-utils" ]