lint:
	golangci-lint run -v
test:
	go clean -testcache && go test -v -cover ./...
run:
	go run cmd/main.go
docker-build:
	docker build -t jsovalles/oauth-service .
docker-push:
	docker push docker.io/jsovalles/oauth-service:latest
kind-cluster:
	kind create cluster --name general-cluster
kube-deploy:
	kubectl apply -f deployment.yaml
kube-service:
	kubectl apply -f service.yaml
kube-pf:
	kubectl port-forward oauth-service-deployment-5697bb97c7-pmlq2 8080:8080