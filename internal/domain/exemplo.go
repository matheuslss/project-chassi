package domain

import (
	"gopkg.in/guregu/null.v4"
)

type Exemplo struct {
	ID null.Int `db:"id" json:"id"`
}
