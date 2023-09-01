package generated

import (
	tfjson "github.com/hashicorp/terraform-json"
	"github.com/lonegunmanb/terraform-azapi-schema/generated/data"
	"github.com/lonegunmanb/terraform-azapi-schema/generated/resource"
)

var Resources map[string]*tfjson.Schema
var DataSources map[string]*tfjson.Schema

func init() {
	resources := make(map[string]*tfjson.Schema, 0)
	dataSources := make(map[string]*tfjson.Schema, 0)  
	resources["azapi_data_plane_resource"] = resource.AzapiDataPlaneResourceSchema()  
	resources["azapi_resource"] = resource.AzapiResourceSchema()  
	resources["azapi_resource_action"] = resource.AzapiResourceActionSchema()  
	resources["azapi_update_resource"] = resource.AzapiUpdateResourceSchema()  
	dataSources["azapi_resource"] = data.AzapiResourceSchema()  
	dataSources["azapi_resource_action"] = data.AzapiResourceActionSchema()  
	dataSources["azapi_resource_id"] = data.AzapiResourceIdSchema()  
	dataSources["azapi_resource_list"] = data.AzapiResourceListSchema()  
	Resources = resources
	DataSources = dataSources
}