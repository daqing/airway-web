FROM golang:1.22-alpine as builder
WORKDIR /app
COPY . /app
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o ./bin/airway-web .

FROM alpine
WORKDIR /app
COPY --from=builder /app/bin/airway-web /app

ENV AIRWAY_ENV=production
ENV AIRWAY_PORT=1999
ENV AIRWAY_ROOT=/app
ENV TZ="Asia/Shanghai"

EXPOSE 1999

CMD ["/app/airway-web"]
