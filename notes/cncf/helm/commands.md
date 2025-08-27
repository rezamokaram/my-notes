
# Helm Commands

- [Helm Commands](#helm-commands)
- [global flags](#global-flags)
- [helm repo](#helm-repo)
  - [helm repo add](#helm-repo-add)
  - [helm repo list](#helm-repo-list)
- [helm search](#helm-search)
- [helm install usage](#helm-install-usage)
  - [helm install example](#helm-install-example)
  - [flags](#flags)
- [helm uninstall](#helm-uninstall)
  - [helm uninstall flags](#helm-uninstall-flags)
- [helm create](#helm-create)
- [helm upgrade](#helm-upgrade)
  - [helm upgrade flags](#helm-upgrade-flags)
- [helm rollback](#helm-rollback)
  - [helm rollback flags](#helm-rollback-flags)
- [helm list](#helm-list)
  - [helm list flags](#helm-list-flags)
- [helm history](#helm-history)
- [helm status](#helm-status)
  - [helm status flags](#helm-status-flags)
- [helm get](#helm-get)
  - [helm get usage](#helm-get-usage)
  - [helm get commands](#helm-get-commands)
- [helm template](#helm-template)
  - [helm template usage](#helm-template-usage)
- [helm package](#helm-package)
  - [helm package usage](#helm-package-usage)
  - [helm package example](#helm-package-example)
  - [Package Correct Usage](#package-correct-usage)
- [helm plugin](#helm-plugin)

# global flags

These flags can be used with any `helm` command to control Helm's behavior globally.

| Flag Usage | Description |
|------------|-------------|
| `--debug` | Enables verbose output for debugging purposes. Useful when troubleshooting errors or understanding what Helm is doing under the hood. |
| `--kube-apiserver <URL>` | Specify the address and port of the Kubernetes API server (e.g., `https://127.0.0.1:6443`). |
| `--kube-as-group <group1,group2,...>` | Specify a list of groups to impersonate when performing operations (can be used multiple times or comma-separated). |
| `--kube-as-user <username>` | Impersonate the given user when interacting with the Kubernetes API server. |
| `--kube-ca-file <path>` | Path to a certificate file for the certificate authority (CA) used to verify the Kubernetes API server’s certificate. |
| `--kube-context <name>` | Use the specified context from the kubeconfig file (helpful if your kubeconfig has multiple clusters). |
| `--kube-token <token>` | Bearer token for authentication to the Kubernetes API server (alternative to kubeconfig credentials). |
| `--kubeconfig <path>` | Path to the kubeconfig file (default is usually `~/.kube/config`). |
| `--namespace <namespace>` or `-n <namespace>` | Specify the namespace scope for this operation (defaults to `default` if not set). |
| `--registry-config <path>` | Path to the registry config file used for OCI registries (default is `~/.config/helm/registry/config.json`). |
| `--repository-cache <path>` | Path to the directory where Helm stores cached repository index files (default is `~/.cache/helm/repository`). |
| `--repository-config <path>` | Path to the file containing Helm repository names and URLs (default is `~/.config/helm/repositories.yaml`). |

> 💡 You can combine multiple global flags in a single command, and they always come **before** the Helm subcommand (e.g., `install`, `upgrade`, `list`).

# helm repo

## helm repo add  
  
```sh
helm repo add bitnami https://charts.bitnami.com/bitnami 
```

## helm repo list  
  
To show the added repositories:  

```sh
helm repo list
```

# helm search

Search provides the ability to search for Helm charts in the various places  
they can be stored including the Artifact Hub and repositories you have added.  
Use search subcommands to search different locations for charts.  

Usage:

```sh
  helm search [command]
```

**Available Commands:**  
  `hub` search for charts in the Artifact Hub or your own hub instance  
  `repo` search repositories for a keyword in charts  

# helm install usage  

```sh
helm install < release-name > < repo-name-in-your-local/chart-name >
```  

## helm install example  

```sh  
helm install nginx bitnami/nginx 
```  

## flags  

- `--create-namespace` create the release namespace if not present
- `--dry-run string[="client"]` simulate an install. If --dry-run is set with no option being specified or as '--dry-run=client', it will not attempt cluster connections. Setting '--dry-run=server' allows attempting cluster connections.  
- `--set stringArray` set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)  

  ***note:*** in case that we need to removing a default value:

  ```bash
  helm install my-release my-chart --set someKey="null"
  ```

# helm uninstall

**This command takes a release name and uninstalls the release.**

Usage:

```sh
helm uninstall RELEASE_NAME [...] [flags]
```

## helm uninstall flags

- `--keep-history` remove all associated resources and mark the release as deleted, but retain the release history

# helm create

```sh
helm create my-app  # Generates a chart template  
helm install my-app ./my-app -f values.yaml  # Deploy with custom configs
```

# helm upgrade

This command upgrades a release to a new version of a chart.

The upgrade arguments must be a release and chart. The chart
argument can be either: a chart reference('example/mariadb'), a path to a chart directory,
a packaged chart, or a fully qualified URL. For chart references, the latest
version will be specified unless the '--version' flag is set.  

To upgrade by image tag (related to chart):

```sh
helm upgrade nginx bitnami/nginx --set "image.tag=1.26.0"
```  

To upgrade by chart version (related to chart):

```sh
helm upgrade nginx bitnami/nginx --version 16.0.0  
```

To upgrade by replica count (related to chart):

```sh
helm upgrade nginx bitnami/nginx --set "replicaCount=1"  
```  

## helm upgrade flags  

- `-f, --values strings` specify values in a YAML file or a URL (can specify multiple)  

# helm rollback

**The first argument of the rollback command is the name of a release, and the second is a revision (version) number. If this argument is omitted or set to 0, it will roll back to the previous release.**

Usage:

```sh
helm rollback <RELEASE> [REVISION] [flags]
```

## helm rollback flags

- `--cleanup-on-fail` allow deletion of new resources created in this rollback when rollback fails
- `--dry-run` simulate a rollback
- `--force` force resource update through delete/recreate if needed

# helm list

**This command lists all of the releases for a specified namespace.**  

## helm list flags

- `--output format`        prints the output in the specified format. Allowed values: table, json, yaml (default table)  
- `--superseded` show superseded releases  
- `--time-format string`   format time using golang time formatter. Example: --time-format "2006-01-02 15:04:05Z0700"  
- `--uninstalled` show uninstalled releases (if 'helm uninstall --keep-history' was used)  

# helm history  

History prints historical revisions for a given release.  
  
Usage:  
  `helm history RELEASE_NAME [flags]`  

# helm status

**This command shows the status of a named release.**  
The status consists of:  

- last deployment time
- k8s namespace in which the release lives
- state of the release (can be: unknown, deployed, uninstalled, superseded, failed, uninstalling, pending-install, pending-upgrade or pending-rollback)
- revision of the release
- description of the release (can be completion message or error message, need to enable --show-desc)
- list of resources that this release consists of (need to enable --show-resources)
- details on last test suite run, if applicable
- additional notes provided by the chart

Usage:  
  `helm status RELEASE_NAME [flags]`  

## helm status flags  

- `--revision int`     if set, display the status of the named release with revision
- `--show-desc`        if set, display the description message of the named release
- `--show-resources`   if set, display the resources of the named release

# helm get

This command consists of multiple subcommands which can be used to
get extended information about the release, including:

- The values used to generate the release
- The generated manifest file
- The notes provided by the chart of the release
- The hooks associated with the release
- The metadata of the release

## helm get usage

```sh  
  helm get [command]
```

## helm get commands  

- `all`         download all information for a named release
- `hooks`       download all hooks for a named release
- `manifest`    download the manifest for a named release
- `metadata`    This command fetches metadata for a given release
- `notes`       download the notes for a named release
- `values`      download the values file for a named release

# helm template  

The helm template command renders a Helm chart into raw Kubernetes manifests (YAML files) using your values, but without installing it.  

## helm template usage  

```bash  
helm template [RELEASE_NAME] [CHART] [flags]
```  

# helm package

Helm packaging is the process of bundling all the files of a Helm chart into a single `.tgz` archive so it can be shared, versioned, and installed easily. Think of it like creating a ZIP file of your application’s Kubernetes configuration, with metadata and templating.  

## helm package usage  

```sh
  helm package [CHART_PATH] [...] [flags]
```

## helm package example

Here’s what a typical chart directory looks like:

```scss
mychart/
├── Chart.yaml       # Chart metadata (name, version, etc.)
├── values.yaml      # Default configuration values
├── templates/       # Kubernetes manifests with Go templating
├── charts/          # Subcharts (optional)
└── README.md        # Optional chart documentation
```

Use the helm package command:

```sh
helm package mychart/
# mychart-1.2.3.tgz <- output
```

Helm does not automatically assign the version when you run `helm package mychart/`  
Instead, it uses the version you manually set in the Chart.yaml file inside the chart directory.  

then we can upload it using:  

```sh
helm repo index . --url https://example.com/charts
```

then, Upgrade the Release Using the New Package:

```bash
helm upgrade myrelease ./mychart-1.2.4.tgz

#If you also changed configuration (or want to override something):  
helm upgrade myrelease ./mychart-1.2.4.tgz -f my-values.yaml

# Or override inline:  
helm upgrade myrelease ./mychart-1.2.4.tgz --set image.tag=v2
```

## Package Correct Usage

| Task               | Command                                                                 |
|--------------------|-------------------------------------------------------------------------|
| First install      | `helm install myrelease ./mychart-1.2.3.tgz`                            |
| Update release     | `helm upgrade myrelease ./mychart-1.2.4.tgz`                            |
| Delete + reinstall | `helm uninstall myrelease && helm install myrelease ./mychart-1.2.4.tgz` |

# helm plugin
