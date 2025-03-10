---
layout: "ibm"
page_title: "IBM : ibm_cd_toolchain_tool_pagerduty"
description: |-
  Get information about cd_toolchain_tool_pagerduty
subcategory: "Continuous Delivery"
---

# ibm_cd_toolchain_tool_pagerduty

Provides a read-only data source for cd_toolchain_tool_pagerduty. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

More information on this Continuous Delivery tool integration can be found [here](https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-pagerduty).

## Example Usage

```hcl
data "ibm_cd_toolchain_tool_pagerduty" "cd_toolchain_tool_pagerduty" {
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

* `id` - The unique identifier of the cd_toolchain_tool_pagerduty.
* `crn` - (String) Tool CRN.

* `href` - (String) URI representing the tool.

* `name` - (String) Tool name.

* `parameters` - (List) Unique key-value pairs representing parameters to be used to create the tool. A list of parameters for each tool integration can be found in the <a href="https://cloud.ibm.com/docs/ContinuousDelivery?topic=ContinuousDelivery-integrations">Configuring tool integrations page</a>.
Nested scheme for **parameters**:
	* `service_id` - (String) The service ID of the PagerDuty service.
	* `service_key` - (String) The PagerDuty service integration key. You can find or create this key in the Integrations section of the PagerDuty service page.
	* `service_url` - (String) The URL of the PagerDuty service to post alerts to.

* `referent` - (List) Information on URIs to access this resource through the UI or API.
Nested scheme for **referent**:
	* `api_href` - (String) URI representing this resource through an API.
	* `ui_href` - (String) URI representing this resource through the UI.

* `resource_group_id` - (String) Resource group where the tool can be found.

* `state` - (String) Current configuration state of the tool.
  * Constraints: Allowable values are: `configured`, `configuring`, `misconfigured`, `unconfigured`.

* `toolchain_crn` - (String) CRN of toolchain which the tool is bound to.


* `updated_at` - (String) Latest tool update timestamp.

