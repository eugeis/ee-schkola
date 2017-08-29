import {SchkolaBase} from "../SharedApiBase"


export class Church extends SchkolaBase {
    readonly name: string
    readonly address: Address
    readonly pastor: PersonName
    readonly contact: Contact

}


export class Graduation extends SchkolaBase {
    readonly name: string
    readonly level: GraduationLevel

}


export class Profile extends SchkolaBase {
    readonly gender: Gender
    readonly name: PersonName
    readonly birthName: string
    readonly birthday: Date
    readonly address: Address
    readonly contact: Contact
    readonly photoData: Blob
    readonly photo: string
    readonly family: Family
    readonly church: ChurchInfo
    readonly education: Education


    findByEmail(email: string = "") : Profile {
        throw new ReferenceError("Not implemented yet.");
    }

    findByPhone(phone: string = "") : Profile {
        throw new ReferenceError("Not implemented yet.");
    }

}






export class Address {
    readonly street: string
    readonly suite: string
    readonly city: string
    readonly code: string
    readonly country: string

}


export class ChurchInfo {
    readonly church: string
    readonly association: string
    readonly member: boolean
    readonly services: string

}


export class Contact {
    readonly phone: string
    readonly email: string
    readonly cellphone: string

}


export class Education {
    readonly graduation: Graduation
    readonly profession: string

}


export class Family {
    readonly maritalState: MaritalState
    readonly childrenCount: number
    readonly partner: PersonName

}


export class PersonName {
    readonly first: string
    readonly last: string

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



