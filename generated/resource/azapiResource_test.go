package resource_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azapi-schema/generated/resource"
	"github.com/stretchr/testify/assert"
)

func TestAzapiResourceSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := resource.AzapiResourceSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
