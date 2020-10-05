.PHONY: setup
setup: mod-tidy npm-install build-webpack bootstrap

.PHONY: mod-tidy
mod-tidy:
	go mod tidy

.PHONY: npm-install
npm-install:
	cd server/assets && \
	npm install

.PHONY: bootstrap
bootstrap:
	go get github.com/codegangsta/gin

.PHONY: dev-server
dev-server:
	gin -i -a 8080 --all -x server/assets -d cmd/server

.PHONY: webpack-dev-server
webpack-dev-server:
	cd server/assets && \
	npm run dev

.PHONY: build
build:
	go build cmd/server/*.go

.PHONY: build-webpack
build-webpack:
	cd server/assets && \
	npm run build
