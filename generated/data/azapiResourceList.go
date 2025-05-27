package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azapiResourceList = `{
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
      "output": {
        "computed": true,
        "description": "The output HCL object containing the properties specified in ` + "`" + `response_export_values` + "`" + `. Here are some examples to use the values.\n\n\t` + "`" + `` + "`" + `` + "`" + `terraform\n\t// it will output \"registry1.azurecr.io\"\n\toutput \"login_server\" {\n\t\tvalue = data.azapi_resource_list.example.output.properties.loginServer\n\t}\n\n\t// it will output \"disabled\"\n\toutput \"quarantine_policy\" {\n\t\tvalue = data.azapi_resource_list.example.output.properties.policies.quarantinePolicy.status\n\t}\n\t` + "`" + `` + "`" + `` + "`" + `\n",
        "description_kind": "markdown",
        "type": "dynamic"
      },
      "parent_id": {
        "description": "The ID of the azure resource in which this resource is created. It supports different kinds of deployment scope for **top level** resources:\n\n  - resource group scope: ` + "`" + `parent_id` + "`" + ` should be the ID of a resource group, it's recommended to manage a resource group by azurerm_resource_group.\n\t- management group scope: ` + "`" + `parent_id` + "`" + ` should be the ID of a management group, it's recommended to manage a management group by azurerm_management_group.\n\t- extension scope: ` + "`" + `parent_id` + "`" + ` should be the ID of the resource you're adding the extension to.\n\t- subscription scope: ` + "`" + `parent_id` + "`" + ` should be like \\x60/subscriptions/00000000-0000-0000-0000-000000000000\\x60\n\t- tenant scope: ` + "`" + `parent_id` + "`" + ` should be /\n\n  For child level resources, the ` + "`" + `parent_id` + "`" + ` should be the ID of its parent resource, for example, subnet resource's ` + "`" + `parent_id` + "`" + ` is the ID of the vnet.\n\n  For type ` + "`" + `Microsoft.Resources/resourceGroups` + "`" + `, the ` + "`" + `parent_id` + "`" + ` could be omitted, it defaults to subscription ID specified in provider or the default subscription (You could check the default subscription by azure cli command: ` + "`" + `az account show` + "`" + `).",
        "description_kind": "markdown",
        "required": true,
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
      "response_export_values": {
        "description": "The attribute can accept either a list or a map.\n\n- **List**: A list of paths that need to be exported from the response body. Setting it to ` + "`" + `[\"*\"]` + "`" + ` will export the full response body. Here's an example. If it sets to ` + "`" + `[\"value\"]` + "`" + `, it will set the following HCL object to the computed property output.\n\n\t` + "`" + `` + "`" + `` + "`" + `text\n\t{\n\t  \"value\" = [\n\t\t{\n\t\t  \"id\" = \"/subscriptions/000000/resourceGroups/demo-rg/providers/Microsoft.Automation/automationAccounts/example\"\n\t\t  \"location\" = \"eastus2\"\n\t\t  \"name\" = \"example\"\n\t\t  \"properties\" = {\n\t\t\t\"creationTime\" = \"2024-10-11T08:18:38.737+00:00\"\n\t\t\t\"disableLocalAuth\" = false\n\t\t\t\"lastModifiedTime\" = \"2024-10-11T08:18:38.737+00:00\"\n\t\t\t\"publicNetworkAccess\" = true\n\t\t  }\n\t\t  \"tags\" = {}\n\t\t  \"type\" = \"Microsoft.Automation/AutomationAccounts\"\n\t\t}\n\t  ]\n\t}\n\t` + "`" + `` + "`" + `` + "`" + `\n\n- **Map**: A map where the key is the name for the result and the value is a JMESPath query string to filter the response. Here's an example. If it sets to ` + "`" + `{\"values\": \"value[].{name: name, publicNetworkAccess: properties.publicNetworkAccess}\", \"names\": \"value[].name\"}` + "`" + `, it will set the following HCL object to the computed property output.\n\n\t` + "`" + `` + "`" + `` + "`" + `text\n\t{\n\t\t\"names\" = [\n\t\t\t\"example\",\n\t\t\t\"fredaccount01\",\n\t\t]\n\t\t\"values\" = [\n\t\t\t{\n\t\t\t  \"name\" = \"example\"\n\t\t\t  \"publicNetworkAccess\" = true\n\t\t\t},\n\t\t\t{\n\t\t\t  \"name\" = \"fredaccount01\"\n\t\t\t  \"publicNetworkAccess\" = null\n\t\t\t},\n\t\t]\n\t}\n\t` + "`" + `` + "`" + `` + "`" + `\n\nTo learn more about JMESPath, visit [JMESPath](https://jmespath.org/).\n",
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
    "description_kind": "plain"
  },
  "version": 0
}`

func AzapiResourceListSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azapiResourceList), &result)
	return &result
}
