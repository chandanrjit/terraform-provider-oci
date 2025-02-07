---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_notebook_session"
sidebar_current: "docs-oci-datasource-datascience-notebook_session"
description: |-
  Provides details about a specific Notebook Session in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_notebook_session
This data source provides details about a specific Notebook Session resource in Oracle Cloud Infrastructure Data Science service.

Gets the specified notebook session's information.

## Example Usage

```hcl
data "oci_datascience_notebook_session" "test_notebook_session" {
	#Required
	notebook_session_id = oci_datascience_notebook_session.test_notebook_session.id
}
```

## Argument Reference

The following arguments are supported:

* `notebook_session_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the notebook session.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the notebook session's compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the notebook session.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information. Example: `My NotebookSession` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the notebook session.
* `lifecycle_details` - Details about the state of the notebook session.
* `notebook_session_config_details` - Details for the notebook session configuration.
	* `block_storage_size_in_gbs` - A notebook session instance is provided with a block storage volume. This specifies the size of the volume in GBs. 
	* `notebook_session_shape_config_details` - Details for the notebook session shape configuration.
		* `memory_in_gbs` - A notebook session instance of type VM.Standard.E3.Flex allows memory to be specified. This specifies the size of the memory in GBs. 
		* `ocpus` - A notebook session instance of type VM.Standard.E3.Flex allows the ocpu count to be specified. 
	* `shape` - The shape used to launch the notebook session compute instance.  The list of available shapes in a given compartment can be retrieved using the `ListNotebookSessionShapes` endpoint. 
	* `subnet_id` - A notebook session instance is provided with a VNIC for network access.  This specifies the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT gateway for egress to the internet. 
* `notebook_session_configuration_details` - Details for the notebook session configuration.
	* `block_storage_size_in_gbs` - A notebook session instance is provided with a block storage volume. This specifies the size of the volume in GBs. 
	* `notebook_session_shape_config_details` - Details for the notebook session shape configuration.
		* `memory_in_gbs` - A notebook session instance of type VM.Standard.E3.Flex allows memory to be specified. This specifies the size of the memory in GBs. 
		* `ocpus` - A notebook session instance of type VM.Standard.E3.Flex allows the ocpu count to be specified. 
	* `shape` - The shape used to launch the notebook session compute instance.  The list of available shapes in a given compartment can be retrieved using the `ListNotebookSessionShapes` endpoint. 
	* `subnet_id` - A notebook session instance is provided with a VNIC for network access.  This specifies the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create a VNIC in.  The subnet should be in a VCN with a NAT gateway for egress to the internet. 
* `notebook_session_runtime_config_details` - Notebook Session runtime configuration details.
	* `custom_environment_variables` - Custom environment variables for Notebook Session. These key-value pairs will be available for customers in Notebook Sessions.
	* `notebook_session_git_config_details` - Git configuration Details.
		* `notebook_session_git_repo_config_collection` - A collection of Git repository configurations.
			* `url` - The repository URL
* `notebook_session_url` - The URL to interact with the notebook session.
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the notebook session.
* `state` - The state of the notebook session.
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 

