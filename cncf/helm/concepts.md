- [Values Hierarchy](#values-hierarchy)
- [Helm Chart Structure](#helm-chart-structure)
  - [`.helmignore`](#helmignore)
    - [Example `.helmignore`:](#example-helmignore)
    - [What is “Packaging a Chart”?](#what-is-packaging-a-chart)
      - [Example:](#example)
  - [`Chart.yaml`](#chartyaml)
    - [Example](#example-1)
  - [`values.yaml`](#valuesyaml)
    - [Example](#example-2)
  - [How it works?](#how-it-works)
  - [`charts`](#charts)
    - [How to Add a Dependency (the Right Way)](#how-to-add-a-dependency-the-right-way)
    - [Summary](#summary)

# Values Hierarchy  

***From lowest to highest:***  

1. **Chart’s values.yaml**
Default values provided by the chart maintainer.

1. **Chart dependencies’ values.yaml files**
If your chart depends on another chart (subcharts), they also have their own defaults.

1. **User-provided values.yaml file**
When installing or upgrading a chart, users can pass a custom values.yaml to override the defaults.

```bash
helm install my-app ./my-chart -f custom-values.yaml
```
4. **--set on the command line**
You can override individual values inline:

```bash
helm install my-app ./my-chart --set image.tag=1.2.3
```
5. **--set-string, --set-file, etc.**
Variants of --set for more control over data types or file contents.

6. **Explicit** values.yaml for subcharts (via dependencies[].values in Chart.yaml)
This lets a parent chart define values for its subcharts.  

# Helm Chart Structure  

Helm charts are a way to define, install, and manage Kubernetes applications. A Helm chart is a collection of files that describe a related set of Kubernetes resources.  

Here’s a breakdown of the basic structure of a Helm chart:   
```scss
my-chart/
│
├── Chart.yaml          # Metadata about the chart (name, version, description, etc.)
├── values.yaml         # The default configuration values for the chart
├── charts/             # Subcharts (other charts your chart depends on)
├── templates/          # Kubernetes YAML templates (with Go templating syntax)
│   ├── deployment.yaml
│   ├── service.yaml
│   ├── ingress.yaml
│   └── _helpers.tpl    # Helper templates (functions/macros)
├── .helmignore         # Files to ignore when packaging the chart
```

## `.helmignore`

`.helmignore` is like `.gitignore` but for Helm. It tells Helm which files or directories to ignore when packaging the chart (we’ll get to packaging in a sec).  

When you run `helm package`, Helm reads your chart directory and creates a `.tgz` archive. The `.helmignore` file makes sure that unnecessary files (like .git, local dev scripts, test data, etc.) don’t end up in that archive.  

### Example `.helmignore`:
```sh
# Ignore all .md files
*.md

# Ignore Git directory
.git/

# Ignore temp files
*.swp
*.tmp

# Ignore CI/CD files
.github/
.gitlab-ci.yml
```

### What is “Packaging a Chart”?

When you **package a chart**, you create a `.tgz` (tarball) archive of your Helm chart directory. This is useful for:

- Distributing it (like uploading to a Helm chart repository)
- Versioning your chart
- Sharing it with your team or across environments

#### Example:
```bash
helm package my-chart
```  
If your my-chart/ directory looks like this:
```scss
my-chart/
├── Chart.yaml
├── values.yaml
├── templates/
└── .helmignore
```

After running `helm package my-chart`, Helm will output something like:  

```sh
Successfully packaged chart and saved it to: ./my-chart-0.1.0.tgz
```
Only files not ignored by `.helmignore` will be included in `my-chart-0.1.0.tgz`.  

***Tip:***
`.helmignore` = cleaner packages, faster installs, smaller uploads, less clutter.

Then you can install the chart directly from the .tgz file:  
```sh
helm install my-release ./my-chart-0.1.0.tgz
# or
helm upgrade my-release ./my-chart-0.1.0.tgz
```

If you're using a Helm chart repository (like GitHub Pages, ChartMuseum, or ArtifactHub), you can upload your .tgz file so others can install it.

Example: GitHub Pages as a Chart Repo

1. Move .tgz to a directory (like docs/charts/)  
2. Run:  
    ```sh
    helm repo index docs/charts --url https://yourname.github.io/repo/charts
    ```
3. Push it to GitHub
4. Add it as a repo:  
    ```sh
    helm repo add myrepo https://yourname.github.io/repo/charts
    helm repo update
    helm install my-release myrepo/my-chart
    ```

## `Chart.yaml`  

`Chart.yaml` is one of the core files in every Helm chart.  
It holds metadata about your chart—basically the "package manifest" for your Kubernetes application.  

It defines key information like:  
- The name and version of the chart  
- A description of what the chart does  
- The app version it’s deploying  
- Dependencies on other charts  
- Maintainers and license info  

### Example  
```yaml
apiVersion: v2
name: my-app
description: A Helm chart for deploying My Awesome App
type: application
version: 0.1.0
appVersion: "1.0.3"
```  

| Field         | Required | Description |
|---------------|----------|-------------|
| `apiVersion`  | ✅       | Chart API version. Use `v2` for Helm 3+. |
| `name`        | ✅       | The name of the chart (usually matches the folder name). |
| `description` | ✅       | Short summary of what the chart does. |
| `type`        | ❌       | `application` (default) or `library`. Library charts are reusable but not installable on their own. |
| `version`     | ✅       | The chart's version (used during packaging). |
| `appVersion`  | ❌       | The version of the app this chart deploys. Purely informational. |

Optional but Useful Fields:  
```yaml
maintainers:
  - name: Jane Doe
    email: jane@example.com

sources:
  - https://github.com/example/my-app

keywords:
  - web
  - frontend
  - awesome

dependencies:
  - name: redis
    version: "17.3.11"
    repository: "https://charts.bitnami.com/bitnami"
```

---

***version vs appVersion***  
*version:* The version of the chart itself (for Helm). You increment this when you update templates, configs, or chart structure.

*appVersion:* The version of the actual application you’re deploying (like a Docker image tag).

---

## `values.yaml`  

`values.yaml` is a central configuration file in a Helm chart. It holds default values that your templates refer to when rendering Kubernetes manifests.

Think of it as the "settings file" for your chart. It makes your chart flexible, customizable, and reusable without modifying the templates directly.  

**It’s where you define all the configurable parts of your chart, such as:**

- Image names and tags
- Number of replicas
- Resource limits
- Service type and ports
- Environment variables
- Custom app settings

### Example  

```yaml
replicaCount: 2

image:
  repository: nginx
  tag: stable
  pullPolicy: IfNotPresent

service:
  type: ClusterIP
  port: 80

resources:
  limits:
    cpu: 100m
    memory: 128Mi
  requests:
    cpu: 100m
    memory: 128Mi

env:
  - name: NODE_ENV
    value: production
```

## How it works?

In your `templates/` (e.g., `deployment.yaml`), you use ***Go templating*** to reference values from `values.yaml`.  
```sh
spec:
  replicas: {{ .Values.replicaCount }}
  containers:
    - image: "{{ .Values.image.repository }}:{{ .Values.image.tag }}"
      imagePullPolicy: {{ .Values.image.pullPolicy }}
```

***Tips***
- Keep your values.yaml clean and well-commented—it's your chart's API.
- Use nested structures for better organization.
- Use default in templates to avoid errors when a value isn’t set.

## `charts`  

The `charts/` folder is where you place other Helm charts that your main chart depends on. These are called subcharts.  

For example, if your app uses a database like PostgreSQL or a cache like Redis, you can include those charts here so everything installs together.  

```scss
my-app/
├── Chart.yaml
├── values.yaml
├── templates/
└── charts/
    └── redis-17.3.11.tgz  ← a dependency packaged chart
```

### How to Add a Dependency (the Right Way)
Instead of manually downloading `.tgz` files into `charts/`, you typically declare dependencies in `Chart.yaml`:  
```yaml
dependencies:
  - name: redis
    version: "17.3.11"
    repository: "https://charts.bitnami.com/bitnami"
```
Then run:  
```sh
helm dependency update
```
This will:  
- Download the Redis chart from the Bitnami repo
- Save it as a .tgz file inside the charts/ folder  

### Summary

| What It Does                | How |
|-----------------------------|-----|
| Manages subcharts (deps)    | Place `.tgz` files or use `dependencies:` in `Chart.yaml` |
| Installs with main chart    | Automatically included on `helm install` |
| Keeps things modular        | Each subchart is independent but can be configured via `values.yaml` |
| Supports version control    | Pin specific versions to avoid breaking changes |

//TODO: best practice git-ops & templates