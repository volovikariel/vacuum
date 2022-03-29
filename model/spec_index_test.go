package model

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"testing"
)

func TestSpecIndex_ExtractRefsStripe(t *testing.T) {

	stripe, _ := ioutil.ReadFile("test_files/stripe.yaml")
	var rootNode yaml.Node
	yaml.Unmarshal(stripe, &rootNode)

	index := NewSpecIndex(&rootNode)

	assert.Len(t, index.allRefs, 385)
	assert.Len(t, index.allMappedRefs, 385)

	combined := index.GetAllCombinedReferences()
	assert.Equal(t, 537, len(combined))

	assert.Len(t, index.rawSequencedRefs, 1972)
	assert.Equal(t, 246, index.pathCount)
	assert.Equal(t, 402, index.operationCount)
	assert.Equal(t, 537, index.schemaCount)
	assert.Equal(t, 0, index.globalTagsCount)
	assert.Equal(t, 0, index.globalLinksCount)
	assert.Equal(t, 0, index.componentParamCount)
	assert.Equal(t, 143, index.operationParamCount)
	assert.Equal(t, 76, index.componentsInlineParamDuplicateCount)
	assert.Equal(t, 67, index.componentsInlineParamUniqueCount)

}

func TestSpecIndex_Asana(t *testing.T) {

	asana, _ := ioutil.ReadFile("test_files/asana.yaml")
	var rootNode yaml.Node
	yaml.Unmarshal(asana, &rootNode)

	index := NewSpecIndex(&rootNode)

	assert.Len(t, index.allRefs, 152)
	assert.Len(t, index.allMappedRefs, 152)
	combined := index.GetAllCombinedReferences()
	assert.Equal(t, 171, len(combined))
	assert.Equal(t, 118, index.pathCount)
	assert.Equal(t, 152, index.operationCount)
	assert.Equal(t, 135, index.schemaCount)
	assert.Equal(t, 26, index.globalTagsCount)
	assert.Equal(t, 0, index.globalLinksCount)
	assert.Equal(t, 30, index.componentParamCount)
	assert.Equal(t, 107, index.operationParamCount)
	assert.Equal(t, 8, index.componentsInlineParamDuplicateCount)
	assert.Equal(t, 69, index.componentsInlineParamUniqueCount)
}

func TestSpecIndex_k8s(t *testing.T) {

	asana, _ := ioutil.ReadFile("test_files/k8s.json")
	var rootNode yaml.Node
	yaml.Unmarshal(asana, &rootNode)

	index := NewSpecIndex(&rootNode)

	assert.Len(t, index.allRefs, 558)
	assert.Len(t, index.allMappedRefs, 558)
	combined := index.GetAllCombinedReferences()
	assert.Equal(t, 563, len(combined))
	assert.Equal(t, 436, index.pathCount)
	assert.Equal(t, 853, index.operationCount)
	assert.Equal(t, 563, index.schemaCount)
	assert.Equal(t, 0, index.globalTagsCount)
	assert.Equal(t, 58, index.operationTagsCount)
	assert.Equal(t, 0, index.globalLinksCount)
	assert.Equal(t, 0, index.componentParamCount)
	assert.Equal(t, 36, index.operationParamCount)
	assert.Equal(t, 26, index.componentsInlineParamDuplicateCount)
	assert.Equal(t, 10, index.componentsInlineParamUniqueCount)
	assert.Equal(t, 58, index.GetTotalTagsCount())
	assert.Equal(t, 2524, index.GetRawReferenceCount())

}

func TestSpecIndex_PetstoreV2(t *testing.T) {

	asana, _ := ioutil.ReadFile("test_files/petstorev2.json")
	var rootNode yaml.Node
	yaml.Unmarshal(asana, &rootNode)

	index := NewSpecIndex(&rootNode)

	assert.Len(t, index.allRefs, 6)
	assert.Len(t, index.allMappedRefs, 6)
	assert.Equal(t, 14, index.pathCount)
	assert.Equal(t, 20, index.operationCount)
	assert.Equal(t, 6, index.schemaCount)
	assert.Equal(t, 3, index.globalTagsCount)
	assert.Equal(t, 3, index.operationTagsCount)
	assert.Equal(t, 0, index.globalLinksCount)
	assert.Equal(t, 1, index.componentParamCount)
	assert.Equal(t, 1, index.GetComponentParameterCount())
	assert.Equal(t, 11, index.operationParamCount)
	assert.Equal(t, 4, index.componentsInlineParamDuplicateCount)
	assert.Equal(t, 7, index.componentsInlineParamUniqueCount)
	assert.Equal(t, 3, index.GetTotalTagsCount())
}

func TestSpecIndex_PetstoreV3(t *testing.T) {

	asana, _ := ioutil.ReadFile("test_files/petstorev3.json")
	var rootNode yaml.Node
	yaml.Unmarshal(asana, &rootNode)

	index := NewSpecIndex(&rootNode)

	assert.Len(t, index.allRefs, 7)
	assert.Len(t, index.allMappedRefs, 7)
	assert.Equal(t, 13, index.pathCount)
	assert.Equal(t, 19, index.operationCount)
	assert.Equal(t, 8, index.schemaCount)
	assert.Equal(t, 3, index.globalTagsCount)
	assert.Equal(t, 3, index.operationTagsCount)
	assert.Equal(t, 0, index.globalLinksCount)
	assert.Equal(t, 0, index.componentParamCount)
	assert.Equal(t, 9, index.operationParamCount)
	assert.Equal(t, 3, index.componentsInlineParamDuplicateCount)
	assert.Equal(t, 6, index.componentsInlineParamUniqueCount)
	assert.Equal(t, 3, index.GetTotalTagsCount())

}

func TestSpecIndex_BurgerShop(t *testing.T) {

	burgershop, _ := ioutil.ReadFile("test_files/burgershop.openapi.yaml")
	var rootNode yaml.Node
	yaml.Unmarshal(burgershop, &rootNode)

	index := NewSpecIndex(&rootNode)

	assert.Len(t, index.allRefs, 4)
	assert.Len(t, index.allMappedRefs, 4)
	assert.Equal(t, 4, len(index.GetMappedReferences()))

	assert.Equal(t, 5, index.pathCount)
	assert.Equal(t, 5, index.GetPathCount())

	assert.Equal(t, 5, len(index.GetAllSchemas()))

	assert.Equal(t, 17, len(index.GetAllSequencedReferences()))
	assert.NotNil(t, index.GetSchemasNode())
	assert.Nil(t, index.GetParametersNode())

	assert.Equal(t, 5, index.operationCount)
	assert.Equal(t, 5, index.GetOperationCount())

	assert.Equal(t, 5, index.schemaCount)
	assert.Equal(t, 5, index.GetComponentSchemaCount())

	assert.Equal(t, 2, index.globalTagsCount)
	assert.Equal(t, 2, index.GetGlobalTagsCount())
	assert.Equal(t, 2, index.GetTotalTagsCount())

	assert.Equal(t, 2, index.operationTagsCount)
	assert.Equal(t, 2, index.GetOperationTagsCount())

	assert.Equal(t, 3, index.globalLinksCount)
	assert.Equal(t, 3, index.GetGlobalLinksCount())

	assert.Equal(t, 0, index.componentParamCount)
	assert.Equal(t, 0, index.GetComponentParameterCount())

	assert.Equal(t, 2, index.operationParamCount)
	assert.Equal(t, 2, index.GetOperationsParameterCount())

	assert.Equal(t, 1, index.componentsInlineParamDuplicateCount)
	assert.Equal(t, 1, index.GetInlineDuplicateParamCount())

	assert.Equal(t, 1, index.componentsInlineParamUniqueCount)
	assert.Equal(t, 1, index.GetInlineUniqueParamCount())

}

func TestSpecIndex_NoRoot(t *testing.T) {

	index := NewSpecIndex(nil)
	refs := index.ExtractRefs(nil, nil, 0, false, "")
	docs := index.ExtractExternalDocuments(nil)
	assert.Nil(t, docs)
	assert.Nil(t, refs)
	assert.Nil(t, index.FindComponent("nothing", nil))
	assert.Equal(t, -1, index.GetOperationCount())
	assert.Equal(t, -1, index.GetPathCount())
	assert.Equal(t, -1, index.GetGlobalTagsCount())
	assert.Equal(t, -1, index.GetOperationTagsCount())
	assert.Equal(t, -1, index.GetTotalTagsCount())
	assert.Equal(t, -1, index.GetOperationsParameterCount())
	assert.Equal(t, -1, index.GetComponentParameterCount())
	assert.Equal(t, -1, index.GetComponentSchemaCount())
	assert.Equal(t, -1, index.GetGlobalLinksCount())

}

func TestSpecIndex_BurgerShopMixedRef(t *testing.T) {

	spec, _ := ioutil.ReadFile("test_files/mixedref-burgershop.openapi.yaml")
	var rootNode yaml.Node
	yaml.Unmarshal(spec, &rootNode)

	index := NewSpecIndex(&rootNode)

	assert.Len(t, index.allRefs, 4)
	assert.Len(t, index.allMappedRefs, 4)
	assert.Equal(t, 5, index.GetPathCount())
	assert.Equal(t, 5, index.GetOperationCount())
	assert.Equal(t, 1, index.GetComponentSchemaCount())
	assert.Equal(t, 2, index.GetGlobalTagsCount())
	assert.Equal(t, 2, index.GetTotalTagsCount())
	assert.Equal(t, 2, index.GetOperationTagsCount())
	assert.Equal(t, 0, index.GetGlobalLinksCount())
	assert.Equal(t, 0, index.GetComponentParameterCount())
	assert.Equal(t, 2, index.GetOperationsParameterCount())
	assert.Equal(t, 1, index.GetInlineDuplicateParamCount())
	assert.Equal(t, 1, index.GetInlineUniqueParamCount())

}

func TestSpecIndex_TestEmptyBrokenReferences(t *testing.T) {

	asana, _ := ioutil.ReadFile("test_files/badref-burgershop.openapi.yaml")
	var rootNode yaml.Node
	yaml.Unmarshal(asana, &rootNode)

	index := NewSpecIndex(&rootNode)
	assert.Equal(t, 5, index.GetPathCount())
	assert.Equal(t, 5, index.GetOperationCount())
	assert.Equal(t, 5, index.GetComponentSchemaCount())
	assert.Equal(t, 2, index.GetGlobalTagsCount())
	assert.Equal(t, 2, index.GetTotalTagsCount())
	assert.Equal(t, 2, index.GetOperationTagsCount())
	assert.Equal(t, 2, index.GetGlobalLinksCount())
	assert.Equal(t, 0, index.GetComponentParameterCount())
	assert.Equal(t, 2, index.GetOperationsParameterCount())
	assert.Equal(t, 1, index.GetInlineDuplicateParamCount())
	assert.Equal(t, 1, index.GetInlineUniqueParamCount())
	assert.Len(t, index.refErrors, 7)

}