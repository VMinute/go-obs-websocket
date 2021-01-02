package obsws

// This file is automatically generated.
// https://github.com/christopher-dG/go-obs-websocket/blob/master/codegen/protocol.py

// SourceOrderChangedEvent : Scene items within a scene have been reordered.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#sourceorderchanged
type SourceOrderChangedEvent struct {
	// Name of the scene where items have been reordered.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	// Ordered list of scene items.
	// Required: Yes.
	SceneItems []map[string]interface{} `json:"scene-items"`
	// Item source name.
	// Required: Yes.
	SceneItemsSourceName string `json:"scene-items.*.source-name"`
	// Scene item unique ID.
	// Required: Yes.
	SceneItemsItemID int `json:"scene-items.*.item-id"`
	_event           `json:",squash"`
}

// SceneItemAddedEvent : A scene item has been added to a scene.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#sceneitemadded
type SceneItemAddedEvent struct {
	// Name of the scene.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	// Name of the item added to the scene.
	// Required: Yes.
	ItemName string `json:"item-name"`
	// Scene item ID.
	// Required: Yes.
	ItemID int `json:"item-id"`
	_event `json:",squash"`
}

// SceneItemRemovedEvent : A scene item has been removed from a scene.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#sceneitemremoved
type SceneItemRemovedEvent struct {
	// Name of the scene.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	// Name of the item removed from the scene.
	// Required: Yes.
	ItemName string `json:"item-name"`
	// Scene item ID.
	// Required: Yes.
	ItemID int `json:"item-id"`
	_event `json:",squash"`
}

// SceneItemVisibilityChangedEvent : A scene item's visibility has been toggled.
//
// Since obs-websocket version: 4.0.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#sceneitemvisibilitychanged
type SceneItemVisibilityChangedEvent struct {
	// Name of the scene.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	// Name of the item in the scene.
	// Required: Yes.
	ItemName string `json:"item-name"`
	// Scene item ID.
	// Required: Yes.
	ItemID int `json:"item-id"`
	// New visibility state of the item.
	// Required: Yes.
	ItemVisible bool `json:"item-visible"`
	_event      `json:",squash"`
}

// SceneItemLockChangedEvent : A scene item's locked status has been toggled.
//
// Since obs-websocket version: 4.8.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#sceneitemlockchanged
type SceneItemLockChangedEvent struct {
	// Name of the scene.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	// Name of the item in the scene.
	// Required: Yes.
	ItemName string `json:"item-name"`
	// Scene item ID.
	// Required: Yes.
	ItemID int `json:"item-id"`
	// New locked state of the item.
	// Required: Yes.
	ItemLocked bool `json:"item-locked"`
	_event     `json:",squash"`
}

// SceneItemTransformChangedEvent : A scene item's transform has been changed.
//
// Since obs-websocket version: 4.6.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#sceneitemtransformchanged
type SceneItemTransformChangedEvent struct {
	// Name of the scene.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	// Name of the item in the scene.
	// Required: Yes.
	ItemName string `json:"item-name"`
	// Scene item ID.
	// Required: Yes.
	ItemID int `json:"item-id"`
	// Scene item transform properties.
	// Required: Yes.
	Transform *SceneItemTransform `json:"transform"`
	_event    `json:",squash"`
}

// SceneItemSelectedEvent : A scene item is selected.
//
// Since obs-websocket version: 4.6.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#sceneitemselected
type SceneItemSelectedEvent struct {
	// Name of the scene.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	// Name of the item in the scene.
	// Required: Yes.
	ItemName string `json:"item-name"`
	// Name of the item in the scene.
	// Required: Yes.
	ItemID int `json:"item-id"`
	_event `json:",squash"`
}

// SceneItemDeselectedEvent : A scene item is deselected.
//
// Since obs-websocket version: 4.6.0.
//
// https://github.com/Palakis/obs-websocket/blob/4.x-current/docs/generated/protocol.md#sceneitemdeselected
type SceneItemDeselectedEvent struct {
	// Name of the scene.
	// Required: Yes.
	SceneName string `json:"scene-name"`
	// Name of the item in the scene.
	// Required: Yes.
	ItemName string `json:"item-name"`
	// Name of the item in the scene.
	// Required: Yes.
	ItemID int `json:"item-id"`
	_event `json:",squash"`
}
