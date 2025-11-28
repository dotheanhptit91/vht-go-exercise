package fooddtos

type DeleteFoodDTO struct {
	Id int `json:"id" uri:"id" binding:"required"`
}

