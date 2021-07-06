package data

type Announcement struct {
	Accommodation Accommodation `json:"accommodation"`
	Location      Location      `json:"location"`
	Options       Options       `json:"options"`
	CreatedOn     string        `json:"created_on"`
	UpdatedOn     string        `json:"updated_on"`
	DeletedOn     string        `json:"deleted_on"`
}
