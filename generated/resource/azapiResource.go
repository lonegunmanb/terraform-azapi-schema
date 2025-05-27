package resource

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azapiResource = `{
  "block": {
    "attributes": {
      "body": {
        "computed": true,
        "description": "A dynamic attribute that contains the request body.",
        "description_kind": "markdown",
        "optional": true,
        "type": "dynamic"
      },
      "create_headers": {
        "description": "A mapping of headers to be sent with the create request.",
        "description_kind": "markdown",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "create_query_parameters": {
        "description": "A mapping of query parameters to be sent with the create request.",
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
      "delete_headers": {
        "description": "A mapping of headers to be sent with the delete request.",
        "description_kind": "markdown",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "delete_query_parameters": {
        "description": "A mapping of query parameters to be sent with the delete request.",
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
      "id": {
        "computed": true,
        "description": "The ID of the Azure resource.",
        "description_kind": "markdown",
        "type": "string"
      },
      "ignore_casing": {
        "computed": true,
        "description": "Whether ignore the casing of the property names in the response body. Defaults to ` + "`" + `false` + "`" + `.",
        "description_kind": "markdown",
        "optional": true,
        "type": "bool"
      },
      "ignore_missing_property": {
        "computed": true,
        "description": "Whether ignore not returned properties like credentials in ` + "`" + `body` + "`" + ` to suppress plan-diff. Defaults to ` + "`" + `true` + "`" + `. It's recommend to enable this option when some sensitive properties are not returned in response body, instead of setting them in ` + "`" + `lifecycle.ignore_changes` + "`" + ` because it will make the sensitive fields unable to update.",
        "description_kind": "markdown",
        "optional": true,
        "type": "bool"
      },
      "location": {
        "computed": true,
        "description": "The location of the Azure resource.",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "locks": {
        "description": "A list of ARM resource IDs which are used to avoid create/modify/delete azapi resources at the same time.",
        "description_kind": "markdown",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "name": {
        "computed": true,
        "description": "Specifies the name of the azure resource. Changing this forces a new resource to be created.",
        "description_kind": "markdown",
        "optional": true,
        "type": "string"
      },
      "output": {
        "computed": true,
        "description": "The output HCL object containing the properties specified in ` + "`" + `response_export_values` + "`" + `. Here are some examples to use the values.\n\n\t` + "`" + `` + "`" + `` + "`" + `terraform\n\t// it will output \"registry1.azurecr.io\"\n\toutput \"login_server\" {\n\t\tvalue = azapi_resource.example.output.properties.loginServer\n\t}\n\n\t// it will output \"disabled\"\n\toutput \"quarantine_policy\" {\n\t\tvalue = azapi_resource.example.output.properties.policies.quarantinePolicy.status\n\t}\n\t` + "`" + `` + "`" + `` + "`" + `\n",
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
      "read_headers": {
        "description": "A mapping of headers to be sent with the read request.",
        "description_kind": "markdown",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "read_query_parameters": {
        "description": "A mapping of query parameters to be sent with the read request.",
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
      "replace_triggers_external_values": {
        "description": "Will trigger a replace of the resource when the value changes and is not ` + "`" + `null` + "`" + `. This can be used by practitioners to force a replace of the resource when certain values change, e.g. changing the SKU of a virtual machine based on the value of variables or locals. The value is a ` + "`" + `dynamic` + "`" + `, so practitioners can compose the input however they wish. For a \"break glass\" set the value to ` + "`" + `null` + "`" + ` to prevent the plan modifier taking effect. \nIf you have ` + "`" + `null` + "`" + ` values that you do want to be tracked as affecting the resource replacement, include these inside an object. \nAdvanced use cases are possible and resource replacement can be triggered by values external to the resource, for example when a dependent resource changes.\n\ne.g. to replace a resource when either the SKU or os_type attributes change:\n\n` + "`" + `` + "`" + `` + "`" + `hcl\nresource \"azapi_resource\" \"example\" {\n  name      = var.name\n  type      = \"Microsoft.Network/publicIPAddresses@2023-11-01\"\n  parent_id = \"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/example\"\n  body = {\n    properties = {\n      sku   = var.sku\n      zones = var.zones\n    }\n  }\n\n  replace_triggers_external_values = [\n    var.sku,\n    var.zones,\n  ]\n}\n` + "`" + `` + "`" + `` + "`" + `\n",
        "description_kind": "markdown",
        "optional": true,
        "type": "dynamic"
      },
      "replace_triggers_refs": {
        "description": "A list of paths in the current Terraform configuration. When the values at these paths change, the resource will be replaced.",
        "description_kind": "markdown",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
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
      "schema_validation_enabled": {
        "computed": true,
        "description": "Whether enabled the validation on ` + "`" + `type` + "`" + ` and ` + "`" + `body` + "`" + ` with embedded schema. Defaults to ` + "`" + `true` + "`" + `.",
        "description_kind": "markdown",
        "optional": true,
        "type": "bool"
      },
      "sensitive_body": {
        "description": "A dynamic attribute that contains the write-only properties of the request body. This will be merge-patched to the body to construct the actual request body.",
        "description_kind": "markdown",
        "optional": true,
        "type": "dynamic",
        "write_only": true
      },
      "tags": {
        "computed": true,
        "description": "A mapping of tags which should be assigned to the Azure resource.",
        "description_kind": "markdown",
        "optional": true,
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
      },
      "update_headers": {
        "description": "A mapping of headers to be sent with the update request.",
        "description_kind": "markdown",
        "optional": true,
        "type": [
          "map",
          "string"
        ]
      },
      "update_query_parameters": {
        "description": "A mapping of query parameters to be sent with the update request.",
        "description_kind": "markdown",
        "optional": true,
        "type": [
          "map",
          [
            "list",
            "string"
          ]
        ]
      }
    },
    "block_types": {
      "identity": {
        "block": {
          "attributes": {
            "identity_ids": {
              "description": "A list of User Managed Identity ID's which should be assigned to the azure resource.",
              "description_kind": "markdown",
              "optional": true,
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
              "description": "The Type of Identity which should be used for this azure resource. Possible values are ` + "`" + `SystemAssigned` + "`" + `, ` + "`" + `UserAssigned` + "`" + ` and ` + "`" + `SystemAssigned,UserAssigned` + "`" + `",
              "description_kind": "markdown",
              "required": true,
              "type": "string"
            }
          },
          "description_kind": "plain"
        },
        "nesting_mode": "list"
      },
      "timeouts": {
        "block": {
          "attributes": {
            "create": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "delete": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "read": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours). Read operations occur during any refresh or planning operation when refresh is enabled.",
              "description_kind": "plain",
              "optional": true,
              "type": "string"
            },
            "update": {
              "description": "A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as \"30s\" or \"2h45m\". Valid time units are \"s\" (seconds), \"m\" (minutes), \"h\" (hours).",
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
    "description": "This resource can manage any Azure Resource Manager resource.",
    "description_kind": "markdown"
  },
  "version": 2
}`

func AzapiResourceSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azapiResource), &result)
	return &result
}
