// Copyright 2022 Dave Shanley / Quobix
// SPDX-License-Identifier: MIT

package openapi

import (
	"fmt"
	"github.com/daveshanley/vacuum/model"
	vacuumUtils "github.com/daveshanley/vacuum/utils"
	"gopkg.in/yaml.v3"
	"regexp"
	"strings"
)

// PathsKebabCase Checks to ensure each segment of a path is using kebab case.
type PathsKebabCase struct {
}

// GetSchema returns a model.RuleFunctionSchema defining the schema of the VerbsInPath rule.
func (vp PathsKebabCase) GetSchema() model.RuleFunctionSchema {
	return model.RuleFunctionSchema{Name: "pathsKebabCase"}
}

// RunRule will execute the PathsKebabCase rule, based on supplied context and a supplied []*yaml.Node slice.
func (vp PathsKebabCase) RunRule(nodes []*yaml.Node, context model.RuleFunctionContext) []model.RuleFunctionResult {

	if len(nodes) <= 0 {
		return nil
	}

	var results []model.RuleFunctionResult

	ops := context.Index.GetPathsNode()

	var opPath string

	if ops != nil {
		for i, op := range ops.Content {
			if i%2 == 0 {
				opPath = op.Value
				continue
			}
			path := fmt.Sprintf("$.paths['%s']", opPath)
			if opPath == "/" {
				continue
			}
			notKebab, segments := checkPathCase(opPath)
			if notKebab {
				results = append(results, model.RuleFunctionResult{
					Message:   fmt.Sprintf("path segments `%s` do not use kebab-case", strings.Join(segments, "`, `")),
					StartNode: op,
					EndNode:   vacuumUtils.BuildEndNode(op),
					Path:      path,
					Rule:      context.Rule,
				})
			}
		}
	}
	return results
}

var pathKebabCaseRegex, _ = regexp.Compile(`^[{}a-z\d-.]+$`)
var variableRegex, _ = regexp.Compile(`^\{(\w.*)}\.?.*$`)

func checkPathCase(path string) (bool, []string) {
	segs := strings.Split(path, "/")[1:]
	var found []string
	for _, seg := range segs {
		if !pathKebabCaseRegex.MatchString(seg) {
			// check if it's a variable, if so, skip
			if seg == "" {
				continue
			}
			// if this is a variable, or a variable at the end of a path then skip
			if variableRegex.MatchString(seg) {
				continue
			}
			found = append(found, seg)
		}
	}
	if len(found) > 0 {
		return true, found
	}
	return false, nil
}
