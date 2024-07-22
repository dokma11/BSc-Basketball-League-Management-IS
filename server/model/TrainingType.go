package model

type TrainingType struct {
	IdTipTrng   int64  `json:"idTipTrng"`
	NazTipTrng  string `json:"nazTipTrng"`  // Name
	CiljTipTrng string `json:"ciljTipTrng"` // Focus
}
