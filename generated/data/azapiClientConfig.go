package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azapiClientConfig = `{
  "block": {
    "attributes": {
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "object_id": {
        "computed": true,
        "description": "The object ID of the identity. E.g. ` + "`" + `00000000-0000-0000-0000-000000000000` + "`" + `",
        "description_kind": "markdown",
        "type": "string"
      },
      "subscription_id": {
        "computed": true,
        "description": "The subscription ID. E.g. ` + "`" + `00000000-0000-0000-0000-000000000000` + "`" + `",
        "description_kind": "markdown",
        "type": "string"
      },
      "subscription_resource_id": {
        "computed": true,
        "description": "The resource ID of the subscription. E.g. ` + "`" + `/subscriptions/00000000-0000-0000-0000-000000000000` + "`" + `",
        "description_kind": "markdown",
        "type": "string"
      },
      "tenant_id": {
        "computed": true,
        "description": "The tenant ID. E.g. ` + "`" + `00000000-0000-0000-0000-000000000000` + "`" + `",
        "description_kind": "markdown",
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

func AzapiClientConfigSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azapiClientConfig), &result)
	return &result
}
