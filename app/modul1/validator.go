package modul1

var (
	getRule = map[string]string{
		"Id": "min=0",
	}

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
