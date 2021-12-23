package core

import (
	"github.com/daveshanley/vaccum/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCasing_RunRule_CamelSuccess(t *testing.T) {

	sampleYaml := `beer: "isYummy"`

	path := "$.beer"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "camel"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 0)
}

func TestCasing_RunRule_CamelFail(t *testing.T) {

	sampleYaml := `beer: "ISGREAT"`

	path := "$.beer"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "camel"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
}

func TestCasing_RunRule_PascalSuccess(t *testing.T) {

	sampleYaml := `spaghetti: "IsMyFav"`

	path := "$.spaghetti"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "pascal"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 0)
}

func TestCasing_RunRule_PascalFail(t *testing.T) {

	sampleYaml := `spaghetti: "is-the-best"`

	path := "$.spaghetti"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "pascal"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
}

func TestCasing_RunRule_KebabSuccess(t *testing.T) {

	sampleYaml := `melody: "is-what-makes-life-worth-living"`

	path := "$.melody"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "kebab"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 0)
}

func TestCasing_RunRule_KebabFail(t *testing.T) {

	sampleYaml := `melody: "is_what-Makes-life_worth-living"`

	path := "$.melody"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "kebab"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
}

func TestCasing_RunRule_CobolSuccess(t *testing.T) {

	sampleYaml := `maddy: "THE-LITTLE-CHAMPION"`

	path := "$.maddy"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "cobol"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 0)
}

func TestCasing_RunRule_CobolFail(t *testing.T) {

	sampleYaml := `maddy: "THE-little-CHAMPION"`

	path := "$.maddy"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "cobol"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
}

func TestCasing_RunRule_SnakeSuccess(t *testing.T) {

	sampleYaml := `ember: "naughty_puppy_get_off_the_couch"`

	path := "$.ember"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "snake"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 0)
}

func TestCasing_RunRule_SnakeFail(t *testing.T) {

	sampleYaml := `ember: "Naughty_ember-get-off-THAT_COUCH"`

	path := "$.ember"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "snake"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
}

func TestCasing_RunRule_MacroSuccess(t *testing.T) {

	sampleYaml := `chicken: "THE_NANNY_DOG"`

	path := "$.chicken"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "macro"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 0)
}

func TestCasing_RunRule_MacroFail(t *testing.T) {

	sampleYaml := `chicken: "THE-Nanny_dog"`

	path := "$.chicken"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "macro"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
}

func TestCasing_RunRule_CamelNoDigits_Success(t *testing.T) {

	sampleYaml := `alchomohol: "afterHoursNoDigits"`

	path := "$.alchomohol"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "camel"
	opts["disallowDigits"] = "true"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 0)
}

func TestCasing_RunRule_CamelNoDigits_Fail(t *testing.T) {

	sampleYaml := `alchomohol: "aft3rHoursN0Dig1ts"`

	path := "$.alchomohol"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "camel"
	opts["disallowDigits"] = "true"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
}

func TestCasing_RunRule_Snake_SeparatingChar_Success(t *testing.T) {

	sampleYaml := `alchomohol: "after_hours,want_a_drink"`

	path := "$.alchomohol"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "snake"
	opts["separator.char"] = ","

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 0)
}

func TestCasing_RunRule_Snake_SeparatingChar_Fail(t *testing.T) {

	sampleYaml := `alchomohol: "after_hours|want_a_drink"`

	path := "$.alchomohol"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "snake"
	opts["separator.char"] = ","

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
}

func TestCasing_RunRule_Snake_AllowLeading_Success(t *testing.T) {

	sampleYaml := `mo_money: ",mo_problems,rub_a,dub_dub"`

	path := "$.mo_money"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "snake"
	opts["separator.char"] = ","
	opts["separator.allowLeading"] = "true"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 0)
}

func TestCasing_RunRule_Snake_AllowLeading_Fail(t *testing.T) {

	sampleYaml := `mo_money: ",mo_problems,rub_a,dub_dub"`

	path := "$.mo_money"

	nodes, _ := utils.FindNodes([]byte(sampleYaml), path)
	assert.Len(t, nodes, 1)

	opts := make(map[string]string)
	opts["type"] = "snake"
	opts["separator.char"] = ","
	opts["separator.allowLeading"] = "false"

	rule := buildCoreTestRule(path, severityError, "casing", "", nil)
	ctx := buildCoreTestContext(rule.Then, opts)

	def := &Casing{}
	res := def.RunRule(nodes, ctx)

	assert.Len(t, res, 1)
}
