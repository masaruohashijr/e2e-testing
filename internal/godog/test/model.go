package test

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
)

type StepDefinition struct {
	Args         []interface{}
	Expr         *regexp.Regexp
	HandlerValue reflect.Value
}

type Background struct {
	Steps []StepDefinition
}

type Suite struct {
	Steps      []StepDefinition
	Background []StepDefinition
}

type ScenarioContext struct {
	Name  string
	Suite Suite
}

type FeatureContext struct {
	Name string
	Scs  []ScenarioContext
}

func (ctx *ScenarioContext) Step(expr, stepFunc interface{}) {
	var regex *regexp.Regexp
	switch t := expr.(type) {
	case *regexp.Regexp:
		regex = t
	case string:
		regex = regexp.MustCompile(t)
	case []byte:
		regex = regexp.MustCompile(string(t))
	default:
		panic(fmt.Sprintf("expecting expr to be a *regexp.Regexp or a string, got type: %T", expr))
	}
	v := reflect.ValueOf(stepFunc)
	typ := v.Type()
	if typ.Kind() != reflect.Func {
		panic(fmt.Sprintf("expected handler to be func, but got: %T", stepFunc))
	}

	if typ.NumOut() != 1 {
		panic(fmt.Sprintf("expected handler to return only one value, but it has: %d", typ.NumOut()))
	}

	def := StepDefinition{
		Expr:         regex,
		HandlerValue: v,
	}
	ctx.Suite.Steps = append(ctx.Suite.Steps, def)
}

func (sd *StepDefinition) Run() interface{} {
	typ := sd.HandlerValue.Type()
	var values []reflect.Value
	for i := 0; i < typ.NumIn(); i++ {
		param := typ.In(i)
		switch param.Kind() {
		case reflect.Int:
			s, err := sd.shouldBeString(i)
			if err != nil {
				return err
			}
			v, err := strconv.ParseInt(s, 10, 0)
			if err != nil {
				return fmt.Errorf(`cannot convert argument %d: "%s" to int: %s`, i, s, err)
			}
			values = append(values, reflect.ValueOf(int(v)))
		case reflect.String:
			s, err := sd.shouldBeString(i)
			if err != nil {
				return err
			}
			values = append(values, reflect.ValueOf(s))
		}
	}
	return sd.HandlerValue.Call(values)[0].Interface()
}

func (sd *StepDefinition) shouldBeString(idx int) (string, error) {
	arg := sd.Args[idx]
	s, ok := arg.(string)
	if !ok {
		return "", fmt.Errorf(`cannot convert argument %d: "%v" of type "%T" to string`, idx, arg, arg)
	}
	return s, nil
}
