package form

type CheckVatForm struct {
	VatNumber string `json:"vat_number"`
}

type CheckMultiVatsForm struct {
	VatNumbers []string `json:"vat_numbers"`
}
