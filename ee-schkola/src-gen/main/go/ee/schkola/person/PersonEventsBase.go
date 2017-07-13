package person

import (
    "ee/schkola"
    "time"
)
type ChurchCreated struct {
    Name  
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
    *schkola.SchkolaBase
}


type ChurchDeleted struct {
    Id  
}


type ChurchUpdated struct {
    Name  
    Address  *Address
    Pastor  *PersonName
    Contact  *Contact
    *schkola.SchkolaBase
}


type GraduationCreated struct {
    Name  
    Level  *GraduationLevel
    *schkola.SchkolaBase
}


type GraduationDeleted struct {
    Id  
}


type GraduationUpdated struct {
    Name  
    Level  *GraduationLevel
    *schkola.SchkolaBase
}


type ProfileCreated struct {
    Gender  *Gender
    Name  *PersonName
    BirthName  
    Birthday  *time.Time
    Address  *Address
    Contact  *Contact
    PhotoData  []byte
    Photo  
    Family  *Family
    Church  *ChurchInfo
    Education  *Education
    *schkola.SchkolaBase
}


type ProfileDeleted struct {
    Id  
}


type ProfileUpdated struct {
    Gender  *Gender
    Name  *PersonName
    BirthName  
    Birthday  *time.Time
    Address  *Address
    Contact  *Contact
    PhotoData  []byte
    Photo  
    Family  *Family
    Church  *ChurchInfo
    Education  *Education
    *schkola.SchkolaBase
}





