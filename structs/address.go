package main

type Address struct {
	stName  string
	houseNo string
}

func (addr Address) getCompleteAddress() string {
	return "HouseNo: " + addr.houseNo + ", " + addr.stName
}
