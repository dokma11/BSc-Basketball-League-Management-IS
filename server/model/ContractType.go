package model

type ContractTypeName int

const (
	PLAYER_CONTRACT ContractTypeName = iota
	MANAGER_CONTRACT
	COACH_CONTRACT
	SCOUT_CONTRACT
)

type ContractType struct {
	ID   int64            `json:"id"`
	Name ContractTypeName `json:"name"`
}

type ContractTypeDAO struct {
	IdTipUgo  int64
	NazTipUgo ContractTypeName
}
