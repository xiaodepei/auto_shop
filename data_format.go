package main

type Datatype struct {
	Device string
	Data   []Shoplist
	Time   string
}

type Shoplist struct {
	Epc []string
}
