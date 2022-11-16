package posapi

type checkAPIRes struct {
	Success  bool `json:"success"`
	Database struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"database"`
	Config struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"config"`
	Network struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
	} `json:"network"`
}

type InformationRes struct {
	RegisterNo string `json:"registerNo"`
	BranchNo   string `json:"branchNo"`
	PosID      string `json:"posId"`
	DBDirPath  string `json:"dbDirPath"`
	ExtraInfo  struct {
		CountBill string `json:"countBill"`
	} `json:"extraInfo"`
}
type PutParams struct {
}
