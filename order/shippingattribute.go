package order

type ShippingAttribute struct {
	Id                 int    `json:",omitempty"`
	CustomerShippingId int    `json:",omitempty"`
	Text1              string `json:",omitempty"`
	Text2              string `json:",omitempty"`
	Text3              string `json:",omitempty"`
	Text4              string `json:",omitempty"`
	Text5              string `json:",omitempty"`
	Text6              string `json:",omitempty"`
}