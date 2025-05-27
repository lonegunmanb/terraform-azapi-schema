package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azapiResource = `{
  "block": {
    "attributes": {
      "headers": {
        "description": "A map of headers to include in the request",
        "description_kind": "markdown",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "id": {
        "computed": true,
        "description": "The ID of the Azure resource.",
        "description_kind": "markdown",
        "type": "string"
      },
      "identity": {
        "computed": true,
        "description_kind": "plain",
        "nested_type": {
          "attributes": {
            "identity_ids": {
              "computed": true,
              "description": "A list of User Managed Identity ID's which should be assigned to the azure resource.",
              "description_kind": "markdown",
              "type": [
                "list",
                "string"
              ]
            },
            "principal_id": {
              "computed": true,
              "description": "The Principal ID for the Service Principal associated with the Managed Service Identity of this Azure resource.",
              "description_kind": "markdown",
              "type": "string"
            },
            "tenant_id": {
              "computed": true,
              "description": "The Tenant ID for the Service Principal associated with the Managed Service Identity of this Azure resource.",
              "description_kind": "markdown",
              "type": "string"
            },
            "type": {
              "computed": true,
              "description": "The Type of Identity which should be used for this azure resource. Possible values are ` + "`" + `SystemAssigned` + "`" + `, ` + "`" + `UserAssigned` + "`" + ` and ` + "`" + `SystemAssigned,UserAssigned` + "`" + `",
              "description_kind": "markdown",
              "type": "string"
            }
          },
          "nesting_mode": "list"
        }
      },
      "location": {
        "computed": true,
        "description": "The location of the Azure resource.",
        "description_kind": "markdown",
        "type": "string"
      },
      "name": {
        "computed": true,
        "description": "Specifies the name of the Azure resource. Exactly one of the arguments ` + "`" + `name` + "`" + ` or ` + "`" + `resource_id` + "`" + ` must be set. It could be omitted if the ` + "`" + `type` + "`" + ` is ` + "`" + `Microsoft.Resources/subscriptions` + "`" + `.",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "output": {
        "computed": true,
        "description": "The output HCL object containing the properties specified in ` + "`" + `response_export_values` + "`" + `. Here are some examples to use the values.\n\n\t` + "`" + `` + "`" + `` + "`" + `terraform\n\t// it will output \"registry1.azurecr.io\"\n\toutput \"login_server\" {\n\t\tvalue = data.azapi_resource.example.output.properties.loginServer\n\t}\n\n\t// it will output \"disabled\"\n\toutput \"quarantine_policy\" {\n\t\tvalue = data.azapi_resource.example.output.properties.policies.quarantinePolicy.status\n\t}\n\t` + "`" + `` + "`" + `` + "`" + `\n",
        "description_kind": "markdown",
        "type": "dynamic"
      },
      "parent_id": {
        "computed": true,
        "description": "The ID of the azure resource in which this resource is created. It supports different kinds of deployment scope for **top level** resources:\n\n  - resource group scope: ` + "`" + `parent_id` + "`" + ` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.\n\t- management group scope: ` + "`" + `parent_id` + "`" + ` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.\n\t- extension scope: ` + "`" + `parent_id` + "`" + ` should be the ID of the resource you're adding the extension to.\n\t- subscription scope: ` + "`" + `parent_id` + "`" + ` should be like \\x60/subscriptions/00000000-0000-0000-0000-000000000000\\x60\n\t- tenant scope: ` + "`" + `parent_id` + "`" + ` should be /\n\n  For child level resources, the ` + "`" + `parent_id` + "`" + ` should be the ID of its parent resource, for example, subnet resource's ` + "`" + `parent_id` + "`" + ` is the ID of the vnet.\n\n  For type ` + "`" + `Microsoft.Resources/resourceGroups` + "`" + `, the ` + "`" + `parent_id` + "`" + ` could be omitted, it defaults to subscription ID specified in provider or the default subscription (You could check the default subscription by azure cli command: ` + "`" + `az account show` + "`" + `).",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "query_parameters": {
        "description": "A map of query parameters to include in the request",
        "description_kind": "markdown",
        "optional": true,
        "type": [
          "map",
          [
            "list",
            "string"
          ]
        ]
      },
      "resource_id": {
        "computed": true,
        "description": "The ID of the Azure resource to retrieve. Exactly one of the arguments ` + "`" + `name` + "`" + ` or ` + "`" + `resource_id` + "`" + ` must be set. It could be omitted if the ` + "`" + `type` + "`" + ` is ` + "`" + `Microsoft.Resources/subscriptions` + "`" + `.",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "response_export_values": {
        "description": "The attribute can accept either a list or a map.\n\n- **List**: A list of paths that need to be exported from the response body. Setting it to ` + "`" + `[\"*\"]` + "`" + ` will export the full response body. Here's an example. If it sets to ` + "`" + `[\"properties.loginServer\", \"properties.policies.quarantinePolicy.status\"]` + "`" + `, it will set the following HCL object to the computed property output.\n\n\t` + "`" + `` + "`" + `` + "`" + `text\n\t{\n\t\tproperties = {\n\t\t\tloginServer = \"registry1.azurecr.io\"\n\t\t\tpolicies = {\n\t\t\t\tquarantinePolicy = {\n\t\t\t\t\tstatus = \"disabled\"\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\t` + "`" + `` + "`" + `` + "`" + `\n\n- **Map**: A map where the key is the name for the result and the value is a JMESPath query string to filter the response. Here's an example. If it sets to ` + "`" + `{\"login_server\": \"properties.loginServer\", \"quarantine_status\": \"properties.policies.quarantinePolicy.status\"}` + "`" + `, it will set the following HCL object to the computed property output.\n\n\t` + "`" + `` + "`" + `` + "`" + `text\n\t{\n\t\t\"login_server\" = \"registry1.azurecr.io\"\n\t\t\"quarantine_status\" = \"disabled\"\n\t}\n\t` + "`" + `` + "`" + `` + "`" + `\n\nTo learn more about JMESPath, visit [JMESPath](https://jmespath.org/).\n",
        "description_kind": "markdown",
        "optional": true,
        "type": "dynamic"
      },
      "retry": {
        "description": "The retry object supports the following attributes:",
        "description_kind": "markdown",
        "nested_type": {
          "attributes": {
            "error_message_regex": {
              "description": "A list of regular expressions to match against error messages. If any of the regular expressions match, the request will be retried.",
              "description_kind": "markdown",
              "required": true,
              "type": [
                "list",
                "string"
              ]
            },
            "interval_seconds": {
              "computed": true,
              "description": "The base number of seconds to wait between retries. Default is ` + "`" + `10` + "`" + `.",
              "description_kind": "markdown",
              "optional": true,
              "type": "number"
            },
            "max_interval_seconds": {
              "computed": true,
              "description": "The maximum number of seconds to wait between retries. Default is ` + "`" + `180` + "`" + `.",
              "description_kind": "markdown",
              "optional": true,
              "type": "number"
            },
            "multiplier": {
              "computed": true,
              "description": "The multiplier to apply to the interval between retries. Default is ` + "`" + `1.5` + "`" + `.",
              "description_kind": "markdown",
              "optional": true,
              "type": "number"
            },
            "randomization_factor": {
              "computed": true,
              "description": "The randomization factor to apply to the interval between retries. The formula for the randomized interval is: ` + "`" + `RetryInterval * (random value in range [1 - RandomizationFactor, 1 + RandomizationFactor])` + "`" + `. Therefore set to zero ` + "`" + `0.0` + "`" + ` for no randomization. Default is ` + "`" + `0.5` + "`" + `.",
              "description_kind": "markdown",
              "optional": true,
              "type": "number"
            }
          },
          "nesting_mode": "single"
        },
        "optional": true
      },
      "tags": {
        "computed": true,
        "description": "A mapping of tags which are assigned to the Azure resource.",
        "description_kind": "markdown",
        "type": [
          "map",
          "string"
        ]
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
    "description": "This resource can access any existing Azure resource manager resource.",
    "description_kind": "markdown"
  },
  "version": 0
}`

func AzapiResourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azapiResource), &result)
	return &result
}
