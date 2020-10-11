FROM golang:1.15.2-alpine as go-builder
WORKDIR /go/src/github.com/YusukeKishino/go-blog
RUN apk add --no-cache make git && \
    rm -rf /var/cache/apk/*

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN make build
RUN make build-migrate

FROM node:14.5.0-alpine as node-builder
WORKDIR /app
ENV NODE_ENV=production
COPY server/assets/package.json server/assets/package-lock.json server/assets/webpack.config.js ./
RUN npm install
COPY server/assets/ ./
RUN npm run build

FROM alpine:3.10 as go-blog
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*
RUN ENTRYKIT_VERSION=0.4.0 && \
    wget -q https://github.com/progrium/entrykit/releases/download/v${ENTRYKIT_VERSION}/entrykit_${ENTRYKIT_VERSION}_Linux_x86_64.tgz && \
    tar -xvzf entrykit_${ENTRYKIT_VERSION}_Linux_x86_64.tgz && \
    rm entrykit_${ENTRYKIT_VERSION}_Linux_x86_64.tgz && \
    mv entrykit /bin/entrykit && \
    chmod +x /bin/entrykit && \
    entrykit --symlink

COPY --from=go-builder /go/src/github.com/YusukeKishino/go-blog/main /usr/local/bin/server
COPY --from=go-builder /go/src/github.com/YusukeKishino/go-blog/migrate /usr/local/bin/
COPY server/views server/views
COPY config/settings.yml config/
COPY --from=node-builder /app/public server/assets/public
COPY --from=node-builder /app/src/images server/assets/src/images
EXPOSE 3000/tcp
ENTRYPOINT [ \
  "switch", \
  "server=/usr/local/bin/server -a :3000", \
  "migrate=/usr/local/bin/migrate", \
  "--", \
  "/bin/sh" \
]
