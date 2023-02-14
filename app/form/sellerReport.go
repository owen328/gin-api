package form

type UploadSellerReportForm struct {
	StartDate  string `form:"start_date"`
	PeriodType int8   `form:"period_type"`
}
