package fooddtos

type GetFoodDTO struct {
	Id int `json:"id" uri:"id" binding:"required"`
}

