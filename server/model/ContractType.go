package model

type ContractTypeName int

const (
	PLAYER_CONTRACT ContractTypeName = iota
	MANAGER_CONTRACT
	COACH_CONTRACT
	SCOUT_CONTRACT
)

type ContractType struct {
	ID   int64
	Name ContractTypeName
}

type ContractTypeDAO struct {
	IdTipUgo  int64
	NazTipUgo ContractTypeName
}

type ContractTypeResponseDTO struct {
	IdTipUgo  int64            `json:"idTipUgo"`
	NazTipUgo ContractTypeName `json:"nazTipUgo"`
}
