package data_test

import (
	"testing"

	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azapi-schema/generated/data"
	"github.com/stretchr/testify/assert"
)

func TestAzapiClientConfigSchema(t *testing.T) {
	defaultSchema := &tfjson.Schema{}
	s := data.AzapiClientConfigSchema()
	assert.NotNil(t, s)
	assert.NotEqual(t, defaultSchema, s)
}
