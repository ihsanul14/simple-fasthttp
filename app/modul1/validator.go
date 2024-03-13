package modul1

var (
	createRule = map[string]string{
		"Nama":  "required",
		"Nomor": "required",
	}

	updateRule = map[string]string{
		"Id":    "required,min=0",
		"Nama":  "required",
		"Nomor": "required",
	}

	deleteRule = map[string]string{
		"Id": "required,min=0",
	}
)
