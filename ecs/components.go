package ecs

import (
	"reflect"
)

type Component interface {
	TypeOf() reflect.Type
}
