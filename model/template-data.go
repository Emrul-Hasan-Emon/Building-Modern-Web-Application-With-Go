package model

type TemplateData struct {
	StringMap      map[string]string
	IntMap         map[string]int
	FloatMap       map[string]float32
	Data           map[string]interface{}
	CSRFToken      string
	FlashMessage   string
	WarningMessage string
	ErrorMessage   string
}
