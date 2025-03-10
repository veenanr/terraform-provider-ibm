---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_securitycompliance"
description: |-
  Manages cd_toolchain_tool_securitycompliance.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_securitycompliance

Provides a resource for cd_toolchain_tool_securitycompliance. This allows cd_toolchain_tool_securitycompliance to be created, updated and deleted.

More information on this Continuous Delivery tool integration can be found [here](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-scc).

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_securitycompliance" "cd_toolchain_tool_securitycompliance_instance" {
  parameters {
		name = "compliance"
		evidence_namespace = "cd"
		trigger_scan = "disabled"
		scope = "my-scope"
		profile = "IBM Cloud Security Best Practices v1.0.0"
		evidence_repo_url = "https://github.example.com/<username>/compliance-evidence-<datestamp>"
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
	* `api_key` - (Optional, String) The IBM Cloud API key used to access the Security and Compliance Center API. This parameter is only relevant when the `trigger_scan` parameter is `enabled`.
	  * Constraints: The value must match regular expression `/\\S/`.
	* `evidence_namespace` - (Required, String) The kind of pipeline evidence to be displayed in Security and Compliance Center for this toolchain. The values are; `cd` which will use evidence generated by a Continuous Deployment pipeline, or `cc` which will use evidence generated by a Continuous Compliance pipeline.
	  * Constraints: Allowable values are: `cd`, `cc`.
	* `evidence_repo_url` - (Required, String) The URL to a Git repository evidence locker. The DevSecOps toolchain templates will collect and store evidence for scans and tasks in an evidence repository. This evidence URL should match the `repo_url` for a Git tool integration in this toolchain. The DevSecOps toolchain goals in the Security and Compliance Center will check the evidence repository for the pass or fail results for those goals.
	* `name` - (Required, String) The name for this tool integration, shown on the toolchain page.
	* `profile` - (Optional, String) The name of a Security and Compliance Center profile. Usually, use the predefined profile "IBM Cloud Security Best Practices v1.0.0", which contains the DevSecOps toolchain goals. Or use a user-authored customized profile that has been configured to contain those goals. When the `trigger_scan` parameter is set to `enabled`, then the Validation scan will use the controls and goals in the configured profile. If configured with a profile that does not check the DevSecOps toolchain goals, it might incorrectly indicate that the toolchain status is passed even though some of the DevSecOps scans had actually failed. This parameter is only relevant when the `trigger_scan` parameter is `enabled`.
	* `scope` - (Optional, String) The name of a Security and Compliance Center scope, which has previously been created in that service. When `trigger_scan` parameter is set to `enabled`, then the Validation scan will scan all the resources in that scope. Select a scope that contains this toolchain, so that the scan will find the evidence that has been recently updated by the DevSecOps pipeline-run. This parameter is only relevant when the `trigger_scan` parameter is `enabled`.
	* `trigger_scan` - (Optional, String) Set to `enabled` to indicate that a DevSecOps pipeline task should trigger a Security and Compliance Center run of a Validation scan. Note, each scan may incur charges. When enabled, other parameters become relevant that are needed to trigger that scan; `api_key`, `scope`, `profile`.
	  * Constraints: The default value is `disabled`. Allowable values are: `disabled`, `enabled`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind the tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cd_toolchain_tool_securitycompliance.
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

You can import the `ibm_cd_toolchain_tool_securitycompliance` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

```
<toolchain_id>/<tool_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind the tool to.
* `tool_id`: A string. ID of the tool bound to the toolchain.

# Syntax
```
$ terraform import ibm_cd_toolchain_tool_securitycompliance.cd_toolchain_tool_securitycompliance <toolchain_id>/<tool_id>
```
