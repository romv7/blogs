// @ineffective-test
package utils_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/rommms07/blogs/internal"
	"github.com/rommms07/blogs/internal/utils"
)

const (
	LOAD_TOOL   = 0
	RELOAD_TOOL = 1
	DROP_TOOL   = 2

	LOAD     = 0
	LOAD_ALL = 1

	RELOAD     = 0
	RELOAD_ALL = 1

	DROP     = 0
	DROP_ALL = 1
)

var (
	config     *internal.ConfigSchema
	schemaTool *utils.SQLSchemaTool
)

func init() {
	config, _ = internal.LoadConfig()
	schemaTool = utils.NewSQLSchemaTool()
}

func callTest(t *testing.T, tool int, method0 int, method1 int, args ...reflect.Value) error {
	if err := testSchemaToolOp(t, tool, method0, schemaTool, args...); err != nil {
		printErr(t, err)
		return err
	}

	query, _ := utils.GetExecQueryProps()

	switch tool {
	case LOAD_TOOL:
		fallthrough
	case RELOAD_TOOL:
		if query != strings.ReplaceAll(utils.TestQuery, utils.TestDb, utils.TestDb+"0") {
			var toolType string

			if tool == LOAD_TOOL {
				toolType = "Load"
			} else {
				toolType = "Reload"
			}

			t.Errorf("[fail] schemaTool.%s did not get the expected query...", toolType)
		}

	case DROP_TOOL:
		if query != fmt.Sprintf("drop table if exists `%s0`;", args[0].String()) {
			t.Errorf("[fail] schemaTool.Drop did not get the expected query... got: (%s)", query)
		}
	}

	if err := testSchemaToolOp(t, tool, method1, schemaTool); err != nil {
		printErr(t, err)
		return err
	}

	return nil
}

func testSchemaToolOp(t *testing.T, tool int, method int, st *utils.SQLSchemaTool, args ...reflect.Value) (err error) {
	vl := reflect.ValueOf(st)
	methodToCall := map[int][]string{
		LOAD_TOOL:   {"Load", "LoadAll"},
		RELOAD_TOOL: {"Reload", "ReloadAll"},
		DROP_TOOL:   {"Drop", "DropAll"},
	}

	if err_v := vl.MethodByName(methodToCall[tool][method]).Call(args); err_v[0].Interface() != nil {
		err = err_v[0].Interface().(error)
	}

	return
}

func printErr(t *testing.T, err error) {
	t.Errorf("[fail] %s", err.Error())
}

func setupMocks() {
	utils.LoadMockFunctions()
}

func detachMocks() {
	utils.Restore()
}

func Test_shouldCreateNewSchemaTool(t *testing.T) {
	setupMocks()
	tool := utils.NewSQLSchemaTool()
	if tool == nil {
		t.Errorf("[fail] utils.NewSQLSchemaTool was not able to create a new instance.")
	}
	detachMocks()
}

func Test_shouldLoadWithoutAnError(t *testing.T) {
	setupMocks()

	connName, name := utils.TestDb, utils.TestDb
	callTest(t, LOAD_TOOL, LOAD, LOAD_ALL, reflect.ValueOf(connName), reflect.ValueOf(name))

	detachMocks()
}

func Test_shouldReloadWithoutAnError(t *testing.T) {
	setupMocks()

	connName, name := utils.TestDb, utils.TestDb
	callTest(t, RELOAD_TOOL, RELOAD, RELOAD_ALL, reflect.ValueOf(connName), reflect.ValueOf(name))

	detachMocks()
}

func Test_shouldDropWithoutAnError(t *testing.T) {
	setupMocks()

	connName, name := utils.TestDb, utils.TestDb
	callTest(t, DROP_TOOL, DROP, DROP_ALL, reflect.ValueOf(connName), reflect.ValueOf(name))

	detachMocks()
}
