package library

import (
    "ee/schkola/shared"
    "encoding/json"
    "fmt"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "gopkg.in/mgo.v2/bson"
    "time"
)
const (
     CreateEvent eventhorizon.EventType = "Create"
     CreatedEvent eventhorizon.EventType = "Created"
     DeleteEvent eventhorizon.EventType = "Delete"
     DeletedEvent eventhorizon.EventType = "Deleted"
     ChangeLocationEvent eventhorizon.EventType = "ChangeLocation"
     UpdateEvent eventhorizon.EventType = "Update"
     UpdatedEvent eventhorizon.EventType = "Updated"
     ChangeLocationedEvent eventhorizon.EventType = "ChangeLocationed"
)




type Create struct {
    Location *Location `json:"location" eh:"optional"`
    Title string `json:"title" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Language string `json:"language" eh:"optional"`
    ReleaseDate *time.Time `json:"releaseDate" eh:"optional"`
    Edition string `json:"edition" eh:"optional"`
    Category string `json:"category" eh:"optional"`
    Author *shared.PersonName `json:"author" eh:"optional"`
    Location *Location `json:"location" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Created struct {
    Location *Location `json:"location" eh:"optional"`
    Title string `json:"title" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Language string `json:"language" eh:"optional"`
    ReleaseDate *time.Time `json:"releaseDate" eh:"optional"`
    Edition string `json:"edition" eh:"optional"`
    Category string `json:"category" eh:"optional"`
    Author *shared.PersonName `json:"author" eh:"optional"`
    Location *Location `json:"location" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Delete struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Deleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type ChangeLocation struct {
    Location *Location `json:"location" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Update struct {
    Location *Location `json:"location" eh:"optional"`
    Title string `json:"title" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Language string `json:"language" eh:"optional"`
    ReleaseDate *time.Time `json:"releaseDate" eh:"optional"`
    Edition string `json:"edition" eh:"optional"`
    Category string `json:"category" eh:"optional"`
    Author *shared.PersonName `json:"author" eh:"optional"`
    Location *Location `json:"location" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type Updated struct {
    Location *Location `json:"location" eh:"optional"`
    Title string `json:"title" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Language string `json:"language" eh:"optional"`
    ReleaseDate *time.Time `json:"releaseDate" eh:"optional"`
    Edition string `json:"edition" eh:"optional"`
    Category string `json:"category" eh:"optional"`
    Author *shared.PersonName `json:"author" eh:"optional"`
    Location *Location `json:"location" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type ChangeLocationed struct {
    Location *Location `json:"location" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}




type BookEventType struct {
	name  string
	ordinal int
}

func (o *BookEventType) Name() string {
    return o.name
}

func (o *BookEventType) Ordinal() int {
    return o.ordinal
}

func (o BookEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *BookEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := BookEventTypes().ParseBookEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid BookEventType %q", lit.Name)
        }
	}
	return
}

func (o BookEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *BookEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := BookEventTypes().ParseBookEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid BookEventType %q", lit)
        }
    }
    return
}

func (o *BookEventType) IsCreate() bool {
    return o == _bookEventTypes.Create()
}

func (o *BookEventType) IsCreated() bool {
    return o == _bookEventTypes.Created()
}

func (o *BookEventType) IsDelete() bool {
    return o == _bookEventTypes.Delete()
}

func (o *BookEventType) IsDeleted() bool {
    return o == _bookEventTypes.Deleted()
}

func (o *BookEventType) IsChangeLocation() bool {
    return o == _bookEventTypes.ChangeLocation()
}

func (o *BookEventType) IsUpdate() bool {
    return o == _bookEventTypes.Update()
}

func (o *BookEventType) IsUpdated() bool {
    return o == _bookEventTypes.Updated()
}

func (o *BookEventType) IsChangeLocationed() bool {
    return o == _bookEventTypes.ChangeLocationed()
}

type bookEventTypes struct {
	values []*BookEventType
    literals []enum.Literal
}

var _bookEventTypes = &bookEventTypes{values: []*BookEventType{
    {name: "Create", ordinal: 0},
    {name: "Created", ordinal: 1},
    {name: "Delete", ordinal: 2},
    {name: "Deleted", ordinal: 3},
    {name: "ChangeLocation", ordinal: 4},
    {name: "Update", ordinal: 5},
    {name: "Updated", ordinal: 6},
    {name: "ChangeLocationed", ordinal: 7}},
}

func BookEventTypes() *bookEventTypes {
	return _bookEventTypes
}

func (o *bookEventTypes) Values() []*BookEventType {
	return o.values
}

func (o *bookEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *bookEventTypes) Create() *BookEventType {
    return _bookEventTypes.values[0]
}

func (o *bookEventTypes) Created() *BookEventType {
    return _bookEventTypes.values[1]
}

func (o *bookEventTypes) Delete() *BookEventType {
    return _bookEventTypes.values[2]
}

func (o *bookEventTypes) Deleted() *BookEventType {
    return _bookEventTypes.values[3]
}

func (o *bookEventTypes) ChangeLocation() *BookEventType {
    return _bookEventTypes.values[4]
}

func (o *bookEventTypes) Update() *BookEventType {
    return _bookEventTypes.values[5]
}

func (o *bookEventTypes) Updated() *BookEventType {
    return _bookEventTypes.values[6]
}

func (o *bookEventTypes) ChangeLocationed() *BookEventType {
    return _bookEventTypes.values[7]
}

func (o *bookEventTypes) ParseBookEventType(name string) (ret *BookEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*BookEventType), ok
	}
	return
}



