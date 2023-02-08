FROM golang:alpine

WORKDIR /build

COPY . .

RUN go build -o image_crop image_crop.go

CMD ["./image_crop"]

EXPOSE 3000
