package repository

import (
	"encoding/json"
	"io/ioutil"

	"biodata.go/models"
)

func GetAllData(path string) (data models.Murid_List) {
	file, _ := ioutil.ReadFile(path)

	_ = json.Unmarshal([]byte(file), &data)

	// for i := 0; i < len(data.Murid_List); i++ {
	// 	fmt.Println("ID: ", data.Murid_List[i].ID)
	// 	fmt.Println("Nama: ", data.Murid_List[i].Nama)
	// 	fmt.Println("Alasan: ", data.Murid_List[i].Alamat)
	// 	fmt.Println("Pekerjaan: ", data.Murid_List[i].Pekerjaan)
	// 	fmt.Println("Alasan: ", data.Murid_List[i].Alasan)
	// }

	return
}
