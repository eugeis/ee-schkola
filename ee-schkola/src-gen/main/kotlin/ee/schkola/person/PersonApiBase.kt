package ee.schkola.person

import ee.schkola.SchkolaBase
import java.util.*


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


enum class GraduationLevel {
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


fun String?.toGraduationLevel(): GraduationLevel {
    return if (this != null) GraduationLevel.valueOf(this) else GraduationLevel.UNKNOWN
}


enum class MaritalState {
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


fun String?.toMaritalState(): MaritalState {
    return if (this != null) MaritalState.valueOf(this) else MaritalState.UNKNOWN
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


    constructor(id: String = "", name: String = "", address: Address = Address(), pastor: PersonName = PersonName(), 
                contact: Contact = Contact()) : super(id) {
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
    val church: String
    val association: String
    val member: Boolean
    val services: MutableList<String>


    constructor(church: String = "", association: String = "", member: Boolean = false, services: MutableList<String> = arrayListOf()) {
        this.church = church
        this.association = association
        this.member = member
        this.services = services
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


    constructor(graduation: Graduation = Graduation(), profession: String = "") {
        this.graduation = graduation
        this.profession = profession
    }

    companion object {
        val EMPTY = Education()
    }
}


open class Family {
    val maritalState: MaritalState
    val childrenCount: Int
    val partner: PersonName


    constructor(maritalState: MaritalState = MaritalState.UNKNOWN, childrenCount: Int = 0, partner: PersonName = PersonName()) {
        this.maritalState = maritalState
        this.childrenCount = childrenCount
        this.partner = partner
    }

    companion object {
        val EMPTY = Family()
    }
}


open class Graduation : SchkolaBase {
    val name: String
    val level: GraduationLevel


    constructor(id: String = "", name: String = "", level: GraduationLevel = GraduationLevel.UNKNOWN) : super(id) {
        this.name = name
        this.level = level
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


open class Profile : SchkolaBase {
    val gender: Gender
    val name: PersonName
    val birthName: String
    val birthday: Date
    val address: Address
    val contact: Contact
    val photoData: ByteArray
    val photo: String
    val family: Family
    val church: ChurchInfo
    val education: Education


    constructor(id: String = "", gender: Gender = Gender.UNKNOWN, name: PersonName = PersonName(), birthName: String = "", 
                birthday: Date = Date(), address: Address = Address(), contact: Contact = Contact(), 
                photoData: ByteArray = ByteArray(0), photo: String = "", family: Family = Family(), 
                church: ChurchInfo = ChurchInfo(), education: Education = Education()) : super(id) {
        this.gender = gender
        this.name = name
        this.birthName = birthName
        this.birthday = birthday
        this.address = address
        this.contact = contact
        this.photoData = photoData
        this.photo = photo
        this.family = family
        this.church = church
        this.education = education
    }

    companion object {
        val EMPTY = Profile()
    }
}



