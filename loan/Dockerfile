FROM golang:1.13

WORKDIR /go/src/ego-api/loan

RUN addgroup --system loan && adduser --system loan --ingroup loan

RUN chown -R loan:loan /go/src/ego-api/loan

USER loan

COPY . .

RUN go get -d -v ./...

#RUN go install -v ./...

#ENTRYPOINT [ "wallet" ]
