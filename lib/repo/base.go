package repo

import (
	"errors"
	"fmt"
	"strings"
)

const EQ = "="
const IN = "IN"

type Separator string

const AND Separator = " AND "
const OR Separator = " OR "
const COMMA Separator = " , "

type KeyValueField interface {
	KeyField() string
	ValueField() any
	Operator() string
}

type Attribute struct {
	Key   string
	Value any
}

func NewKV(key string, value any) *Attribute {
	return &Attribute{key, value}
}

func (attr *Attribute) KeyField() string {
	return attr.Key
}

func (attr *Attribute) ValueField() any {
	return attr.Value
}

func (attr *Attribute) Operator() string {
	return EQ
}

type InQuery[T any] struct {
	Field  string
	Values []T
}

func NewIn[T any](field string, values []T) *InQuery[T] {
	return &InQuery[T]{field, values}
}

func (in *InQuery[T]) KeyField() string {
	return in.Field
}

func (in *InQuery[T]) ValueField() any {
	var result []string

	for _, v := range in.Values {
		result = append(result, fmt.Sprintf("%v", v))
	}

	return strings.Join(result, ",")
}

func (in *InQuery[T]) Operator() string { return IN }

func buildCondQuery(conds []KeyValueField, start int, sep Separator) (condQuery string, values []any, dollar int) {
	if len(conds) == 0 {
		return "1=1", nil, 0
	}

	dollar = start
	var condString = []string{}

	for _, cond := range conds {
		dollar += 1

		var part string

		switch cond.Operator() {
		case IN:
			part = fmt.Sprintf("%s IN ($%d)", cond.KeyField(), dollar)
		default:
			part = fmt.Sprintf("%s %s $%d", cond.KeyField(), cond.Operator(), dollar)
		}

		condString = append(condString, part)
		values = append(values, cond.ValueField())
	}

	condQuery = strings.Join(condString, string(sep))

	dollar++

	return
}

var ErrorNotFound = errors.New("record_not_found")
var ErrorCountNotMatch = errors.New("count_not_match")

const InvalidCount = -1
