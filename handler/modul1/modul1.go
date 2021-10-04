package modul1

import (
	ctrl "simple-fasthttp/controller/modul1"
	mdl "simple-fasthttp/models/modul1"

	"github.com/gofiber/fiber/v2"
)

const modul = "modul1"

func Get(c *fiber.Ctx) error {
	var err error
	c.Send([]byte("Hello World"))
	return err
}

func DataHandler(c *fiber.Ctx) error {
	var param mdl.Request
	var err error
	if err := c.BodyParser(param); err != nil {
		if param.Id >= 0 {
			result, err := ctrl.GetData(param)
			if err == nil {
				result.Code = 200
				result.Message = "success retrieve data"
				c.JSON(result)
			} else {
				result.Code = 500
				result.Message = "internal server error"
				c.JSON(result)
			}
		}
		// return err
	}
	return err
	// else {
	// 	c.JSON(map[string]{"code": 400, "message": fmt.Sprintf("kode %s must be a string", modul)})
	// }
	// return err
}

func CreateHandler(c *fiber.Ctx) error {
	var param mdl.Request
	var err error
	if err := c.BodyParser(param); err != nil {
		return err
	}
	return err
	// if _, err := strconv.Atoi(param.Nama); err != nil {
	// 	err = ctrl.CreateData(param)
	// 	if err == nil {
	// 		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success create data"})
	// 	} else {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
	// 	}
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "this account doesn't have access to create data"})
	// }
}

func UpdateHandler(c *fiber.Ctx) error {
	var param mdl.Request
	var err error
	if err := c.BodyParser(param); err != nil {
		return err
	}
	return err
	// if param.Id >= 0 {
	// 	err := ctrl.UpdateData(param)
	// 	if err == nil {
	// 		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success update data"})
	// 	} else {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
	// 	}
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "this account doesn't have access to update data"})
	// }
}

func DeleteHandler(c *fiber.Ctx) error {
	var param mdl.Request
	var err error
	if err := c.BodyParser(param); err != nil {
		return err
	}
	return err
	// if param.Id >= 0 {
	// 	err := ctrl.DeleteData(param)
	// 	if err == nil {
	// 		c.JSON(http.StatusOK, gin.H{"code": 200, "message": "success delete data"})
	// 	} else {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": err.Error()})
	// 	}
	// } else {
	// 	c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "this account doesn't have access to delete data"})
	// }
}
