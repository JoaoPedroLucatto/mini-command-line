FROM golang:1.19 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY  . .

RUN GCO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /line-code /app/main.go

#docker build . -t my-line-code 
#docker run  my-line-code /line-code server --host amazon.com.br


#FROM gcr.io/distroless/static-debian11

#COPY --from=build /line-code /line-code

CMD ["/line-code"]