# Helm Detailed README

## Overview
Helm is a package manager for Kubernetes that simplifies deployment and management of applications using Helm charts. Helm charts are pre-configured application resources that can be deployed and managed as a single unit.

## Helm Flowchart
The following steps outline the Helm workflow:

1. **Create a Helm Chart** → Develop a Helm chart with templates and values.
2. **Package the Chart** → Helm packages the chart into a `.tgz` archive.
3. **Push to Chart Repository** → Store the chart in a Helm repository.
4. **Install the Chart** → Deploy the application using `helm install`.
5. **Upgrade or Rollback** → Modify or rollback the release as needed.
6. **Uninstall the Chart** → Remove the deployed application.

## Helm Commands with Examples

### 1. Helm Installation
```sh
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

### 2. Create a New Helm Chart
```sh
helm create mychart
```
This creates a `mychart` directory with a predefined structure.

### 3. Package the Chart
```sh
helm package mychart
```
It creates a `.tgz` file that can be stored in a Helm repository.

### 4. Add a Chart Repository
```sh
helm repo add myrepo https://example.com/helm-charts
```

### 5. Update the Repository
```sh
helm repo update
```

### 6. Install a Chart
```sh
helm install myrelease mychart --namespace mynamespace
```

### 7. List Installed Releases
```sh
helm list --namespace mynamespace
```

### 8. Upgrade a Release
```sh
helm upgrade myrelease mychart --set image.tag=2.0
```

### 9. Rollback a Release
```sh
helm rollback myrelease 1
```

### 10. Uninstall a Release
```sh
helm uninstall myrelease
```

### 11. Get Information About a Release
```sh
helm get all myrelease
```

### 12. Get Values of a Release
```sh
helm get values myrelease
```

### 13. Get Manifest of a Release
```sh
helm get manifest myrelease
```

### 14. Lint a Chart
```sh
helm lint mychart
```

### 15. Render Templates Without Installing
```sh
helm template mychart
```

### 16. Delete a Chart From Repository
```sh
helm repo remove myrepo
```

### 17. Search for Charts
```sh
helm search repo mychart
```

### 18. Show Chart Information
```sh
helm show chart mychart
```

### 19. Show Chart Values
```sh
helm show values mychart
```

### 20. Test a Release
```sh
helm test myrelease
```

## Custom Helm Chart
To create a custom Helm chart:

### 1. Define Templates
Modify `templates/deployment.yaml`:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: {{ .Release.Name }}
spec:
  replicas: {{ .Values.replicas }}
  template:
    spec:
      containers:
        - name: mycontainer
          image: {{ .Values.image.repository }}:{{ .Values.image.tag }}
```

### 2. Define Values
Edit `values.yaml`:
```yaml
replicas: 3
image:
  repository: nginx
  tag: latest
```

### 3. Install the Custom Chart
```sh
helm install customrelease ./mychart
```

This README provides a complete guide to Helm workflow, essential commands, and custom chart creation.


