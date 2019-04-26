FROM golang:alpine as builder

LABEL AUTHOR="Suhail Abood" \
      DESCRIPTION="Go builder for business service" \
      USAGE="docker build -t suhdev/nutrient-business ."

COPY ./build.sh build.sh
RUN chmod +x build.sh && ./build.sh
WORKDIR $GOPATH/src/github.com/suhdev/nutrient-business
COPY . ./
RUN glide install
RUN go build main.go && chmod +x main

CMD ["./main"]