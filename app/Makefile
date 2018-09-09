build:
	docker build -t car-ping-hub-service .
run:
	docker run \
	--volume="$(PWD)/app:/usr/src/app" \
    --workdir=/usr/src/app \
    --env=GOOS=linux \
    --env=GOARCH=amd64 \
    -p 8200:8200 \
    golang bash -c "go get -d ./... && go build . && ./app"
