package models

type Contrat struct {
	NumeroContrat string `json:"numero_contrat" pg:"numero_contrat,pk"`
	NumeroAdherent string `json:"numero_adherent" pg:"numero_adherent,unique,notnull" validate:"required"`
}

