package server

import (
	"AddressService/gen/gen-go/addrservice"
	"AddressService/server/helper"
	"database/sql"
	"encoding/json"
	"fmt"
)

type AddressServiceHandler struct {
}

func NewAddressServiceHandler() *AddressServiceHandler {
	return &AddressServiceHandler{}
}

func (p *AddressServiceHandler) GetAllAddress(req *addrservice.Request) (res *addrservice.Response, err error) {
	fmt.Print("Resquest:", req.GetId(), "\n")

	db, err := sql.Open("mysql", "password@tcp(host:port)/schema")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	fmt.Println("Open success")

	rows, err := db.Query("SELECT * FROM xm_region")
	if err != nil {
		panic(err.Error())
	}

	result, err := helper.ProcessResult(rows)
	if err != nil {
		panic(err.Error())
	}

	// fmt.Println(result)

	var b []byte
	b, err = json.Marshal(result)
	if err != nil {
		panic(err.Error())
	}

	res = addrservice.NewResponse()
	res.Code = 200
	res.Desc = "OK"
	res.Data = string(b)

	fmt.Println("Call finishied")
	fmt.Println("----------------------------")

	return res, err
}
