export class Church {
    name: string
    address: Address
    pastor: PersonName
    contact: Contact
    association: string
    id: string
}

export class Graduation {
    name: string
    level: GraduationLevel
    id: string
}

export class Profile {
    gender: Gender = 0;
    name: PersonName = {first: '', last: ''};
    birthName = ''
    birthday: Date = new Date();
    address: Address = {street: '', suite: '', city: '', code: '', country: ''};
    contact: Contact
    photoData: Blob
    photo: string
    family: Family
    church: ChurchInfo
    education: Education
    id: string

    findByEmail(email: string = ''): Profile {
        throw new ReferenceError('Not implemented yet.');
    }

    findByPhone(phone: string = ''): Profile {
        throw new ReferenceError('Not implemented yet.');
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
    other: string
    profession: string
}

export class Family {
    maritalState: MaritalState
    childrenCount: number
    partner: PersonName
}

export class PersonName {
    first: string
    last: string
}



export enum Gender {
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

export enum MaritalState {
    UNKNOWN,
    SINGLE,
    MARRIED,
    SEPARATED,
    DIVORCED,
    WIDOWED
}



