// Copyright 2020-2021 Dave Shanley / Quobix
// SPDX-License-Identifier: MIT

package core

import (
	"fmt"
	"github.com/daveshanley/vacuum/model"
	vacuumUtils "github.com/daveshanley/vacuum/utils"
	"github.com/pb33f/libopenapi/utils"
	"gopkg.in/yaml.v3"
	"strings"
)

// Enumeration is a rule that will check that a set of values meet the supplied 'values' supplied via functionOptions.
type Enumeration struct{}

// GetSchema returns a model.RuleFunctionSchema defining the schema of the Enumeration rule.
func (e Enumeration) GetSchema() model.RuleFunctionSchema {
	return model.RuleFunctionSchema{
		Name:     "enumeration",
		Required: []string{"values"},
		Properties: []model.RuleFunctionProperty{
			{
				Name:        "values",
				Description: "'enumeration' requires a set of values to operate against",
			},
		},
		MinProperties: 1,
		MaxProperties: 2,
		ErrorMessage:  "'enumerate' needs 'values' to operate. A valid example of 'values' are: 'cake, egg, milk'",
	}
}

// RunRule will execute the Enumeration rule, based on supplied context and a supplied []*yaml.Node slice.
func (e Enumeration) RunRule(nodes []*yaml.Node, context model.RuleFunctionContext) []model.RuleFunctionResult {

	if len(nodes) != 1 { // there can only be a single node passed in to this function.
		return nil
	}

	var results []model.RuleFunctionResult
	var values []string

	message := context.Rule.Message

	// check supplied values (required)
	props := utils.ConvertInterfaceIntoStringMap(context.Options)
	if props["values"] == "" {
		return nil
	}
	values = strings.Split(props["values"], ",")

	pathValue := "unknown"
	if path, ok := context.Given.(string); ok {
		pathValue = path
	}

	ruleMessage := context.Rule.Description
	if context.Rule.Message != "" {
		ruleMessage = context.Rule.Message
	}

	for _, node := range nodes {
		if !e.checkValueAgainstAllowedValues(node.Value, values) {
			results = append(results, model.RuleFunctionResult{
				Message: vacuumUtils.SuppliedOrDefault(message,
					fmt.Sprintf("%s: `%s` must equal to one of: %v", ruleMessage,
						node.Value, values)),
				StartNode: node,
				EndNode:   vacuumUtils.BuildEndNode(node),
				Path:      pathValue,
				Rule:      context.Rule,
			})
		}
	}
	return results
}

func (e Enumeration) checkValueAgainstAllowedValues(value string, allowed []string) bool {
	found := false
	for _, allowedValue := range allowed {
		if strings.TrimSpace(allowedValue) == strings.TrimSpace(value) {
			found = true
			break
		}
	}
	return found
}
