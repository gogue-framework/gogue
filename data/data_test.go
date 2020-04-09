package data

import (
	"github.com/gogue-framework/gogue/ecs"
	"github.com/gogue-framework/gogue/ui"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestNewFileLoader(t *testing.T) {
	fileLoader, err := NewFileLoader("testdata")

	assert.Nil(t, err, "NewFileLoader raised an error")
	assert.Equal(t, fileLoader.dataFilesLocation, "testdata")

	fileLoader, err = NewFileLoader("does/not/exist")

	assert.NotNil(t, err, "NewFileLoader did not raise an error")
	assert.Nil(t, fileLoader)
}

func TestFileLoader_LoadDataFromFile(t *testing.T) {
	fileLoader, _ := NewFileLoader("testdata")
	dataFile := "enemies.json"

	dataMap, err := fileLoader.LoadDataFromFile(dataFile)

	assert.Nil(t, err, "err is not Nil")
	assert.Equal(t, len(dataMap), 1)

	levelOne := dataMap["level_1"].(map[string]interface{})

	assert.Equal(t, len(levelOne), 3)

	smallRat := levelOne["small_rat"].(map[string]interface{})
	components := smallRat["components"].(map[string]interface{})
	appearance := components["appearance"].(map[string]interface{})

	assert.Equal(t, appearance["Name"], "Small rat")

	// Ensure a non-existant file will properly raise an error
	dataMap, err = fileLoader.LoadDataFromFile("does_not_exist.json")

	assert.Nil(t, dataMap, "dataMap is not nil")
	assert.NotNil(t, err)
}

func TestFileLoader_LoadAllFromFiles(t *testing.T) {
	fileLoader, _ := NewFileLoader("testdata")

	dataMap, err := fileLoader.LoadAllFromFiles()

	assert.Nil(t, err, "err is not Nil")
	assert.Equal(t, len(dataMap), 2)

	levelOne := dataMap["testdata/enemies"]["level_1"].(map[string]interface{})
	levelTwo := dataMap["testdata/enemies_2"]["level_2"].(map[string]interface{})

	assert.Equal(t, len(levelOne), 3)
	assert.Equal(t, len(levelTwo), 3)

}

func TestNewEntityLoader(t *testing.T) {
	controller := ecs.NewController()
	entityLoader := NewEntityLoader(controller)

	assert.NotNil(t, entityLoader)
	assert.Equal(t, entityLoader.controller, controller)
}

// Components for EntityLoader tests
type PositionComponent struct {
	X int
	Y int
}

func (pc PositionComponent) TypeOf() reflect.Type {
	return reflect.TypeOf(pc)
}

type AppearanceComponent struct {
	Name        string
	Description string
	Glyph       ui.Glyph
	Layer       int
}

func (ac AppearanceComponent) TypeOf() reflect.Type {
	return reflect.TypeOf(ac)
}

func TestEntityLoader_CreateSingleEntity(t *testing.T) {
	controller := ecs.NewController()

	// Load a couple of components into the controller
	controller.MapComponentClass("position", PositionComponent{})
	controller.MapComponentClass("appearance", AppearanceComponent{})

	dataLoader, _ := NewFileLoader("testdata")
	entityLoader := NewEntityLoader(controller)

	dataMap, _ := dataLoader.LoadDataFromFile("enemies.json")
	levelOne := dataMap["level_1"].(map[string]interface{})
	smallRat := levelOne["small_rat"].(map[string]interface{})
	caveBat := levelOne["cave_bat"].(map[string]interface{})

	entityID := entityLoader.CreateSingleEntity(smallRat)

	assert.Equal(t, entityID, 0)
	assert.True(t, controller.HasComponent(entityID, reflect.TypeOf(PositionComponent{})))
	assert.True(t, controller.HasComponent(entityID, reflect.TypeOf(AppearanceComponent{})))

	appearance := controller.GetComponent(entityID, AppearanceComponent{}.TypeOf()).(AppearanceComponent)

	assert.Equal(t, "Small rat", appearance.Name)
	assert.Equal(t, "r", appearance.Glyph.Char())
	assert.Equal(t, "brown", appearance.Glyph.Color())

	entityID = entityLoader.CreateSingleEntity(caveBat)

	assert.Equal(t, entityID, 1)
	assert.True(t, controller.HasComponent(entityID, reflect.TypeOf(PositionComponent{})))
	assert.True(t, controller.HasComponent(entityID, reflect.TypeOf(AppearanceComponent{})))

	appearance = controller.GetComponent(entityID, AppearanceComponent{}.TypeOf()).(AppearanceComponent)

	assert.Equal(t, "Cave bat", appearance.Name)
	assert.Equal(t, "b", appearance.Glyph.Char())
	assert.Equal(t, "gray", appearance.Glyph.Color())
}
