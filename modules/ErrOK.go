package modules

import (
	"fmt"
	"log"
)

func Critical(err error) {
	if err != nil {
		fmt.Println("		치명적인 오류 발생: ", err)
		log.Fatal(err)
	}
}
func ErrOK(err error) error {
	if err != nil {
		fmt.Println("		ErrOK: ", err)
		return err
	} else {
		return nil
	}
}
