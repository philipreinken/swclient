package article

type PropertyGroup struct {
	Id         int    `json:",omitempty"`
	Name       string `json:",omitempty"`
	Position   int    `json:",omitempty"`
	Comparable bool   `json:",omitempty"`
	SortMode   int    `json:",omitempty"`
}
