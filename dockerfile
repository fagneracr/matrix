FROM golang:latest
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
COPY ./internal/ginserver/* ./internal/ginserver/
COPY ./internal/verifymatrix/* ./internal/verifymatrix/
COPY ./internal/database/* ./internal/database/
RUN go test ./internal/verifymatrix/ -race 
RUN go test ./internal//database/ -race 
RUN go build -o /matrixtest
ENV hostredis=redis
ENV portredis=6379
EXPOSE 3001
CMD [ "/matrixtest" ]