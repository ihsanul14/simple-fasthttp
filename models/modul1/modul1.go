package modul1

type Request struct {
	Id         int    `json:"id"`
	Nama       string `json:"nama"`
	Nomor      int    `json:"nomor"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type Response struct {
	Id         int    `json:"id" gorm:"id"`
	Nama       string `json:"nama" gorm:"nama"`
	Nomor      int    `json:"nomor" gorm:"nomor"`
	Created_at string `json:"created_at" gorm:"created_at"`
	Updated_at string `json:"updated_at" gorm:"updated_at"`
}

type ResponseAll struct {
	Code    int        `json:"code"`
	Message string     `json:"message"`
	Data    []Response `json:"data"`
}
