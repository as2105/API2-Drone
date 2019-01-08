FROM golang:alpine

RUN apk --no-cache add alpine-sdk
ENV GO111MODULE=on
WORKDIR /app
COPY ./ /app
RUN go build -o fhirapi

FROM alpine
COPY --from=0 /app/fhirapi /usr/bin
EXPOSE 8000

ENTRYPOINT [ "fhirapi" ]
