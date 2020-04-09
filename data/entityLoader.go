package data

import (
	"github.com/gogue-framework/gogue/ecs"
	"github.com/gogue-framework/gogue/ui"
	"reflect"
)

// EntityLoader responsible for taking loaded data from the text files (in the form of a map of strings), and turning
// those into entities, as required. For example, we may have loaded several potion definitions into memory from the
// definition files, and we now want to use those in the game. In order to do that, we would find the potion we want
// to load, and then take that definition and turn it into an entity in the ECS. We can do this as many times as we
// need per potion definition. In this way, we have an easy way of loading data file information into the ECS.
type EntityLoader struct {
	controller *ecs.Controller
}

// NewEntityLoader creates a new instance of an EntityLoader
func NewEntityLoader(controller *ecs.Controller) *EntityLoader {
	entityLoader := EntityLoader{}
	entityLoader.controller = controller

	return &entityLoader
}

// CreateSingleEntity takes a map of generic interface data (returned from Gogues data loader), and creates a single
// instance entity out of it. It will add any indicated components, and any data associated with those components. This
// will return the entity ID.
func (el *EntityLoader) CreateSingleEntity(data map[string]interface{}) int {
	// First, check to ensure there is a components property in the map. If this is not present, we cannot continue
	if _, ok := data["components"]; ok {
		componentList := data["components"].(map[string]interface{})

		// Create a new entity
		newEntity := el.controller.CreateEntity([]ecs.Component{})

		for componentName, values := range componentList {
			// Grab the component type off the controller. Also, ensure that this component type has been registered
			// with the controller
			component := el.controller.GetMappedComponentClass(componentName)

			if component != nil {
				newComponentValue := el.getInterfaceValue(component)

				valuesMap := values.(map[string]interface{})
				el.setFieldValues(valuesMap, newComponentValue)
				// Finally, update the new component with the changes we made based on the property values
				newComponentInterface := newComponentValue.Interface()
				updatedNewComponent := newComponentInterface

				// Add the component to the created entity
				el.controller.AddComponent(newEntity, updatedNewComponent.(ecs.Component))
			}
		}

		return newEntity
	}
	return -1
}

// getInterfaceValue returns the reflect.Value of a generic interface
func (el *EntityLoader) getInterfaceValue(reflectedInterface interface{}) reflect.Value {
	iType := reflect.TypeOf(reflectedInterface)
	iPointer := reflect.New(iType)
	iValue := iPointer.Elem()
	iInterface := iValue.Interface()
	newInterface := iInterface
	newInterfaceType := reflect.TypeOf(newInterface)
	newInterfaceValue := reflect.New(newInterfaceType).Elem()

	return newInterfaceValue
}

// setFieldValues will dynamically set interface property values for a given generic interface reflect.Value.
func (el *EntityLoader) setFieldValues(values map[string]interface{}, value reflect.Value) {
	for propertyName, propertyValue := range values {
		field := value.FieldByName(propertyName)

		if field.IsValid() && field.CanSet() {
			if field.Kind() == reflect.Int64 {
				field.SetInt(propertyValue.(int64))
			} else if field.Kind() == reflect.Float64 {
				field.SetFloat(propertyValue.(float64))
			} else if field.Kind() == reflect.String {
				field.SetString(propertyValue.(string))
			} else if field.Kind() == reflect.Interface {
				// There are only a few nested interface properties, so we'll just handle them manually
				if propertyName == "Glyph" {
					// Glyphs are a little weird. They don't expose any public setters, so we can't dynamically discover
					// and set their properties. We have to do it a bit more...manually
					glyphValues := propertyValue.(map[string]interface{})
					color := ""
					char := ""
					for glyphPropName, glyphPropValue := range glyphValues {
						if glyphPropName == "Color" {
							color = glyphPropValue.(string)
						} else if glyphPropName == "Char" {
							char = glyphPropValue.(string)
						}
					}
					glyph := ui.NewGlyph(char, color, "")
					field.Set(reflect.ValueOf(glyph))
				}
			}
		}
	}
}
