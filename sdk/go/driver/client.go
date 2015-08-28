package driver

import (
	"AddressService/gen/gen-go/addrservice"
	"fmt"
)

func HandleClient(client *addrservice.AddressServiceClient) (err error) {

	req := addrservice.NewRequest()
	req.Id = "123"

	res, err := client.GetAllAddress(req)

	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	fmt.Println(res.GetCode(), res.GetDesc(), res.GetData())
	return err

}
