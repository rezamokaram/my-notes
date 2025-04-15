
- [helm repo](#helm-repo)
  - [helm repo add](#helm-repo-add)
  - [helm repo list](#helm-repo-list)
  - [helm repo list](#helm-repo-list-1)
- [helm search](#helm-search)
- [helm install](#helm-install)
- [helm create](#helm-create)
- [helm upgrade](#helm-upgrade)
- [helm list](#helm-list)
  - [flags](#flags)
- [helm history](#helm-history)
- [helm status](#helm-status)
  - [flags](#flags-1)


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

## helm repo list

shows existing repositories that we added using helm repo add

# helm search
# helm install

Usage: `helm install < release-name > < repo-name-in-your-local/chart-name >`  

example: 
```sh  
helm install nginx bitnami/nginx 
```  

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

# helm list

**This command lists all of the releases for a specified namespace**  

## flags

- `--output` format        prints the output in the specified format. Allowed values: table, json, yaml (default table)  
- `--superseded`           show superseded releases  
- `--time-format` string   format time using golang time formatter. Example: --time-format "2006-01-02 15:04:05Z0700"  
- `--uninstalled`          show uninstalled releases (if 'helm uninstall --keep-history' was used)  

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

## flags  
- `--revision int`     if set, display the status of the named release with revision
- `--show-desc`        if set, display the description message of the named release
- `--show-resources`   if set, display the resources of the named release


4cmp