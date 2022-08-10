export class Church {
    name = ''
    address: Address = new Address()
    pastor: PersonName = new PersonName()
    contact: Contact = new Contact()
    association = ''
    id = ''
}

export class Graduation {
    name = ''
    level: GraduationLevel = GraduationLevel.UNKNOWN
    id = ''
}

export class Profile {
    gender: Gender = Gender.UNKNOWN
    name: PersonName = new PersonName()
    birthName = ''
    birthday: Date = new Date()
    address: Address = new Address()
    contact: Contact = new Contact()
    photoData: Blob /*= new ByteArray(0)*/
    photo = ''
    family: Family = new Family()
    church: ChurchInfo = new ChurchInfo()
    education: Education = new Education()
    id = ''

    findByEmail(email = ''): Profile {
        throw new ReferenceError('Not implemented yet.');
    }

    findByPhone(phone = ''): Profile {
        throw new ReferenceError('Not implemented yet.');
    }
}





export class Address {
    street = ''
    suite = ''
    city = ''
    code = ''
    country = ''
}

export class ChurchInfo {
    church = ''
    member = false
    services = ''
}

export class Contact {
    phone = ''
    email = ''
    cellphone = ''
}

export class Education {
    graduation: Graduation = new Graduation()
    other = ''
    profession = ''
}

export class Family {
    maritalState: MaritalState = MaritalState.UNKNOWN
    childrenCount = 0
    partner: PersonName = new PersonName()
}

export class PersonName {
    first = ''
    last = ''
}



export enum Gender {
    UNKNOWN,
    MALE,
    FEMALE
}

export enum GraduationLevel {
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



