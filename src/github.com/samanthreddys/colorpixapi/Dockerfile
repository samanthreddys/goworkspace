FROM golang:latest
RUN mkdir -p /go/src/colorpixapi
WORKDIR /go/src/colorpixapi

COPY . /go/src/colorpixapi

RUN  go install colorpixapi

CMD /go/bin/colorpixapi

EXPOSE 8080

