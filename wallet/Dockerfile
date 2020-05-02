FROM golang:1.13

WORKDIR /go/src/ego-api/wallet

RUN addgroup --system wallet && adduser --system wallet --ingroup wallet

RUN chown -R wallet:wallet /go/src/ego-api/wallet

USER wallet

COPY . .

RUN go get -d -v ./...

#RUN go install -v ./...

#ENTRYPOINT [ "wallet" ]
