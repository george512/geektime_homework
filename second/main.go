package main

import (
	"fmt"
	"handling-error/dao"
	"log"
)

func main() {
	defer func() {
		dao.DB.Close()
	}()

	user, err := dao.Get(2)
	if err != nil && dao.IsNoRows(err) {
		// handling no rows
		log.Printf("now rows error")
		return
	} else if err != nil {
		// handling error
		log.Printf("%+v", err)
		return
	}

	fmt.Println(user)
}
