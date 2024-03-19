MINIKUBE := minikube
KUBECTL := kubectl

.DEFAULT_GOAL := help

help:
	@echo "Usage:"
	@echo "  make start          Start Minikube cluster"
	@echo "  make stop           Stop Minikube cluster"

start:
	$(MINIKUBE) start
	eval $($(MINIKUBE) -p minikube docker-env)
	$(MINIKUBE) image load gotmpl-app

	$(KUBECTL) apply -f mongo-config.yaml
	$(KUBECTL) apply -f mongo-secret.yaml
	$(KUBECTL) apply -f mongo.yaml
	$(KUBECTL) apply -f webapp.yaml
	$(KUBECTL) get pod
	$(KUBECTL) get svc
	$(MINIKUBE) ip


stop:
	$(MINIKUBE) stop
