package person

import (
    "ee/schkola"
    "time"
)
type ChurchCreated struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
    *schkola.SchkolaBase
}


type ChurchDeleted struct {
    Id  string
}


type ChurchUpdated struct {
    Name  string
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
    *schkola.SchkolaBase
}


type GraduationCreated struct {
    Name  string
    Level  *GraduationLevel
    *schkola.SchkolaBase
}


type GraduationDeleted struct {
    Id  string
}


type GraduationUpdated struct {
    Name  string
    Level  *GraduationLevel
    *schkola.SchkolaBase
}


type ProfileCreated struct {
    Gender  *Gender
    Name  *PersonName
    BirthName  string
    Birthday  *time.Time
    Address  *Address
    Contact  *Contact
    PhotoData  []byte
    Photo  string
    Family  *Family
    Church  *ChurchInfo
    Education  *Education
    *schkola.SchkolaBase
}


type ProfileDeleted struct {
    Id  string
}


type ProfileUpdated struct {
    Gender  *Gender
    Name  *PersonName
    BirthName  string
    Birthday  *time.Time
    Address  *Address
    Contact  *Contact
    PhotoData  []byte
    Photo  string
    Family  *Family
    Church  *ChurchInfo
    Education  *Education
    *schkola.SchkolaBase
}





