
# Helm Concepts

- [Helm Concepts](#helm-concepts)
- [Values Hierarchy](#values-hierarchy)
- [Helm Chart Structure](#helm-chart-structure)
  - [`.helmignore`](#helmignore)
    - [Example `.helmignore`](#example-helmignore)
    - [What is “Packaging a Chart”?](#what-is-packaging-a-chart)
      - [Example](#example)
  - [`Chart.yaml`](#chartyaml)
    - [Chart.yaml Example](#chartyaml-example)
  - [`values.yaml`](#valuesyaml)
    - [values.yaml Example](#valuesyaml-example)
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
- [Built-In Objects](#built-in-objects)
  - [Root Object](#root-object)
- [Template Development](#template-development)
  - [action](#action)
  - [invalid actions](#invalid-actions)
    - [Examples of Invalid Actions](#examples-of-invalid-actions)
  - [Quote Function](#quote-function)
    - [Quote Function Example](#quote-function-example)
  - [Pipeline](#pipeline)
  - [Upper](#upper)
  - [Lower](#lower)
  - [Squote](#squote)
  - [Default](#default)
  - [White Spaces](#white-spaces)
  - [Indent](#indent)
  - [Nindent](#nindent)
  - [ToYaml](#toyaml)
  - [if / else statement (control flow)](#if--else-statement-control-flow)
    - [if / else statement Example](#if--else-statement-example)
  - [`with`](#with)
    - [`with` Example](#with-example)
  - [`range`](#range)
    - [`range` Example](#range-example)
  - [Variables](#variables)
    - [Example Of Usage](#example-of-usage)
  - [Define (named template)](#define-named-template)
    - [Define Example](#define-example)
  - [Template (named template)](#template-named-template)
    - [Template Example](#template-example)
  - [Include (named template) | pipeline](#include-named-template--pipeline)
    - [Include Example](#include-example)
  - [Dependencies](#dependencies)
    - [Types of Dependencies](#types-of-dependencies)
    - [Dependencies Example](#dependencies-example)
    - [Local Dependencies Example](#local-dependencies-example)
    - [Chart.lock](#chartlock)
      - [Purpose of Chart.lock](#purpose-of-chartlock)
    - [Dependencies Common Pitfalls](#dependencies-common-pitfalls)
    - [Dependencies Key Takeaways](#dependencies-key-takeaways)
  - [helm dependencies alias](#helm-dependencies-alias)
    - [alias features](#alias-features)
    - [alias features Example](#alias-features-example)
    - [How to Reference Aliased Dependencies](#how-to-reference-aliased-dependencies)
  - [helm dependencies condition](#helm-dependencies-condition)
    - [helm dependencies condition Example](#helm-dependencies-condition-example)
  - [helm chart tags](#helm-chart-tags)
    - [Purpose Of Tags In Helm](#purpose-of-tags-in-helm)
    - [Tags Usage Example](#tags-usage-example)
    - [Tags Key Rules](#tags-key-rules)
    - [how to override subchart values from parent chart](#how-to-override-subchart-values-from-parent-chart)

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

1. **--set on the command line**
You can override individual values inline:

```bash
helm install my-app ./my-chart --set image.tag=1.2.3
```

1. **--set-string, --set-file, etc.**
  Variants of --set for more control over data types or file contents.

1. **Explicit** values.yaml for subcharts (via dependencies[].values in Chart.yaml)
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

### Example `.helmignore`

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

#### Example

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

### Chart.yaml Example  

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

### values.yaml Example

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

# Built-In Objects  

In Helm, built-in objects are special variables that are available by default in your Helm templates. These objects provide access to useful information about the release, chart, values, files, and more. They are a core part of how Helm templates dynamically render Kubernetes manifests.

| Built-in Object   | Description                                                      |
|-------------------|------------------------------------------------------------------|
| `.Release`        | Info about the current release (name, namespace, revision, etc.) |
| `.Chart`          | Info about the chart (name, version, description, etc.)          |
| `.Values`         | Values passed into the chart (`values.yaml`, `--set`, etc.)      |
| `.Files`          | Access non-template files in the chart (e.g., config files)      |
| `.Capabilities`   | Info about the Kubernetes version and API capabilities           |
| `.Template`       | Info about the template that is currently being rendered         |
| `.Chart.Name`     | Shorthand to access the chart’s name                             |
| `.Release.Name`   | The name of the release being installed/upgraded                 |

## Root Object  

In Helm charts, the root object refers to the top-level context available within templates, represented by the dot (.). This object encompasses several built-in objects that provide essential information and functionalities for rendering Kubernetes manifests.​

***The root object (.) is a dictionary containing various built-in objects, including:​***

- `.Release`: Information about the release, such as its name, namespace, and revision.

- `.Values`: User-defined values supplied in the values.yaml file or via the command line.

- `.Chart`: Metadata about the chart, including its name and version.

- `.Capabilities`: Information about the capabilities of the Kubernetes cluster.

- `.Files`: Access to non-template files within the chart.​

# Template Development

## action

actions are instructions wrapped in `{{ ... }}` — they tell Helm what to do when rendering templates.  

```yaml
metadata:
  name: {{ .Release.Name }}
```

- `{{ ... }}`: The action block  
- `.Release.Name`: A built-in object — this will be replaced with the release name  

Other common actions:

- `{{ .Values.someKey }}` – Access values from values.yaml

- `{{ if ... }}`, `{{ range ... }}`, `{{ with ... }}` – Control structures

- `{{ include "template.name" . }}` – Include another template

## invalid actions  

An invalid action is anything that breaks Go templating rules or Helm conventions. These cause template rendering to fail.  

### Examples of Invalid Actions

| Example | Why it’s Invalid |
|---------|------------------|
| `{{ .Values }}` alone | Outputs an entire map (not useful, often breaks YAML) |
| `{{ .Values.someKey` | Missing closing `}}` |
| `{{ .Release.Name }}` inside a string but unquoted: `name: {{ .Release.Name }}-svc` | May break YAML parsing if the result contains special characters |
| `{{ include "mytpl" }}` | Missing `.` context parameter |
| `{{ if }}` | Missing condition |

***$Tip$: Use helm template --debug to see exactly what failed.***

## Quote Function  

```gotemplate
{{ quote VALUE }}
```

### Quote Function Example

```yaml
name: {{ .Values.appName }}
nameQuote: {{ quote .Values.appName }}
```

output:

```yaml
app: my-app
nameQuote: "my-app"
```

## Pipeline

To pass something as input of next function, usage:

```yaml
name: {{ .Values.appName | quote }}  
```  

## Upper

To make all the characters uppercase, usage:

```yaml
name: {{ .Values.appName | quote | upper }}  
```  

## Lower

To make all the characters lowercase, usage:

```yaml
name: {{ .Values.appName | quote | lower }}  
```  

## Squote

To put inside two single quotations(''), usage:

```yaml
name: {{ .Values.appName | squote }}  
```  

## Default

To set default value for any value if it is no provided, usage:  

```yaml
replicas: {{ default 2 .Values.ReplicaCount }}
name: {{ default "my-default-value" .Values.Name | lower }}
```  

***$Tip$: in numeric values 0 considered as empty(not provided), so the default value will assigned.***  

## White Spaces

To remove a white space we can use hyphen(`-`), usage:  

```yaml
leadingWhiteSpaces: "    {{- .Chart.Name }}    sample"
trailingWhiteSpaces: "    {{ .Chart.Name -}}    sample"
leadTrailWhiteSpaces: "    {{- .Chart.Name -}}    sample"
```  

## Indent  

To put spaces before value, usage:  

```yaml
indentName: "  {{- .Chart.Name | indent 4 -}}  "
```

## Nindent  

To put spaces before value in new line, usage:  

```yaml
indentName: "  {{- .Chart.Name | nindent 4 -}}  "
```

## ToYaml  

this is a type conversion function.  
there is a lots of other type conversion funcs in [helm doc](https://helm.sh/docs/chart_template_guide/function_list/) for functions.

Convert list, slice, array, dict, or object to indented yaml, can be used to copy chunks of yaml from any source.  

```yaml
containers:
- name: nginx
  image: sample.com/nginx
  ports:
  - containerPort: 80
  resources:
  {{- toYaml .Value.resource | nindent 10}}
```

## if / else statement (control flow)  

### if / else statement Example  

```yaml
spec:
{{- if eq .Values.myapp.env "prod" }}
  replicas: 4
{{- else if eq .Values.myapp.env "qa" }}
  replicas: 2
{{- else }}
  replicas: 1
{{- end }}
```  

- `eq`: return true if two next values are same and return false otherwise.
- we have other boolean operators: `or`, `and`, `not` and ...
- In Helm templates, parentheses () function like regular parentheses in programming—they enforce the priority of an expression's evaluation.  

## `with`  

Limits (change) the scope of a block to a specific object, avoiding repetitive references.  
  
When using control structures like `with` or `range`, the context (.) changes to the current item within the block. To access the original root context in such cases, Helm provides the $ variable, which always points to the root context. This is particularly useful when you need to reference top-level objects from within nested scopes.  

### `with` Example

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: {{ .Release.Name }}-configmap
data:
  myvalue: "Hello World"
  {{- $relname := .Release.Name -}}
  {{- with .Values.favorite }}
  drink: {{ .drink | default "tea" | quote }}
  food: {{ .food | upper | quote }}
  release: {{ $relname }}
  {{- end }}
```

- `{{ .Release.Name }}` accesses the release name from the root context.

- `{{- $relname := .Release.Name -}}` assigns the release name to a variable `$relname` before entering the `with` block.

- Within the with block, `.drink` and `.food` refer to properties under .Values.favorite.

- `{{ $relname }}` accesses the release name using the variable defined earlier, ensuring access to the root context within the nested block.​  

- `$` refers to the root context of the template, allowing access to global variables even inside scoped blocks (e.g., `with`/`range`).

## `range`  

Iterates over lists or maps, similar to a for loop.  
we can use range for lists.

### `range` Example  

If your values.yaml defines a map:  

```yaml
env:
  LOG_LEVEL: debug
  TIMEOUT: 30
```  
  
You can iterate over the key-value pairs:  

```yaml
{{- range $key, $value := .Values.env }}
- name: {{ $key }}
  value: "{{ $value }}"
{{- end }}
```

This will generate:

```yaml
- name: LOG_LEVEL
  value: "debug"
- name: TIMEOUT
  value: "30"
```  

## Variables  

### Example Of Usage  

```yaml
{{- $chartname := .Chart.Name }}

# usage 
name: {{ $chartname }}
```  

## Define (named template)

In Go's `text/template` or `html/template` package, a named template is a way to define a reusable block of template logic with a specific name that can be invoked later — very similar to what you see in Helm templates (since Helm uses Go templates under the hood).  

### Define Example  

```yaml
{{- define "mychart.labels" }}
app: {{ .Chart.Name }}
chart: {{ .Chart.Name }}-{{ .Chart.Version | replace "+" "_" }}
release: {{ .Release.Name }}
heritage: {{ .Release.Service }}
{{- end }}
```

## Template (named template)

### Template Example

```yaml
# This writes output directly
metadata:
  name: {{ template "mychart.fullname" . }}
```

## Include (named template) | pipeline

### Include Example

```yaml
metadata:
  labels:
    {{- include "mychart.labels" . | nindent 4 }}
```

| Term                  | Description |
|-----------------------|-------------|
| `define`              | Declares a named template block (used with `include` or `template`) |
| `include`             | Calls a named template and returns its output as a string (can be piped) |
| `template`            | Renders a named template inline (writes output directly) |

***$Tip$*** named template can used inside another named template

## Dependencies  

Dependencies are external Helm charts that your parent chart relies on (e.g., databases, Redis, or shared libraries). They are defined in:

- `Chart.yaml`: Under the `dependencies` field.
- `charts/` directory: Where downloaded/subcharts are stored.

### Types of Dependencies

| Type | Description |
|------|-------------|
| Subcharts | Charts stored locally in the charts/ directory. |
| External | Charts fetched from repositories (e.g., Bitnami, Artifact Hub). |

### Dependencies Example  

```yaml
# Parent Chart.yaml
dependencies:
  - name: postgresql
    version: "12.0.0"
    repository: "https://charts.bitnami.com/bitnami"
```

then:

```sh
helm install myapp ./mychart
# This generates a Chart.lock
helm install myapp ./mychart
```

### Local Dependencies Example  

Directory Structure:

```bash
my-parent-chart/
├── Chart.yaml          # Parent chart (declares dependency on local subchart)
├── values.yaml         # Parent values (can override subchart values)
├── charts/             # Local subcharts (manually placed here)
│   └── my-subchart/    # Local subchart
│       ├── Chart.yaml
│       ├── values.yaml
│       └── templates/
│           └── deployment.yaml
└── templates/          # Parent chart templates
    └── NOTES.txt
```

then the chart will be like:

```yaml
apiVersion: v2
name: my-parent-chart
description: A chart with a local subchart dependency
version: 0.1.0
dependencies:
  - name: my-subchart
    version: "0.1.0"               # Must match the subchart's version
    repository: "file://./charts/my-subchart"  # Local path
    condition: my-subchart.enabled  # Optional: Enable/disable via parent values
```

### Chart.lock

The Chart.lock file is an auto-generated dependency manifest that locks the exact versions of subcharts or external dependencies used in a Helm chart. It ensures reproducible deployments by freezing dependency versions.

#### Purpose of Chart.lock

- **Version Pinning**: Records the precise versions of dependencies resolved during helm dependency update.

- **Reproducibility**: Ensures the same chart versions are used across environments (dev, staging, prod).

- **Consistency**: Prevents unexpected updates when sharing charts with others.  

### Dependencies Common Pitfalls

- **Version Conflicts**: Ensure dependency versions are compatible.

- **Broken Repos**: Verify repository URLs with helm repo list.

- **Orphaned Files**: Manually delete charts/ if dependencies change.  

### Dependencies Key Takeaways

- `Reusability`: Share common services (DBs, caches) across charts.

- `Control`: Enable/disable dependencies via values.yaml.

- `Isolation`: Subcharts can be tested independently.

## helm dependencies alias

In Helm, when managing complex charts, you often have dependencies—other charts that your main chart relies on. Sometimes, you might want to include multiple instances of the same dependency or give dependencies alternative names within your chart. This is where dependencies aliasing comes into play.

Alias in Helm dependencies allows you to specify an alternative name for a dependency chart within your parent chart's requirements.yaml (Helm v2) or Chart.yaml (Helm v3). This is particularly useful when:

- You want to include multiple instances of the same dependency (e.g., deploying two different Redis instances).
- You want to differentiate between multiple dependencies of the same chart.
- You want to avoid naming conflicts.

### alias features

- Alias is used to assign a custom name to a dependency chart.
- It allows multiple instances of the same dependency or better organization.
- You specify alias alongside the dependency in Chart.yaml.
- You reference the alias in your templates and values.

### alias features Example

```yaml
dependencies:
  - name: redis
    version: 14.8.0
    repository: "https://charts.bitnami.com/bitnami"
    alias: redis-primary
  - name: redis
    version: 14.8.0
    repository: "https://charts.bitnami.com/bitnami"
    alias: redis-secondary
```

### How to Reference Aliased Dependencies

Once you've set aliases, you can refer to them in your Helm templates using the alias name:

```yaml
# Example in a deployment.yaml template
containers:
  - name: redis-primary
    image: {{ .Values.redis-primary.image }}
  - name: redis-secondary
    image: {{ .Values.redis-secondary.image }}
```

And in your values.yaml, you might specify:

```yaml
redis-primary:
  image: redis:6.0
redis-secondary:
  image: redis:6.2
```

## helm dependencies condition

When defining dependencies in your Helm chart, you specify them in the Chart.yaml under the dependencies section. Each dependency can have a condition field which controls whether that subchart is enabled or disabled during deployment.

- The condition is a string that points to a value in your Helm chart's values.yaml.
- If the value at this path evaluates to true, the dependency is enabled (installed).
- If it evaluates to false, the dependency is skipped (not installed).

### helm dependencies condition Example

Suppose you have a dependency like this in your Chart.yaml:

```yaml
dependencies:
  - name: mysubchart
    version: 1.2.3
    repository: "https://example.com/charts"
    condition: mysubchart.enabled
```

In your values.yaml, you'd then have:

```yaml
mysubchart:
  enabled: true
```

## helm chart tags

tags are used primarily within Helm Chart's values.yaml or templates to control the inclusion or configuration of resources based on certain conditions. They help you manage different deployment scenarios, environments, or feature toggles.

### Purpose Of Tags In Helm

- **Conditional resource inclusion**: Tags allow you to specify whether certain parts of the chart should be deployed.
- **Feature toggles**: Enable or disable features based on tags.
- **Customization**: Customize chart behavior for different environments or use cases.

### Tags Usage Example

```yaml
dependencies:
  - name: postgresql
    version: "12.1.0"
    repository: "https://charts.bitnami.com/bitnami"
    condition: postgresql.enabled
    tags:
      - database
      - backend
```

then, Enable/disable tags via --set when installing:

```bash
helm install myapp . --set tags.database=true --set tags.frontend=false
```

### Tags Key Rules

1. Default Behavior: If no tag is specified, Helm installs all dependencies (unless disabled by condition).

1. Tag Precedence:

    - If any tag in a dependency is true, the dependency is enabled.
    - Example: redis above is enabled if either database or cache is true.

1. Combining with condition:

    - Tags are evaluated before condition.
    - If a tag disables a dependency, the condition is ignored.

### how to override subchart values from parent chart

1. Parent `Chart.yaml`:

    ```yaml
    dependencies:
      - name: redis
        version: "16.0.0"
        repository: "https://charts.bitnami.com/bitnami"
    ```

1. Parent `values.yaml`:

    ```yaml
    redis:
      architecture: "standalone"  # Overrides default (replication)
      auth:
        password: "mypassword"
    ```

1. Install with CLI overrides:

    ```yaml
    helm install my-app . --set redis.auth.password="newpass"
    ```
