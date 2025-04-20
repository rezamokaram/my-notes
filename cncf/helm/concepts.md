- [Values Hierarchy](#values-hierarchy)
  - [Helm Values Hierarchy (from lowest to highest priority):](#helm-values-hierarchy-from-lowest-to-highest-priority)

# Values Hierarchy  

## Helm Values Hierarchy (from lowest to highest priority):
1. **Chart’s values.yaml**
Default values provided by the chart maintainer.

2. **Chart dependencies’ values.yaml files**
If your chart depends on another chart (subcharts), they also have their own defaults.

3. **User-provided values.yaml file**
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