run:
	go run ./cmd/api

test:
	chmod +x scripts/run-tests.sh
	./scripts/run-tests.sh

cov: test
	go tool cover -html=reports/coverage.out -o reports/coverage.html

tv:
	go mod tidy && go mod vendor

image:
	docker build -t productivity:latest .

container-up:
	chmod +x scripts/run-docker.sh
	./scripts/run-docker.sh

container-down:
	docker stop productivity

db-up:
	chmod +x scripts/db-up.sh
	./scripts/db-up.sh

db-down:
	chmod +x scripts/db-down.sh
	./scripts/db-down.sh

mailhog-up:
	chmod +x scripts/mailhog-up.sh
	./scripts/mailhog-up.sh

mailhog-down:
	chmod +x scripts/mailhog-down.sh
	./scripts/mailhog-down.sh