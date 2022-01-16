package main

import (
	"context"
	"fmt"
	"log"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

func main() {

	
	packet := radius.New(radius.CodeAccessRequest, []byte(`123456`))
	rfc2865.UserName_SetString(packet, "user3")

	rfc2865.UserPassword_SetString(packet, "Thteam47@")

	//res, er := radius.
	//radius.
	response, err := radius.Exchange(context.Background(), packet, "192.168.3.128:1812")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.Code)
}
