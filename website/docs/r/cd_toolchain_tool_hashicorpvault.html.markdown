---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_hashicorpvault"
description: |-
  Manages cd_toolchain_tool_hashicorpvault.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_hashicorpvault

Provides a resource for cd_toolchain_tool_hashicorpvault. This allows cd_toolchain_tool_hashicorpvault to be created, updated and deleted.

More information on this Continuous Delivery tool integration can be found [here](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-hashicorpvault).

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_hashicorpvault" "cd_toolchain_tool_hashicorpvault_instance" {
  parameters {
		name = "hcv_tool_01"
		server_url = "https://hcv.mycompany.example.com:8200"
		authentication_method = "approle"
		role_id = "<role_id>"
		secret_id = "<secret_id>"
		dashboard_url = "https://hcv.mycompany.example.com:8200/ui"
		path = "generic/project/test_project"
  }
  toolchain_id = ibm_cd_toolchain.cd_toolchain.id
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `name` - (Optional, String) Name of the tool.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
* `parameters` - (Required, List) Unique key-value pairs representing parameters to be used to create the tool. A list of parameters for each tool integration can be found in the <a href="https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-integrations">Configuring tool integrations page</a>.
Nested scheme for **parameters**:
	* `authentication_method` - (Required, String) The authentication method for your HashiCorp Vault instance.
	  * Constraints: Allowable values are: `token`, `approle`, `userpass`, `github`.
	* `dashboard_url` - (Required, String) The URL of the HashiCorp Vault server dashboard for this integration. In the graphical UI, this is the dashboard that the browser will navigate to when you click the HashiCorp Vault integration tile.
	* `default_secret` - (Optional, String) A default secret name that will be selected or used if no list of secret names are returned from your HashiCorp Vault instance.
	* `name` - (Required, String) The name used to identify this tool integration. Secret references include this name to identify the secrets store where the secrets reside. All secrets store tools integrated into a toolchain should have a unique name to allow secret resolution to function properly.
	* `password` - (Optional, String) The authentication password for your HashiCorp Vault instance when using the 'userpass' authentication method. This parameter is ignored for other authentication methods.
	* `path` - (Required, String) The mount path where your secrets are stored in your HashiCorp Vault instance.
	* `role_id` - (Optional, String) The authentication role ID for your HashiCorp Vault instance when using the 'approle' authentication method. This parameter is ignored for other authentication methods. Note, 'role_id' should be treated as a secret and should not be shared in plaintext.
	* `secret_filter` - (Optional, String) A regular expression to filter the list of secret names returned from your HashiCorp Vault instance.
	* `secret_id` - (Optional, String) The authentication secret ID for your HashiCorp Vault instance when using the 'approle' authentication method. This parameter is ignored for other authentication methods.
	* `server_url` - (Required, String) The server URL for your HashiCorp Vault instance.
	* `token` - (Optional, String) The authentication token for your HashiCorp Vault instance when using the 'github' and 'token' authentication methods. This parameter is ignored for other authentication methods.
	* `username` - (Optional, String) The authentication username for your HashiCorp Vault instance when using the 'userpass' authentication method. This parameter is ignored for other authentication methods.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind the tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cd_toolchain_tool_hashicorpvault.
* `crn` - (String) Tool CRN.
* `href` - (String) URI representing the tool.
* `referent` - (List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (String) URI representing this resource through an API.
	* `ui_href` - (String) URI representing this resource through the UI.
* `resource_group_id` - (String) Resource group where the tool can be found.
* `state` - (String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.
* `toolchain_crn` - (String) CRN of toolchain which the tool is bound to.
* `tool_id` - (String) Tool ID.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `updated_at` - (String) Latest tool update timestamp.

## Provider Configuration

The IBM Cloud provider offers a flexible means of providing credentials for authentication. The following methods are supported, in this order, and explained below:

- Static credentials
- Environment variables

To find which credentials are required for this resource, see the service table [here](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-provider-reference#required-parameters).

### Static credentials

You can provide your static credentials by adding the `ibmcloud_api_key`, `iaas_classic_username`, and `iaas_classic_api_key` arguments in the IBM Cloud provider block.

Usage:
```
provider "ibm" {
    ibmcloud_api_key = ""
    iaas_classic_username = ""
    iaas_classic_api_key = ""
}
```

### Environment variables

You can provide your credentials by exporting the `IC_API_KEY`, `IAAS_CLASSIC_USERNAME`, and `IAAS_CLASSIC_API_KEY` environment variables, representing your IBM Cloud platform API key, IBM Cloud Classic Infrastructure (SoftLayer) user name, and IBM Cloud infrastructure API key, respectively.

```
provider "ibm" {}
```

Usage:
```
export IC_API_KEY="ibmcloud_api_key"
export IAAS_CLASSIC_USERNAME="iaas_classic_username"
export IAAS_CLASSIC_API_KEY="iaas_classic_api_key"
terraform plan
```

Note:

1. Create or find your `ibmcloud_api_key` and `iaas_classic_api_key` [here](https://cloud.ibm.com/iam/apikeys).
  - Select `My IBM Cloud API Keys` option from view dropdown for `ibmcloud_api_key`
  - Select `Classic Infrastructure API Keys` option from view dropdown for `iaas_classic_api_key`
2. For iaas_classic_username
  - Go to [Users](https://cloud.ibm.com/iam/users)
  - Click on user.
  - Find user name in the `VPN password` section under `User Details` tab

For more informaton, see [here](https://registry.terraform.io/providers/IBM-Cloud/ibm/latest/docs#authentication).

## Import

You can import the `ibm_cd_toolchain_tool_hashicorpvault` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

```
<toolchain_id>/<tool_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind the tool to.
* `tool_id`: A string. ID of the tool bound to the toolchain.

# Syntax
```
$ terraform import ibm_cd_toolchain_tool_hashicorpvault.cd_toolchain_tool_hashicorpvault <toolchain_id>/<tool_id>
```
