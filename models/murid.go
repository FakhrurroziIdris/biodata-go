package models

type Murid struct {
	ID        int16  `json:"ID"`
	Nama      string `json:"Nama"`
	Alamat    string `json:"Alamat"`
	Pekerjaan string `json:"Pekerjaan"`
	Alasan    string `json:"Alasan"`
}

type Murid_List struct {
	Murid_List []Murid `json:"Murid_List"`
}
