---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_bitbucketgit"
description: |-
  Manages cd_toolchain_tool_bitbucketgit.
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_bitbucketgit

Provides a resource for cd_toolchain_tool_bitbucketgit. This allows cd_toolchain_tool_bitbucketgit to be created, updated and deleted.

More information on this Continuous Delivery tool integration can be found [here](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-bitbucket).

## Example Usage

```hcl
resource "ibm_cd_toolchain_tool_bitbucketgit" "cd_toolchain_tool_bitbucketgit_instance" {
  initialization {
		git_id = "bitbucketgit"
		owner_id = "<bitbucket-user-id>"
		repo_name = "myrepo"
		source_repo_url = "https://bitbucket.org/source-repo-owner/source-repo"
		type = "clone"
		private_repo = true
  }
  parameters {
		enable_traceability = false
		integration_owner = "my-userid"
		toolchain_issues_enabled = true
  }
  toolchain_id = ibm_cd_toolchain.cd_toolchain.id
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `initialization` - (Required, List) 
Nested scheme for **initialization**:
	* `git_id` - (Optional, Forces new resource, String) Set this value to 'bitbucketgit' for bitbucket.org, or to the GUID of a custom Bitbucket server.
	* `owner_id` - (Optional, Forces new resource, String) The Bitbucket user or group that owns the repository.  This parameter is required when creating a new repository, cloning, or forking a repository.  The value will be computed when linking to an existing repository.
	* `private_repo` - (Optional, Forces new resource, Boolean) Set this value to 'true' to make the repository private when creating a new repository or when cloning or forking a repository.  This parameter is not used when linking to an existing repository.
	  * Constraints: The default value is `false`.
	* `repo_name` - (Optional, Forces new resource, String) The name of the new Bitbucket repository to create.  This parameter is required when creating a new repository, cloning, or forking a repository.  The value will be computed when linking to an existing repository.
	* `repo_url` - (Optional, Forces new resource, String) The URL of the bitbucket repository for this tool integration.  This parameter is required when linking to an existing repository.  The value will be computed when creating a new repository, cloning, or forking a repository.
	* `source_repo_url` - (Optional, Forces new resource, String) The URL of the repository that you are forking or cloning.  This parameter is required when forking or cloning a repository.  It is not used when creating a new repository or linking to an existing repository.
	* `type` - (Required, Forces new resource, String) The operation that should be performed to initialize the new tool integration.  Use 'new' to create a new git repository, 'clone' to clone an existing repository into a new git repository, 'fork' to fork an existing git repository, or 'link' to link to an existing git repository.
	  * Constraints: Allowable values are: `new`, `fork`, `clone`, `link`, `new_if_not_exists`, `clone_if_not_exists`, `fork_if_not_exists`.
* `name` - (Optional, String) Name of the tool.
  * Constraints: The maximum length is `128` characters. The minimum length is `0` characters. The value must match regular expression `/^([^\\x00-\\x7F]|[a-zA-Z0-9-._ ])+$/`.
* `parameters` - (Required, List) Unique key-value pairs representing parameters to be used to create the tool. A list of parameters for each tool integration can be found in the <a href="https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-integrations">Configuring tool integrations page</a>.
Nested scheme for **parameters**:
	* `api_root_url` - (Computed, String) The API root URL for the Bitbucket Server.
	* `default_branch` - (Computed, String) The default branch of the git repository.
	* `enable_traceability` - (Optional, Boolean) Set this value to 'true' to track the deployment of code changes by creating tags, labels and comments on commits, pull requests and referenced issues.
	  * Constraints: The default value is `false`.
	* `git_id` - (Computed, String) Set this value to 'bitbucketgit' for bitbucket.org, or to the GUID of a custom Bitbucket server.
	* `integration_owner` - (Optional, String) Select the user which git operations will be performed as.
	* `owner_id` - (Computed, String) The Bitbucket user or group that owns the repository.  This parameter is required when creating a new repository, cloning, or forking a repository.  The value will be computed when linking to an existing repository.
	* `private_repo` - (Computed, Boolean) Set this value to 'true' to make the repository private when creating a new repository or when cloning or forking a repository.  This parameter is not used when linking to an existing repository.
	  * Constraints: The default value is `false`.
	* `project_id` - (Computed, String) The ID of the project.
	* `repo_name` - (Computed, String) The name of the new Bitbucket repository to create.  This parameter is required when creating a new repository, cloning, or forking a repository.  The value will be computed when linking to an existing repository.
	* `repo_url` - (Computed, String) The URL of the bitbucket repository for this tool integration.  This parameter is required when linking to an existing repository.  The value will be computed when creating a new repository, cloning, or forking a repository.
	* `source_repo_url` - (Computed, String) The URL of the repository that you are forking or cloning.  This parameter is required when forking or cloning a repository.  It is not used when creating a new repository or linking to an existing repository.
	* `token_url` - (Computed, String) The token URL used for authorizing with the Bitbucket server.
	* `toolchain_issues_enabled` - (Optional, Boolean) Setting this value to true will enable issues on the Bitbucket repository and add an issues tool card to the toolchain.  Setting the value to false will remove the tool card from the toolchain, but will not impact whether or not issues are enabled on the Bitbucket repository itself.
	  * Constraints: The default value is `true`.
	* `type` - (Computed, String) The operation that should be performed to initialize the new tool integration.  Use 'new' to create a new git repository, 'clone' to clone an existing repository into a new git repository, 'fork' to fork an existing git repository, or 'link' to link to an existing git repository.
	  * Constraints: Allowable values are: `new`, `fork`, `clone`, `link`, `new_if_not_exists`, `clone_if_not_exists`, `fork_if_not_exists`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain to bind the tool to.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the cd_toolchain_tool_bitbucketgit.
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

You can import the `ibm_cd_toolchain_tool_bitbucketgit` resource by using `id`.
The `id` property can be formed from `toolchain_id`, and `tool_id` in the following format:

```
<toolchain_id>/<tool_id>
```
* `toolchain_id`: A string. ID of the toolchain to bind the tool to.
* `tool_id`: A string. ID of the tool bound to the toolchain.

# Syntax
```
$ terraform import ibm_cd_toolchain_tool_bitbucketgit.cd_toolchain_tool_bitbucketgit <toolchain_id>/<tool_id>
```
