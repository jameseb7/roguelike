package entity

import "sort"
import "log"

type ConstContainer interface {
	ListContents() []ID
	ItemOfType(t ItemType) ID
}

type Container interface {
	ConstContainer
	AddItem(eid ID, t ItemType)
	RemoveItem(eid ID, t ItemType)
}

type inventory struct {
	items map[ID]bool
	types map[ItemType]ID
}

func newInventory() *inventory {
	inv := new(inventory)
	inv.items = make(map[ID]bool, 52)
	inv.types = make(map[ItemType]ID, 52)
	return inv
}

func (inv *inventory) AddItem(eid ID, t ItemType) {
	inv.items[eid] = true
	if inv.types[t] == 0 {
		inv.types[t] = eid
	}
	log.Println("Item", eid, "added. Inventory contents:", inv.ListContents())
}

func (inv *inventory) RemoveItem(eid ID, t ItemType) {
	delete(inv.items, eid)
	if inv.types[t] == eid {
		delete(inv.types, t)
	}
}

func (inv inventory) ListContents() []ID {
	var list = make([]ID, 0, len(inv.items))
	for k, v := range inv.items {
		if v == true {
			list = append(list, k)
		}
	}
	sort.Sort(idSlice(list))
	return list
}

func (inv inventory) ItemOfType(t ItemType) ID {
	return inv.types[t]
}

type idSlice []ID
func (ids idSlice) Len() int {return len(ids)}
func (ids idSlice) Less(i,j int) bool {return (ids[i] < ids[j])}
func (ids idSlice) Swap(i,j int) {ids[i], ids[j] = ids[j], ids[i]}