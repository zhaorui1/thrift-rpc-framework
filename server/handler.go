package server

import (
	"AddressService/gen/gen-go/addrservice"
	"fmt"
)

type AddressServiceHandler struct {
}

func NewAddressServiceHandler() *AddressServiceHandler {
	return &AddressServiceHandler{}
}

func (p *AddressServiceHandler) GetAllAddress(req *addrservice.Request) (res *addrservice.Response, err error) {
	fmt.Print("Resquest:", req.GetId(), "\n")

	res = addrservice.NewResponse()
	res.Code = 200
	res.Desc = "OK"
	res.Data = "asdasd"

	return res, err
}
