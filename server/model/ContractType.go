package model

type ContractTypeName int

const (
	PLAYER_CONTRACT ContractTypeName = iota
	MANAGER_CONTRACT
	COACH_CONTRACT
	SCOUT_CONTRACT
)

type ContractType struct {
	IdTipUgo  int64            `json:"idTipUgo"`
	NazTipUgo ContractTypeName `json:"nazTipUgo"`
}
