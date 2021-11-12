build-image:
	docker build -t xhachimanx/finance .

run-app:
	docker-compose -f .devops/app.yml up -d

lint:
	golint ./...
	go fmt -n ./...

unit-test:
	go test ./...