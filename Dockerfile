FROM node:18-alpine as node-builder

WORKDIR /node
COPY . .

RUN npm config set registry https://registry.npmmirror.com
RUN npm install -g pnpm

RUN pnpm install
RUN pnpx tailwindcss -i ./app/assets/stylesheets/application.css -o application.css


FROM golang:1.22-alpine as go-builder
WORKDIR /go
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o ./bin/airway-web .

FROM alpine
WORKDIR /airway

RUN mkdir app
RUN mkdir -p public/assets
RUN mkdir -p public/stylesheets

COPY --from=node-builder /node/application.css ./public/stylesheets/application.css
COPY --from=node-builder /node/public/assets ./public/assets

COPY --from=go-builder /go/bin/airway-web ./server
COPY --from=go-builder /go/app/views ./app/views

ENV AIRWAY_ENV=production
ENV AIRWAY_PORT=1999
ENV AIRWAY_ROOT=/app
ENV TZ="Asia/Shanghai"

EXPOSE 1999

CMD ["/airway/server"]
