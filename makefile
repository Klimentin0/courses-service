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
APP     		:= courses
BASE_IMAGE_NAME := Klimentin0/courses-service
SERVICE_NAME 	:= courses-api
VERSION         := 0.0.1
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(VERSION)
METRICS_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME)-metrics:$(VERSION)

#======================
# Building containers

all: service

service:
	docker build \
		-f config/docker/dockerfile.service \
		-t $(SERVICE_IMAGE) \
		--build-arg BUILD_REF=$(VERSION) \
		--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"`

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
