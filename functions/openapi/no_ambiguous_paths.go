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

// AmbiguousPaths will determine if paths can be confused by a compiler.
type AmbiguousPaths struct {
}

// GetSchema returns a model.RuleFunctionSchema defining the schema of the AmbiguousPaths rule.
func (ap AmbiguousPaths) GetSchema() model.RuleFunctionSchema {
	return model.RuleFunctionSchema{Name: "ambiguousPaths"}
}

// RunRule will execute the AmbiguousPaths rule, based on supplied context and a supplied []*yaml.Node slice.
func (ap AmbiguousPaths) RunRule(nodes []*yaml.Node, context model.RuleFunctionContext) []model.RuleFunctionResult {

	if len(nodes) <= 0 {
		return nil
	}

	var results []model.RuleFunctionResult
	var seen []string

	ops := context.Index.GetPathsNode()

	var opPath string

	if ops != nil {
		for i, op := range ops.Content {
			if i%2 == 0 {
				opPath = op.Value
				continue
			}
			path := fmt.Sprintf("$.paths['%s']", opPath)
			for _, p := range seen {
				ambiguous := checkPaths(p, opPath)
				if ambiguous {

					results = append(results, model.RuleFunctionResult{
						Message:   fmt.Sprintf("paths are ambiguous with one another: `%s` and `%s`", p, opPath),
						StartNode: op,
						EndNode:   vacuumUtils.BuildEndNode(op),
						Path:      path,
						Rule:      context.Rule,
					})

				}
			}
			seen = append(seen, opPath)

		}
	}
	return results
}

var reggie, _ = regexp.Compile(`^{.+?}$`)

func checkPaths(pA, pB string) bool {
	segsA := strings.Split(pA, "/")[1:]
	segsB := strings.Split(pB, "/")[1:]

	if len(segsA) != len(segsB) {
		return false
	}

	a := 0
	b := 0
	amb := true
	for i, part := range segsA {
		aVar := reggie.MatchString(part)
		bVar := reggie.MatchString(segsB[i])
		if aVar || bVar {
			if aVar {
				a++
			}
			if bVar {
				b++
			}
			continue
		} else {
			if segsA[i] != segsB[i] {
				amb = false
			}
		}
	}
	return amb && a == b
}
