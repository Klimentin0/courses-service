SHELL := /bin/bash

run:
	go run main.go

# - - - - - -
# Building containers

VERSION := 1.0

all: courses-service

courses-service:
	docker build \
	-f config/docker/Dockerfile \
	-t courses-service-amd64:$(VERSION) \
	--build-arg VCS_REF=$(VERSION) \
	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
	.

# - - - - - -
# Running in k8s/kind

KIND_CLUSTER := klim-starter-cluster

kind-up:
	kind create cluster \
		--image kindest/node@sha256:c48c62eac5da28cdadcf560d1d8616cfa6783b58f0d94cf63ad1bf49600cb027 \
		--name ${KIND_CLUSTER} \
		--config config/k8s/kind/kind-config.yaml

kind-down:
	kind delete cluster --name ${KIND_CLUSTER}

kind-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces
