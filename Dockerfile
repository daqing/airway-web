# 构建 Dashboard
FROM node:18-alpine as tailwind

# 设置工作目录
WORKDIR /code
COPY . .

# 安装 pnpm
RUN npm config set registry https://registry.npmmirror.com
RUN npm install -g pnpm

RUN pnpm install
RUN pnpx tailwindcss -i ./app/assets/stylesheets/application.css -o application.css

FROM golang:1.22-alpine as builder
WORKDIR /airway
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o ./bin/server .

FROM alpine
WORKDIR /web
RUN mkdir app
RUN mkdir -p public/assets
RUN mkdir -p public/stylesheets
RUN mkdir static

COPY --from=builder /airway/bin/server /web
COPY --from=builder /airway/app/views /web/app/views

COPY --from=tailwind /code/application.css /web/public/stylesheets/application.css
COPY --from=tailwind /code/public/assets /web/public/assets

COPY --from=tailwind /code/static /web/static

ENV AIRWAY_ENV=production
ENV AIRWAY_PORT=1999
ENV AIRWAY_ROOT=/web
ENV TZ="Asia/Shanghai"

EXPOSE 1999

CMD ["/web/server"]
