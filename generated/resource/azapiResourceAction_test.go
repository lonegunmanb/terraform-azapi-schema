package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azapi-schema/v2/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzapiResourceActionSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzapiResourceActionSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
