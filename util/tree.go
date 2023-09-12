package util

import "gateway/models"

type TreeItem interface {
	GetID() uint
	GetParentID() *uint
}

type FieldMapperFunc func(item TreeItem) map[string]interface{}

type IDFunc func(item TreeItem) uint

type ParentIDFunc func(item TreeItem) *uint

func ConvertToTree(items interface{}, fieldMapper FieldMapperFunc) []map[string]interface{} {
	var treeItems []TreeItem

	switch v := items.(type) {
	case []models.Menu:
		for _, item := range v {
			treeItems = append(treeItems, item)
		}
	case []models.Route:
		for _, item := range v {
			treeItems = append(treeItems, item)
		}
	default:
		panic("Unsupported item type for ConvertToTree")
	}

	var transformedItems []map[string]interface{}
	itemMap := make(map[uint]*map[string]interface{})

	for _, item := range treeItems {
		transformedItem := fieldMapper(item)
		itemID := item.GetID()
		itemMap[itemID] = &transformedItem
	}

	for _, item := range treeItems {
		parentID := item.GetParentID()
		if parentID != nil && itemMap[*parentID] != nil {
			parentItem := itemMap[*parentID]
			if children, ok := (*parentItem)["children"].([]map[string]interface{}); ok {
				(*parentItem)["children"] = append(children, *itemMap[item.GetID()])
			}
		}
	}

	for _, item := range treeItems {
		if item.GetParentID() == nil {
			transformedItems = append(transformedItems, *itemMap[item.GetID()])
		}
	}

	return transformedItems
}

func MapMenuToTreeItem(item TreeItem) map[string]interface{} {
	menu := item.(models.Menu) // 类型断言
	return map[string]interface{}{
		"id":       menu.ID,
		"name":     menu.Name,
		"value":    menu.ID,
		"label":    menu.Name,
		"path":     menu.Path,
		"icon":     menu.Icon,
		"status":   menu.Status,
		"children": []map[string]interface{}{},
	}
}

func MapRouteToTreeItem(item TreeItem) map[string]interface{} {
	route := item.(models.Route) // 类型断言
	return map[string]interface{}{
		"id":       route.ID,
		"value":    route.ID,
		"name":     route.Name,
		"label":    route.Name,
		"path":     route.Path,
		"status":   route.Status,
		"children": []map[string]interface{}{},
	}
}
