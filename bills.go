package bills

type Bill struct {
	Id          int    `json:"id" db:"id"`
	Amount      int    `json:"amount" db:"amount"`
	Description string `json:"description" db:"description"`
	Email       string `json:"email" db:"email"`
	CompanyId   int    `json:"company_id" db:"company_id"`
}
