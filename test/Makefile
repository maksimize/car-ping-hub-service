build:
	docker run \
	--volume="$(PWD)/app:/usr/src/app" \
    --workdir=/usr/src/app \
    --env=GOOS=linux \
    --env=GOARCH=amd64 \
    golang bash -c "go get -d ./... && go build . && ./app"

run:
	docker run \
	--volume="$(PWD)/app:/usr/src/app" \
    --workdir=/usr/src/app \
    slim ./app
