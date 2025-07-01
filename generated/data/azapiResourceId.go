package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azapiResourceId = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description": "The ID of the Azure resource.",
        "description_kind": "markdown",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "The name of the Azure resource.",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "parent_id": {
        "computed": true,
        "description": "The ID of the azure resource in which this resource is created. It supports different kinds of deployment scope for **top level** resources:\n\n  - resource group scope: ` + "`" + `parent_id` + "`" + ` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.\n\t- management group scope: ` + "`" + `parent_id` + "`" + ` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.\n\t- extension scope: ` + "`" + `parent_id` + "`" + ` should be the ID of the resource you're adding the extension to.\n\t- subscription scope: ` + "`" + `parent_id` + "`" + ` should be like \\x60/subscriptions/00000000-0000-0000-0000-000000000000\\x60\n\t- tenant scope: ` + "`" + `parent_id` + "`" + ` should be /\n\n  For child level resources, the ` + "`" + `parent_id` + "`" + ` should be the ID of its parent resource, for example, subnet resource's ` + "`" + `parent_id` + "`" + ` is the ID of the vnet.\n\n  For type ` + "`" + `Microsoft.Resources/resourceGroups` + "`" + `, the ` + "`" + `parent_id` + "`" + ` could be omitted, it defaults to subscription ID specified in provider or the default subscription (You could check the default subscription by azure cli command: ` + "`" + `az account show` + "`" + `).",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "parts": {
        "computed": true,
        "description": "A map of the resource ID parts, where the key is the part name and the value is the part value. e.g. ` + "`" + `virtualNetworks=myVnet` + "`" + `.",
        "description_kind": "markdown",
        "type": [
          "map",
          "string"
        ]
      },
      "provider_namespace": {
        "computed": true,
        "description": "The Azure resource provider namespace of the Azure resource.",
        "description_kind": "markdown",
        "type": "string"
      },
      "resource_group_name": {
        "computed": true,
        "description": "The resource group name of the Azure resource.",
        "description_kind": "markdown",
        "type": "string"
      },
      "resource_id": {
        "computed": true,
        "description": "The ID of an existing Azure source.",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "subscription_id": {
        "computed": true,
        "description": "The subscription ID of the Azure resource.",
        "description_kind": "markdown",
        "type": "string"
      },
      "type": {
        "description": "In a format like ` + "`" + `\u003cresource-type\u003e@\u003capi-version\u003e` + "`" + `. ` + "`" + `\u003cresource-type\u003e` + "`" + ` is the Azure resource type, for example, ` + "`" + `Microsoft.Storage/storageAccounts` + "`" + `. ` + "`" + `\u003capi-version\u003e` + "`" + ` is version of the API used to manage this azure resource.",
        "description_kind": "markdown",
        "required": true,
        "type": "string"
      }
    },
    "block_types": {
      "timeouts": {
        "block": {
          "attributes": {
            "read": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "single"
      }
    },
    "description": "This data source allows you to parse an Azure Resource ID into its components. It can be used to extract information such as the resource type, name, parent ID, and other parts of the resource ID. It can also be used to build a resource ID from the resource type, name, and parent ID. This is useful for constructing resource IDs for use in other resources or data sources.",
    "description_kind": "markdown"
  },
  "version": 0
}`

func AzapiResourceIdSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azapiResourceId), &result)
	return &result
}
