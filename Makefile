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

.PHONY: build-migrate
build-migrate:
	go build -o migrate cmd/migrate/main.go

.PHONY: build-webpack
build-webpack:
	cd server/assets && \
	npm run build

.PHONY: docker-build
docker-build:
	docker build --target go-blog -t go-blog .
	docker tag go-blog:latest ${AWS_CONTAINER_REPOSITORY_URL}

.PHONY: docker-login
docker-login:
	aws ecr get-login-password | docker login --username AWS --password-stdin ${AWS_ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com

.PHONY: docker-push
docker-push:
	docker push ${AWS_CONTAINER_REPOSITORY_URL}

.PHONY: update-service
update-service:
	aws ecs update-service --region ${AWS_REGION} --cluster ${AWS_ECS_CLUSTER} --service ${AWS_ECS_SERVICE} --force-new-deployment

.PHONY: wait-stable
wait-stable:
	aws ecs wait services-stable --region ${AWS_REGION} --cluster ${AWS_ECS_CLUSTER} --services ${AWS_ECS_SERVICE}

.PHONY: deploy
deploy: docker-build docker-login docker-push update-service wait-stable
