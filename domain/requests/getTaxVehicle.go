package requests

type getTaxVehicle struct {
	Vehicle string `json:"vehicle" binding:"required,email"`
	dates   string `json:"password" binding:"required,min=6,containsany=!@*"`
	Name    string `json:"name" binding:"required,min=4,max=50"`
	Age     int8   `json:"age" binding:"required,min=2,max=140"`
}
