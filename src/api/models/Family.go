package models

type Family struct {
	FamilyId    uint32 `gorm:"primary_key;auto_increment" json:"familyId"`
	Name        string `gorm:"size:50;not null;unique" json:"name"`
	MaxPerson   int    `gorm:"type:integer;not null" json:"maxPerson"`
	CurrentSize int    `gorm:"type:integer;default:'0'" json:"currentSize"`
}

func (f *Family) AddPerson(idFamily uint32) {
	f.CurrentSize = f.CurrentSize + 1
}
