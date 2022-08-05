package service

import (
	"biodata.go/models"
	"biodata.go/repository"
)

func GetMurid(kode int16) (data models.Murid) {
	murid_list := repository.GetAllData("./murid_list.json").Murid_List

	for _, v := range murid_list {
		if v.ID == kode {
			return v
		}
	}
	return models.Murid{}
}
