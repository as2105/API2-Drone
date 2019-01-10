FROM golang:alpine

RUN apk --no-cache add alpine-sdk
ARG gittoken
RUN git config --global url."https://${gittoken}:@github.com/".insteadOf "https://github.com/"
RUN go get -u github.com/gobuffalo/packr/v2/packr2
ENV GO111MODULE=on
WORKDIR /app
COPY ./ /app
RUN make

FROM alpine
COPY --from=0 /app/app /usr/bin/fhirapi
EXPOSE 8080

ENTRYPOINT [ "fhirapi" ]
