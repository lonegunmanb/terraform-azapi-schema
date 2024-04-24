package data

import (
	"encoding/json"

	tfjson "github.com/hashicorp/terraform-json"
)

const azapiResourceAction = `{
  "block": {
    "attributes": {
      "action": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "body": {
        "description_kind": "plain",
        "optional": true,
        "type": "dynamic"
      },
      "id": {
        "computed": true,
        "description_kind": "plain",
        "type": "string"
      },
      "method": {
        "computed": true,
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "output": {
        "computed": true,
        "description_kind": "plain",
        "type": "dynamic"
      },
      "resource_id": {
        "description_kind": "plain",
        "optional": true,
        "type": "string"
      },
      "response_export_values": {
        "description_kind": "plain",
        "optional": true,
        "type": [
          "list",
          "string"
        ]
      },
      "type": {
        "description_kind": "plain",
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

func AzapiResourceActionSchema() *tfjson.Schema {
	var result tfjson.Schema
	_ = json.Unmarshal([]byte(azapiResourceAction), &result)
	return &result
}
