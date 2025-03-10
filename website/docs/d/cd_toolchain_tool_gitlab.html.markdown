---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_gitlab"
description: |-
  Get information about cd_toolchain_tool_gitlab
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_gitlab

Provides a read-only data source for cd_toolchain_tool_gitlab. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

More information on this Continuous Delivery tool integration can be found [here](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-gitlab).

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_gitlab" "cd_toolchain_tool_gitlab" {
	tool_id = "9603dcd4-3c86-44f8-8d0a-9427369878cf"
	toolchain_id = data.ibm_cd_toolchain.cd_toolchain.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `tool_id` - (Required, Forces new resource, String) ID of the tool bound to the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.
* `toolchain_id` - (Required, Forces new resource, String) ID of the toolchain.
  * Constraints: The maximum length is `36` characters. The minimum length is `36` characters. The value must match regular expression `/^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$/`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the cd_toolchain_tool_gitlab.
* `crn` - (String) Tool CRN.

* `href` - (String) URI representing the tool.

* `name` - (String) Tool name.

* `parameters` - (List) Unique key-value pairs representing parameters to be used to create the tool. A list of parameters for each tool integration can be found in the <a href="https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-integrations">Configuring tool integrations page</a>.
Nested scheme for **parameters**:
	* `api_root_url` - (String) The API root URL for the GitLab Server.
	* `default_branch` - (String) The default branch of the git repository.
	* `enable_traceability` - (Boolean) Set this value to 'true' to track the deployment of code changes by creating tags, labels and comments on commits, pull requests and referenced issues.
	  * Constraints: The default value is `false`.
	* `git_id` - (String) Set this value to 'gitlab' for gitlab.com, or to the GUID of a custom GitLab server.
	* `integration_owner` - (String) Select the user which git operations will be performed as.
	* `owner_id` - (String) The GitLab user or group that owns the repository.  This parameter is required when creating a new repository, cloning, or forking a repository.  The value will be computed when linking to an existing repository.
	* `private_repo` - (Boolean) Set this value to 'true' to make the repository private when creating a new repository or when cloning or forking a repository.  This parameter is not used when linking to an existing repository.
	  * Constraints: The default value is `true`.
	* `project_id` - (String) The ID of the project.
	* `repo_name` - (String) The name of the new GitLab repository to create.  This parameter is required when creating a new repository, cloning, or forking a repository.  The value will be computed when linking to an existing repository.
	* `repo_url` - (String) The URL of the GitLab repository for this tool integration.  This parameter is required when linking to an existing repository.  The value will be computed when creating a new repository, cloning, or forking a repository.
	* `source_repo_url` - (String) The URL of the repository that you are forking or cloning.  This parameter is required when forking or cloning a repository.  It is not used when creating a new repository or linking to an existing repository.
	* `token_url` - (String) The token URL used for authorizing with the GitLab server.
	* `toolchain_issues_enabled` - (Boolean) Setting this value to true will enable issues on the GitLab repository and add an issues tool card to the toolchain.  Setting the value to false will remove the tool card from the toolchain, but will not impact whether or not issues are enabled on the GitLab repository itself.
	  * Constraints: The default value is `true`.
	* `type` - (String) The operation that should be performed to initialize the new tool integration.  Use 'new' to create a new git repository, 'clone' to clone an existing repository into a new git repository, 'fork' to fork an existing git repository, or 'link' to link to an existing git repository.
	  * Constraints: Allowable values are: `new`, `fork`, `clone`, `link`, `new_if_not_exists`, `clone_if_not_exists`, `fork_if_not_exists`.

* `referent` - (List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (String) URI representing this resource through an API.
	* `ui_href` - (String) URI representing this resource through the UI.

* `resource_group_id` - (String) Resource group where the tool can be found.

* `state` - (String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (String) CRN of toolchain which the tool is bound to.


* `updated_at` - (String) Latest tool update timestamp.

