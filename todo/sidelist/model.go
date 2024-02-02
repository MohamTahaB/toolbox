package sidelist

import "github.com/charmbracelet/bubbles/list"

type SideList struct {
	List list.Model
}

type SideListItem struct {
	title, desc string
}

// Implement the item interface.
func (i SideListItem) Title() string {
	return i.title
}

func (i SideListItem) Desctiption() string {
	return i.desc
}

func (i SideListItem) FilterValue() string {
	return i.title
}
