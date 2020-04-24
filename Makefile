build-image:
	docker build -t eduardocfalcao/finance .

run-app:
	docker-compose -f .devops/app.yaml up -d

lint:
	golint ./...
	go fmt -n ./...