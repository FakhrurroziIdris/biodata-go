package main

import (
	"fmt"
	"os"
	"strconv"

	"biodata.go/models"
	"biodata.go/service"
)

func main() {
	if len(os.Args) > 1 {
		kode, _ := strconv.Atoi(os.Args[1])
		murid := service.GetMurid(int16(kode))

		if (models.Murid{}) == murid {
			fmt.Println("Tidak ditemukan data")
		} else {
			fmt.Println(murid)
		}
	} else {
		fmt.Println("No user arguments")
	}
}
