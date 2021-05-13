# kcp upgrade cluster

Upgrades Kubernetes cluster on one or more Kyma Runtimes.

## Synopsis

Upgrade Kubernetes cluster and/or machine images on targets of Runtimes.
The upgrade is performed by Kyma Control Plane (KCP) within a new orchestration asynchronously. The ID of the orchestration is returned by the command upon success.
The targets of Runtimes are specified via the `--target` and `--target-exclude` options. At least one `--target` must be specified.
The version of Kubernetes and machine images is configured by Kyma Environment Broker (KEB).
Additional Kyma configurations to use for the upgrade are taken from Kyma Control Plane during the processing of the orchestration.

```bash
kcp upgrade cluster --target {TARGET SPEC} ... [--target-exclude {TARGET SPEC} ...] [flags]
```

## Examples

```
kcp upgrade cluster --target all --schedule maintenancewindow    Upgrade Kubernetes cluster on Runtime in their next respective maintenance window hours.
  kcp upgrade cluster --target "account=CA.*"                       Upgrade Kubernetes cluster on Runtimes of all global accounts starting with CA.
  kcp upgrade cluster --target all --target-exclude "account=CA.*"  Upgrade Kubernetes cluster on Runtimes of all global accounts not starting with CA.
  kcp upgrade cluster --target "region=europe|eu|uk"                Upgrade Kubernetes cluster on Runtimes whose region belongs to Europe.
```

## Global Options

```
      --config string                Path to the KCP CLI config file. Can also be set using the KCPCONFIG environment variable. Defaults to $HOME/.kcp/config.yaml .
      --gardener-kubeconfig string   Path to the kubeconfig file of the corresponding Gardener project which has permissions to list/get Shoots. Can also be set using the KCP_GARDENER_KUBECONFIG environment variable.
      --gardener-namespace string    Gardener Namespace (project) to use. Can also be set using the KCP_GARDENER_NAMESPACE environment variable.
  -h, --help                         Option that displays help for the CLI.
      --keb-api-url string           Kyma Environment Broker API URL to use for all commands. Can also be set using the KCP_KEB_API_URL environment variable.
      --kubeconfig-api-url string    OIDC Kubeconfig Service API URL used by the kcp kubeconfig and taskrun commands. Can also be set using the KCP_KUBECONFIG_API_URL environment variable.
      --oidc-client-id string        OIDC client ID to use for login. Can also be set using the KCP_OIDC_CLIENT_ID environment variable.
      --oidc-client-secret string    OIDC client secret to use for login. Can also be set using the KCP_OIDC_CLIENT_SECRET environment variable.
      --oidc-issuer-url string       OIDC authentication server URL to use for login. Can also be set using the KCP_OIDC_ISSUER_URL environment variable.
  -v, --verbose int                  Option that turns verbose logging to stderr. Valid values are 0 (default) - 6 (maximum verbosity).
```

## See also

* [kcp upgrade](kcp_upgrade.md)	 - Performs upgrade operations on Kyma Runtimes.

