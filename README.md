# K8S Image Admission controler
## Badges

[![Build Status](https://github.com/thibaultserti/k8s-image-admission-controller/actions/workflows/release.yaml/badge.svg)](https://github.com/thibaultserti/k8s-image-admission-controller/actions/workflows/release.yaml)
[![License](https://img.shields.io/github/license/thibaultserti/k8s-image-admission-controller)](/LICENSE)
[![Release](https://img.shields.io/github/release/thibaultserti/k8s-image-admission-controller.svg)](https://github.com/thibaultserti/k8s-image-admission-controller/releases/latest)
[![GitHub Releases Stats of k8s-image-admission-controller](https://img.shields.io/github/downloads/thibaultserti/k8s-image-admission-controller/total.svg?logo=github)](https://somsubhra.github.io/github-release-stats/?username=thibaultserti&repository=k8s-image-admission-controller)

[![Maintainability](https://api.codeclimate.com/v1/badges/4133d7da3d73fa0c0884/maintainability)](https://codeclimate.com/github/thibaultserti/k8s-image-admission-controller/maintainability)
[![codecov](https://codecov.io/gh/thibaultserti/k8s-image-admission-controller/branch/main/graph/badge.svg?token=5BO47LR632)](https://codecov.io/gh/thibaultserti/k8s-image-admission-controller)
[![Go Report Card](https://goreportcard.com/badge/github.com/thibaultserti/test-saas-ci)](https://goreportcard.com/report/github.com/thibaultserti/k8s-image-admission-controller)

## Test

### Local


### In a K8S cluster

Create k3d cluster

```bash
k3d cluster create k8s-image-admission-test
```


Install cert-manager
```
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.11.0/cert-manager.yaml
```

Install resources
```bash
kubectl apply -f k8s/00_namespace.yaml
kubectl apply -f k8s/10_ca_certificate.yaml
kubectl apply -f k8s/10_certificate.yaml
kubectl apply -f k8s/20_deployment.yaml
kubectl apply -f k8s/20_service.yaml
kubectl apply -f k8s/30_validatingwebhookconfiguration.yaml
```

Test that pod is denied
```bash
kubectl apply -f k8s/90_pod-test.yaml
```
