package entity

import "sort"

type Container interface {
	AddItem(eid ID)
	RemoveItem(eid ID)
	ListContents() []ID
}

type inventory map[ID]bool

func (inv inventory) AddItem(eid ID) {
	inv[eid] = true
}

func (inv inventory) RemoveItem(eid ID) {
	delete(inv,eid)
}

func (inv inventory) ListContents() []ID {
	var list = make([]ID, 0, len(inv))
	for k, v := range inv {
		if v == true {
			list = append(list, k)
		}
	}
	sort.Sort(idSlice(list))
	return list
}

type idSlice []ID
func (ids idSlice) Len() int {return len(ids)}
func (ids idSlice) Less(i,j int) bool {return (ids[i] < ids[j])}
func (ids idSlice) Swap(i,j int) {ids[i], ids[j] = ids[j], ids[i]}