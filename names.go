package ComApart

type SearchCounter struct {
	Gas         int `json:"gas" db:"gas"`
	Water       int `json:"water" db:"water"`
	Electricity int `json:"electricity" db:"electricity"`
}

type SearchPrice struct {
	CostGas          float64 `json:"cost_gas" db:"cost_gas"`
	CostElectricity  float64 `json:"cost_electricity" db:"cost_electricity"`
	CostWater        float64 `json:"cost_water" db:"cost_water"`
	Intercom         int `json:"intercom" db:"intercom"`
	Repair           int `json:"repair" db:"repair"`
	CommunalServices int `json:"communal_services" db:"communal_services"`
}