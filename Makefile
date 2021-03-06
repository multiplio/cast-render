project=multipl
name=cast-render

.PHONY:all
all: build image

.PHONY:image
image:
	docker image build \
		-t ${project}/${name}:latest \
		.

.PHONY:build
build:
	go build ./...

.PHONY:test
test: build
	go test

.PHONY:run
run:
	docker container run \
		--rm \
		--name ${project}-${name}-dev \
		-p 3000:3000 \
		--env-file .env \
		-t ${project}/${name}:latest

.PHONY:kill
kill:
	docker kill $$( \
		docker ps -aq \
			--filter="name=${project}-${name}-dev" )

