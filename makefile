# Check to see if we can use ash, in Alpine images, or default to BASH.
SHELL_PATH = /bin/ash
SHELL = $(if $(wildcard $(SHELL_PATH)),/bin/ash,/bin/bash)

run:
	go run app/services/courses-api/main.go | go run app/tooling/logfmt/main.go

#======================
# Define dependencies

GOLANG 			:= golang:1.23.4
ALPINE          := alpine:3.21
KIND            := kindest/node:v1.32.0
POSTGRES        := postgres:17.2
GRAFANA         := grafana/grafana:11.4.0
PROMETHEUS      := prom/prometheus:v3.0.0
TEMPO           := grafana/tempo:2.6.0
LOKI            := grafana/loki:3.3.0
PROMTAIL        := grafana/promtail:3.3.0

KIND_CLUSTER    := klim-starter-cluster
NAMESPACE       := courses-system
COURSES_APP     := courses
AUTH_APP        := auth
BASE_IMAGE_NAME := localhost/klim
VERSION         := 0.0.1
COURSES_IMAGE   := $(BASE_IMAGE_NAME)/$(COURSES_APP):$(VERSION)
METRICS_IMAGE   := $(BASE_IMAGE_NAME)/metrics:$(VERSION)
AUTH_IMAGE      := $(BASE_IMAGE_NAME)/$(AUTH_APP):$(VERSION)

#======================
# Running with k8s/kind

dev-up:
	kind create cluster \
		--image $(KIND) \
		--name $(KIND_CLUSTER) \
		--config config/k8s/dev/kind-config.yaml

		kubectl wait --timeout=120s --namespace=local-path-storage --for=condition=Available deployment/local-path-provisioner

dev-down:
	kind delete cluster --name $(KIND_CLUSTER)


#----------------------

dev-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces
