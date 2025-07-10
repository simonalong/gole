package matcher

import (
	"fmt"
	"github.com/expr-lang/expr"
	"github.com/expr-lang/expr/vm"
	"github.com/simonalong/gole/constants"
	"github.com/simonalong/gole/logger"
	"reflect"
	"strings"
)

type ConditionMatch struct {
	BlackWhiteMatch

	expression string
	Program    *vm.Program
}

func (conditionMatch *ConditionMatch) Match(_ map[string]interface{}, object any, field reflect.StructField, fieldValue any) bool {
	env := map[string]any{
		"root":    object,
		"current": fieldValue,
	}

	output, err := expr.Run(conditionMatch.Program, env)
	if err != nil {
		logger.Errorf("表达式 %v 执行失败: %v", conditionMatch.expression, err.Error())
		return false
	}

	result, err := CastBool(fmt.Sprintf("%v", output))
	if err != nil {
		return false
	}

	if result {
		conditionMatch.SetBlackMsg("属性 %v 的值 %v 命中禁用条件 [%v] ", field.Name, fieldValue, conditionMatch.expression)
	} else {
		conditionMatch.SetWhiteMsg("属性 %v 的值 %v 不符合条件 [%v] ", field.Name, fieldValue, conditionMatch.expression)
	}
	return result
}

func (conditionMatch *ConditionMatch) IsEmpty() bool {
	return conditionMatch.Program == nil
}

func BuildConditionMatcher(objectTypeFullName string, fieldKind reflect.Kind, objectFieldName string, tagName string, subCondition string, errCode, errMsg string) {
	if constants.MATCH != tagName {
		return
	}

	if fieldKind == reflect.Slice {
		return
	}
	if !strings.Contains(subCondition, constants.Condition) || !strings.Contains(subCondition, constants.EQUAL) {
		return
	}

	index := strings.Index(subCondition, constants.EQUAL)
	expression := subCondition[index+1:]

	if expression == "" {
		return
	}

	program, err := expr.Compile(rmvWell(expression))
	if err != nil {
		logger.Errorf("脚本: %v 编译异常：%v", expression, err.Error())
		return
	}
	addMatcher(objectTypeFullName, objectFieldName, &ConditionMatch{Program: program, expression: expression}, errCode, errMsg, true)
}
