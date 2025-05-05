
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
  - [Global variables](#global-variables)
    - [Key Points](#key-points)
    - [Global Variables Example](#global-variables-example)
  - [implicit and Explicit in values](#implicit-and-explicit-in-values)
- [Helm Starters Chart](#helm-starters-chart)
  - [Release Lifecycle](#release-lifecycle)
  - [Helm Hooks](#helm-hooks)
    - [How They Work](#how-they-work)
      - [Common Hook Annotations](#common-hook-annotations)
      - [Other Useful Hook Annotations](#other-useful-hook-annotations)
      - [Common Use Cases for Helm Hooks](#common-use-cases-for-helm-hooks)
      - [Helm Hook Example](#helm-hook-example)
  - [Helm Test](#helm-test)
    - [Helm Test Example](#helm-test-example)
  - [Helm Resource Policies Explained](#helm-resource-policies-explained)
  - [helm sign and verify charts](#helm-sign-and-verify-charts)
    - [Toolset](#toolset)
  - [helm repository host on gitlab](#helm-repository-host-on-gitlab)
  - [artifact hub](#artifact-hub)
  - [validate values by json](#validate-values-by-json)
  - [use oci registry](#use-oci-registry)

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

## Global variables  

Global variables can be overridden by values set in the `values.yaml` of subcharts. However, if you want to override the global value for a specific chart (or subchart), you can specify that override directly.

### Key Points

1. Definition: You define global variables inside the `values.yaml` file under the global key.

1. Access: To access global variables in your templates, use the global prefix (e.g., {{ .Values.global.VARIABLE }}).

1. Priority: Local values (defined in subcharts) can override global variables, but global variables can be accessed universally.

1. Inheritance: Subcharts can use global variables, which are automatically inherited from the parent chart.

### Global Variables Example

```yaml
# values.yaml (Parent chart)
global:
  image:
    tag: "v1.0.0"

# Subchart 1 (app)
app:
  image:
    tag: "{{ .Values.global.image.tag }}"  # Using global value

# Subchart 2 (frontend)
frontend:
  image:
    tag: "{{ .Values.global.image.tag }}"  # Using global value
```

Here's a concrete example of a `deployment.yaml` template in Helm that uses global variables:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deployment
spec:
  replicas: {{ .Values.replicas | default 3 }}
  template:
    spec:
      containers:
        - name: my-container
          image: "{{ .Values.global.image.repository }}:{{ .Values.global.image.tag }}"
          imagePullPolicy: "{{ .Values.global.image.pullPolicy }}"
```

## implicit and Explicit in values

- Explicit values are those that are directly defined or overridden by the user, making their source and intent clear.
- Implicit values are either default values provided by the chart developer, built-in Helm variables, or outcomes of the template logic, which are used automatically unless explicitly changed by the user.

# Helm Starters Chart

A Helm starter chart is essentially a basic, often minimal, Helm chart structure that provides a starting point for developers to create their own custom Helm charts. Instead of beginning with an empty directory, a starter chart offers pre-configured files and directories with some common elements already in place.  

Think of it like a template or a scaffolding tool specifically for Helm charts. It helps you quickly set up the fundamental structure and potentially some basic resource definitions, allowing you to focus on the specifics of your application deployment.

## Release Lifecycle

- ***Chart:*** A package containing all the necessary resource definitions (templates), metadata, and configuration files to deploy an application on Kubernetes.

- ***Release:*** A specific instance of a chart running in a Kubernetes cluster. A single chart can have multiple releases.

- ***Revision:*** An incremental version of a release, capturing the state of the deployed resources at a specific point in time.

- ***Values:*** Configuration parameters provided during installation or upgrade that are used to render the chart's templates.

- ***Hooks:*** Special lifecycle events within a release that allow chart developers to execute specific actions at certain points (e.g., pre-install, post-upgrade, pre-delete).

## Helm Hooks

Helm hooks are Kubernetes manifests (like Deployments, Jobs, Pods, etc.) that Helm manages and executes at predefined points during a release's lifecycle. These points include actions like installation, upgrade, rollback, and deletion.

### How They Work

When Helm processes a chart, it looks for special annotations within the Kubernetes manifests in the `templates/` directory. These annotations tell Helm that a particular manifest defines a hook and specifies at which point(s) in the release lifecycle the hook should be executed.

#### Common Hook Annotations

The most important annotation is `helm.sh/hook`. This annotation specifies the event that triggers the hook. Some common hook types include:

- **`pre-install`**: Executes *before* any resources in the chart are installed.
- **`post-install`**: Executes *after* all resources in the chart have been successfully installed.
- **`pre-upgrade`**: Executes *before* any resources in the chart are upgraded.
- **`post-upgrade`**: Executes *after* all resources in the chart have been successfully upgraded.
- **`pre-rollback`**: Executes *before* a rollback operation.
- **`post-rollback`**: Executes *after* a rollback operation has completed.
- **`pre-delete`**: Executes *before* any resources in the release are deleted during an uninstall operation.
- **`post-delete`**: Executes *after* all resources in the release have been deleted.
- **`test`**: Executes when you run the `helm test <release-name>` command.

You can specify multiple hook types for a single manifest using a comma-separated list (e.g., `helm.sh/hook: "pre-install,post-upgrade"`).

#### Other Useful Hook Annotations

- **`helm.sh/hook-weight`**: Allows you to define the order in which hooks of the same type are executed. Hooks with lower weights are executed first. This is useful when you have dependencies between pre- or post- actions.
- **`helm.sh/hook-delete-policy`**: Defines when the hook resource should be deleted after it has been executed. Possible values include:
- `before-hook-creation`: Delete previous hook instances before creating a new one.
- `hook-succeeded`: Delete the hook resource only if it executed successfully. (Default for most hooks)
- `hook-failed`: Delete the hook resource if it failed to execute.
- `before-hook-creation,hook-succeeded`: Delete previous instances before creating a new one, and delete the new instance if it succeeds.
- `before-hook-creation,hook-failed`: Delete previous instances before creating a new one, and delete the new instance if it fails.
- `never`: Never delete the hook resource. This can be useful for debugging or for resources that need to persist.

#### Common Use Cases for Helm Hooks

- **Database Migrations:** Running database schema migrations before a new version of your application is deployed (`pre-upgrade`).
- **Backend Initialization:** Performing setup tasks like creating default users or populating initial data after installation (`post-install`).
- **Service Availability Checks:** Verifying that dependent services are running before installing or upgrading your application (`pre-install`, `pre-upgrade`).
- **Notifications:** Sending notifications (e.g., via Slack or email) after a successful installation or upgrade (`post-install`, `post-upgrade`).
- **Cleanup Tasks:** Removing temporary resources or performing cleanup actions before or after deletion (`pre-delete`, `post-delete`).
- **Testing:** Running integration or smoke tests after a deployment to ensure it's working correctly (`test`).

#### Helm Hook Example

```yaml
# templates/migrations-job.yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: migrations-job-{{ .Release.Name }}-{{ .Release.Revision }}
  annotations:
    "helm.sh/hook": "pre-upgrade"
    "helm.sh/hook-weight": "-5"
    "helm.sh/hook-delete-policy": "hook-succeeded"
spec:
  template:
    spec:
      restartPolicy: Never
      containers:
      - name: migration
        image: your-migration-image:latest
        command: ["/app/migrate.sh"]
```

## Helm Test

Helm tests let you define and run **verification steps** within your Helm chart to ensure your deployed application is working correctly *after* installation or upgrade.

Think of it like including **mini-checks** alongside your application deployment definition. You define these checks as Kubernetes Pods (usually Jobs or Pods that run once and exit). When you run `helm test <release-name>`, Helm launches these test Pods in your cluster.

If the test Pods complete successfully (exit with a 0 status), the test is considered **passed**. If they fail (non-zero exit status), the test is **failed**.

This helps automate basic sanity checks and gives you confidence that your Helm deployment is healthy right after it's deployed.

### Helm Test Example

template:

```yaml
# mychart/templates/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-app
spec:
  selector:
    matchLabels:
      app: my-app
  replicas: 1
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: my-app-container
        image: nginx:latest
        ports:
        - containerPort: 80
```

test:

```yaml
# mychart/templates/tests/test-connection.yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-app-test-connection
  annotations:
    "helm.sh/hook": test
spec:
  containers:
  - name: test-container
    image: curlimages/curl:latest
    command: ['curl', '-I', 'http://my-app:80']
  restartPolicy: Never
```

$Explain$

1. mychart/templates/deployment.yaml: This is a standard Kubernetes Deployment that deploys a single instance of the nginx:latest image, exposing port 80.

2. mychart/templates/tests/test-connection.yaml: This defines a Kubernetes Pod that acts as our test:

- `apiVersion`: v1, kind: Pod: It's a simple Pod.
- `metadata.annotations."helm.sh/hook"`: test: This crucial annotation tells Helm that this Pod is a test hook. It will be executed when you run helm test my-release.
- `spec.containers`: Defines a single container named test-container using the `curlimages/curl:latest` image, which is a lightweight image with the curl command-line tool.
- `spec.containers.command`: The command executed within the test container is `curl -I http://my-app:80`. This sends an HTTP HEAD request to the my-app service on port 80. A successful response (HTTP status code in the 2xx or 3xx range) indicates the service is reachable.
- `spec.restartPolicy`: Never: Once the `curl` command finishes (either successfully or with an error), the Pod will not be restarted.

## Helm Resource Policies Explained

Helm Resource Policies are annotations you can add to Kubernetes manifests within your Helm chart to control how Helm handles existing resources during `helm install` and `helm upgrade` operations. They allow you to deviate from Helm's default behavior of managing all defined resources for a release.

**Key Resource Policy Annotations:**

You define Resource Policies using the `helm.sh/resource-policy` annotation in the `metadata.annotations` section of your Kubernetes manifests.

- **`keep`**: Tells Helm to **not** manage the resource. If a resource with the same name exists, Helm leaves it untouched during install, upgrade, and uninstall. Helm won't track it in the release history.

    ```yaml
    apiVersion: v1
    kind: PersistentVolumeClaim
    metadata:
      name: my-existing-pvc
      annotations:
        "helm.sh/resource-policy": keep
    spec:
      accessModes:
        - ReadWriteOnce
      resources:
        requests:
          storage: 10Gi
    ```

- **`retain`**: Tells Helm to **keep** the resource during an uninstall operation (`helm uninstall`). Helm will still manage and potentially update this resource during upgrades.

    ```yaml
    apiVersion: v1
    kind: PersistentVolumeClaim
    metadata:
      name: my-data-pvc
      annotations:
        "helm.sh/resource-policy": retain
    spec:
      accessModes:
        - ReadWriteOnce
      resources:
        requests:
          storage: 20Gi
    ```

**How to Use:**

Add the `helm.sh/resource-policy: <policy-value>` annotation to the `metadata.annotations` of the desired Kubernetes resource definition in your chart's `templates/` directory.

**Important Considerations:**

- **Scope:** Policies apply to individual resources.
- **Override Default:** They override Helm's default management.
- **`keep` Implications:** Helm won't update `kept` resources even if their definition changes in the chart.
- **`retain` Implications:** Retained resources must be managed manually after uninstall.
- **No "Adopt" Policy:** Helm doesn't have a direct policy to adopt existing, unmanaged resources.

In essence, Helm Resource Policies provide fine-grained control over how Helm interacts with specific Kubernetes resources during a release lifecycle, allowing you to handle externally managed resources or ensure the persistence of certain resources across release uninstalls.

## helm sign and verify charts

**Sign:** Chart developers cryptographically sign their Helm charts to guarantee their **authenticity** and **integrity**. This proves the chart hasn't been tampered with since it was signed and confirms the publisher.

**Verify:** Users can cryptographically verify a signed Helm chart before installation to ensure it's the original, untampered chart from a trusted source. This adds a layer of security and trust to Helm deployments.
gnupg -> use this

### Toolset

**1. For Signing:**

- **Helm CLI:** The core `helm` command-line tool itself provides the functionality to sign charts using the `helm package --sign` command. This requires a PGP keypair.
- **GnuPG (GPG):** Helm relies on GPG to handle the cryptographic signing process. You'll need GPG installed to generate and manage your PGP keys.
- **Keyring:** A keyring (usually `~/.gnupg/secring.gpg` for private keys) where your private signing key is stored. You'll need to specify the key to use during the signing process.
- **`helm-sign` (Optional):** A third-party Python tool that offers more flexibility in using existing GPG environments for signing Helm charts.

**2. For Verification:**

- **Helm CLI:** The `helm verify` command is used to check the signature of a packaged chart (`.tgz` file). Additionally, the `helm install --verify` and `helm pull --verify` flags allow verification during installation or pulling of charts.
- **GnuPG (GPG):** Helm uses GPG to verify the cryptographic signatures. You'll need GPG installed to import and manage the public keys of chart publishers you trust.
- **Public Keyring:** A keyring (usually `~/.gnupg/pubring.gpg` for public keys) containing the public keys of the chart signers you want to trust. You might need to import the public key of the chart publisher into your keyring.

## helm repository host on gitlab

## artifact hub

## validate values by json

## use oci registry
