package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azapi-schema/v2/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzapiResourceIdSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzapiResourceIdSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
