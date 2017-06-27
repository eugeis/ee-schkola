package ee.schkola.Person

import ee.schkola.SchkolaBase


enum class Gender {
    UNKNOWN,
    MALE,
    FEMALE;

    fun isUnknown() : Boolean = this == UNKNOWN
    fun isMale() : Boolean = this == MALE
    fun isFemale() : Boolean = this == FEMALE
}


fun String?.toGender(): Gender {
    return if (this != null) Gender.valueOf(this) else Gender.UNKNOWN
}


enum class GraduationType {
    UNKNOWN,
    MIDDLE_SCHOOL,
    SECONDARY_SCHOOL,
    HIGH_SCHOOL,
    TECHNICAL_COLLEGE,
    COLLEGE;

    fun isUnknown() : Boolean = this == UNKNOWN
    fun isMiddleSchool() : Boolean = this == MIDDLE_SCHOOL
    fun isSecondarySchool() : Boolean = this == SECONDARY_SCHOOL
    fun isHighSchool() : Boolean = this == HIGH_SCHOOL
    fun isTechnicalCollege() : Boolean = this == TECHNICAL_COLLEGE
    fun isCollege() : Boolean = this == COLLEGE
}


fun String?.toGraduationType(): GraduationType {
    return if (this != null) GraduationType.valueOf(this) else GraduationType.UNKNOWN
}


enum class MaritalStatus {
    UNKNOWN,
    SINGLE,
    MARRIED,
    SEPARATED,
    DIVORCED,
    WIDOWED;

    fun isUnknown() : Boolean = this == UNKNOWN
    fun isSingle() : Boolean = this == SINGLE
    fun isMarried() : Boolean = this == MARRIED
    fun isSeparated() : Boolean = this == SEPARATED
    fun isDivorced() : Boolean = this == DIVORCED
    fun isWidowed() : Boolean = this == WIDOWED
}


fun String?.toMaritalStatus(): MaritalStatus {
    return if (this != null) MaritalStatus.valueOf(this) else MaritalStatus.UNKNOWN
}




open class Address {
    val street: String
    val suite: String
    val city: String
    val code: String
    val country: String


    constructor(street: String = "", suite: String = "", city: String = "", code: String = "", country: String = "") {
        this.street = street
        this.suite = suite
        this.city = city
        this.code = code
        this.country = country
    }

    companion object {
        val EMPTY = Address()
    }
}


open class Church : SchkolaBase {
    val name: String
    val address: Address
    val pastor: PersonName
    val contact: Contact


    constructor(name: String = "", address: Address = Address.EMPTY, pastor: PersonName = PersonName.EMPTY, 
                contact: Contact = Contact.EMPTY) {
        this.name = name
        this.address = address
        this.pastor = pastor
        this.contact = contact
    }

    companion object {
        val EMPTY = Church()
    }
}


open class ChurchInfo {
    val name: String
    val member: Boolean
    val services: MutableList<String>
    val church: String


    constructor(name: String = "", member: Boolean = false, services: MutableList<String> = arrayListOf(), church: String = "") {
        this.name = name
        this.member = member
        this.services = services
        this.church = church
    }

    companion object {
        val EMPTY = ChurchInfo()
    }
}


open class Contact {
    val phone: String
    val email: String
    val cellphone: String


    constructor(phone: String = "", email: String = "", cellphone: String = "") {
        this.phone = phone
        this.email = email
        this.cellphone = cellphone
    }

    companion object {
        val EMPTY = Contact()
    }
}


open class Education {
    val graduation: Graduation
    val profession: String


    constructor(graduation: Graduation = Graduation.EMPTY, profession: String = "") {
        this.graduation = graduation
        this.profession = profession
    }

    companion object {
        val EMPTY = Education()
    }
}


open class Family {
    val maritalStatus: MaritalStatus
    val childrenCount: Int
    val partner: PersonName


    constructor(maritalStatus: MaritalStatus = MaritalStatus.EMPTY, childrenCount: Int = 0, partner: PersonName = PersonName.EMPTY) {
        this.maritalStatus = maritalStatus
        this.childrenCount = childrenCount
        this.partner = partner
    }

    companion object {
        val EMPTY = Family()
    }
}


open class Graduation : SchkolaBase {
    val name: String
    val type: GraduationType


    constructor(name: String = "", type: GraduationType = GraduationType.EMPTY) {
        this.name = name
        this.type = type
    }

    companion object {
        val EMPTY = Graduation()
    }
}


open class PersonName {
    val first: String
    val last: String


    constructor(first: String = "", last: String = "") {
        this.first = first
        this.last = last
    }

    companion object {
        val EMPTY = PersonName()
    }
}


open class User : SchkolaBase {
    val gender: Gender
    val name: PersonName
    val birthName: String
    val birthday: 
    val address: Address
    val contact: Contact
    val photo: String
    val family: Family
    val church: ChurchInfo
    val education: Education


    constructor(gender: Gender = Gender.EMPTY, name: PersonName = PersonName.EMPTY, birthName: String = "", birthday:  = (), 
                address: Address = Address.EMPTY, contact: Contact = Contact.EMPTY, photo: String = "", 
                family: Family = Family.EMPTY, church: ChurchInfo = ChurchInfo.EMPTY, 
                education: Education = Education.EMPTY) {
        this.gender = gender
        this.name = name
        this.birthName = birthName
        this.birthday = birthday
        this.address = address
        this.contact = contact
        this.photo = photo
        this.family = family
        this.church = church
        this.education = education
    }

    companion object {
        val EMPTY = User()
    }
}



