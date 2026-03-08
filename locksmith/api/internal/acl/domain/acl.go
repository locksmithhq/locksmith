package domain

type Acl struct {
	Id    string  `json:"id" db:"id" paginate:"id"`
	PType string  `json:"p_type" db:"p_type" paginate:"p_type"`
	V0    *string `json:"v0" db:"v0" paginate:"v0"`
	V1    *string `json:"v1" db:"v1" paginate:"v1"`
	V2    *string `json:"v2" db:"v2" paginate:"v2"`
	V3    *string `json:"v3" db:"v3" paginate:"v3"`
	V4    *string `json:"v4" db:"v4" paginate:"v4"`
	V5    *string `json:"v5" db:"v5" paginate:"v5"`
}
