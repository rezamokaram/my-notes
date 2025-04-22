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
  - [`templates/`](#templates)
    - [`templates/_helpers.tpl`](#templates_helperstpl)
    - [`templates/NOTES.txt`](#templatesnotestxt)
      - [What Can You Do in `NOTES.txt`?](#what-can-you-do-in-notestxt)
    - [`templates/tests/`](#templatestests)
      - [Helm Hook for Tests](#helm-hook-for-tests)
        - [Example: templates/tests/test-connection.yaml](#example-templatesteststest-connectionyaml)
        - [How to Run Helm Tests](#how-to-run-helm-tests)
  - [`README.md`](#readmemd)
  - [`LICENSE`](#license)

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

## `templates/`

This is the most important folder — it contains the YAML templates for your Kubernetes resources.

Typical files include:

- `deployment.yaml:` Template for a Kubernetes Deployment.
- `service.yaml:` Template for a Kubernetes Service.
- `ingress.yaml:` (optional) Ingress resource.
- `_helpers.tpl:` Reusable template functions/macros.
- `NOTES.txt:` Message shown after helm install.

### `templates/_helpers.tpl`
Used to define reusable template helpers (like functions). Example:

```tpl
{{- define "mychart.fullname" -}}
{{ .Release.Name }}-{{ .Chart.Name }}
{{- end }}
```

Then used like:
```yaml
name: {{ include "mychart.fullname" . }}
```



$important-note$ -> ***The files in here are rendered using the values from values.yaml.***

Files That Start with _ (Underscore)
These files are not rendered directly into Kubernetes manifests.  
They are used to define reusable template code, such as functions, logic, or text snippets.  
Helm uses them as helpers, and you include their content in other templates using the include function.  


### `templates/NOTES.txt`  

- It's a template file (found in the templates/ directory).  
- It is not rendered into Kubernetes resources, but instead prints helpful information to the user after helm install or helm upgrade.  
- It's used to give the user instructions, like how to access the deployed app, next steps, or service URLs.  

Example output you might see:
```bash
NOTES:
1. Get the application URL by running:
  export POD_NAME=$(kubectl get pods ...)
```
  
Example `NOTES.txt`:  

```gotemplate
{{- if .Values.service.enabled }}
1. To access your service, run:
  kubectl port-forward svc/{{ include "mychart.fullname" . }} 8080:80
{{- else }}
The service is disabled in values.yaml.
{{- end }}
```

#### What Can You Do in `NOTES.txt`?

- Variables: Access `.Values`, `.Release`, .`Chart`, etc.

- Functions: `include`, `printf`, `indent`, `repeat`, etc.
    - include "mychart.fullname" . → Uses a helper defined in _helpers.tpl.  
    - printf → Builds a string (like a port-forward command).  
    - indent 4 → Adds 4 spaces to the start of each line (pretty formatting).  
    - quote → Wraps a string in " " safely.  
    - if → Conditional logic to show the admin password only if it's set.  

Example:  
```gotemplate
{{- /*
This NOTES.txt gives user instructions after deployment
*/ -}}

{{- $fullname := include "mychart.fullname" . }}
{{- $port := .Values.service.port }}
{{- $url := printf "http://%s:%d" $fullname $port }}

1. Your application "{{ $fullname }}" has been deployed.

2. To access it, you can run the following command:
   {{ printf "kubectl port-forward svc/%s 8080:%d" $fullname $port | indent 4 }}

3. Open your browser and visit:
   {{ $url | quote }}

{{- if .Values.adminPassword }}
4. Admin password is:
   {{ .Values.adminPassword | quote }}
{{- else }}
4. No admin password was set. Please check your values.yaml.
{{- end }}
```

- Conditionals:  
```gotemplate
{{ if .Values.service.enabled }}
{{ else }}
{{ end }}
```

- Loops:  
```gotemplate
{{ range .Values.users }}
User: {{ . }}
{{ end }}
```

- String formatting:
```gotemplate
{{ printf "App name: %s" .Release.Name }}
```

### `templates/tests/`  
- inside, you place Kubernetes Job manifests that act as tests for your chart.  
- These jobs run after Helm installs or upgrades your chart.  
- They are triggered using a special Helm hook.  

#### Helm Hook for Tests
To make a resource a test, add this annotation:  
```yaml
annotations:
  "helm.sh/hook": test
```  

This tells Helm:  
*“This is a test job. Run it after installing or upgrading the chart.”*

##### Example: templates/tests/test-connection.yaml  

```yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: "{{ include "mychart.fullname" . }}-test-connection"
  annotations:
    "helm.sh/hook": test
spec:
  template:
    spec:
      containers:
        - name: wget
          image: busybox
          command: ['wget']
          args: ['{{ include "mychart.fullname" . }}:{{ .Values.service.port }}']
      restartPolicy: Never
```

##### How to Run Helm Tests

Usage:  
```bash
helm test <release-name>
```

```bash
helm install myapp ./mychart
helm test myapp
```

## `README.md`
-Explains what the Helm chart does, how to use it, and any important notes.  

## `LICENSE`
- important for legal clarity and sharing the chart openly.  
