proto:
	protoc --go_out=./pkg/ --go-grpc_out=./pkg/ pkg/proto/*.proto

server:
	go run cmd/server/main.go

docker_build:
	docker build -t crm-serv:latest .

docker_run_dev:
	docker run -p "50056:50056" --env-file envs/dev.env crm-serv:latest 

docker_run_staging:
	docker run -p "50056:50056" --env-file envs/.env crm-serv:latest

build_prod_server:
	docker build -t crm-svc:latest -f Dockerfile.prod .

run_prod_server:
	docker run -p "50056:50056" --env-file .env crm-svc:latest