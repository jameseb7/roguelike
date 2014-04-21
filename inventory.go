package main

import "github.com/jameseb7/roguelike/entity"

type item struct{
	name string
	id entity.ID
	checked bool
}

var inventory [52]*item
var invIDs map[entity.ID]byte = make(map[entity.ID]byte, 52)

func inventoryIndex(char byte) byte{
	if ('a' <= char) && (char <= 'z') {
		return char - 'a'
	}
	if ('A' <= char) && (char <= 'Z') {
		return char - 'A' + byte(26)
	}
	return 255
}

func inventoryChar(index byte) byte{
	if (0 <= index) && (index <= 25) {
		return index + 'a'
	}
	if (26 <= index) && (index <= 51) {
		return index + 'A' - byte(26)
	}
	return 0
}

func updateInventory(){
	ids := player.ListContents()

	//set all the items unchecked
	for i, _ := range inventory {
		if inventory[i] != nil {
			inventory[i].checked = false
		}
	}

	//check the existing items that are still there
	for _, id := range ids {
		if i := inventoryIndex(invIDs[id]); i != 255 {
			inventory[i].id = id
			inventory[i].checked = true
		}
	}

	//remove unchecked items
	for i, v := range inventory {
		if inventory[i] != nil {
			if inventory[i].checked == false {
				delete(invIDs, v.id)
				inventory[i] = nil
			}
		}
	}

	//add in any remaining items
	for _, id := range ids {
		if inventoryIndex(invIDs[id]) == 255 {
			for i, _ := range inventory {
				if inventory[i] == nil {
					inventory[i] = new(item)
					inventory[i].id = id
					inventory[i].checked = true
					inventory[i].name = currentLevel.EntityByID(id).EntityName()
					invIDs[id] = inventoryChar(byte(i))
					break
				}
			}
		}
	}
	
	//update item names
	for _, v := range inventory {
		if v != nil {
			v.name = currentLevel.EntityByID(v.id).EntityName()
		}
	}
}