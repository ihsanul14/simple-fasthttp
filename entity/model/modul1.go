package modul1

type Request struct {
	Id         int    `json:"id" xml:"id" form:"id"`
	Nama       string `json:"nama" xml:"nama" form:"nama"`
	Nomor      int    `json:"nomor" xml:"nomor" form:"nomor"`
	Created_at string `json:"created_at" xml:"created_at" form:"created_at"`
	Updated_at string `json:"updated_at" xml:"updated_at" form:"upadated_at"`
}

func (c Request) TableName() string {
	return "testing"
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
