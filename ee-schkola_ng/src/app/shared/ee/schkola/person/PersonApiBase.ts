import {PersonName} from "../SharedApiBase"


export class Church {
    name: string
    address: Address
    pastor: PersonName
    contact: Contact
    id: string

}


export class Graduation {
    name: string
    level: GraduationLevel
    id: string

}


export class Profile {
    gender: Gender
    name: PersonName
    birthName: string
    birthday: Date
    address: Address
    contact: Contact
    photoData: Blob
    photo: string
    family: Family
    church: ChurchInfo
    education: Education
    id: string


    findByEmail(email: string = "") : Profile {
        throw new ReferenceError("Not implemented yet.");
    }

    findByPhone(phone: string = "") : Profile {
        throw new ReferenceError("Not implemented yet.");
    }

}






export class Address {
    street: string
    suite: string
    city: string
    code: string
    country: string

}


export class ChurchInfo {
    church: string
    association: string
    member: boolean
    services: string

}


export class Contact {
    phone: string
    email: string
    cellphone: string

}


export class Education {
    graduation: Graduation
    profession: string

}


export class Family {
    maritalState: MaritalState
    childrenCount: number
    partner: PersonName

}




enum Gender {
    UNKNOWN,
    MALE,
    FEMALE
}


enum GraduationLevel {
    UNKNOWN,
    MIDDLE_SCHOOL,
    SECONDARY_SCHOOL,
    HIGH_SCHOOL,
    TECHNICAL_COLLEGE,
    COLLEGE
}


enum MaritalState {
    UNKNOWN,
    SINGLE,
    MARRIED,
    SEPARATED,
    DIVORCED,
    WIDOWED
}



