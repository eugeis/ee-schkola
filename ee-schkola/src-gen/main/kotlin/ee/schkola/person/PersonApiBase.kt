package ee.schkola.person

import ee.schkola.SchkolaBase
import java.util.*

enum class Gender() {
    UNKNOWN(),
    MALE(),
    FEMALE();


    fun isUnknown() : Boolean = this == UNKNOWN
    fun isMale() : Boolean = this == MALE
    fun isFemale() : Boolean = this == FEMALE
}

fun String?.toGender(): Gender {
    return if (this != null) {
        Gender.valueOf(this)
    } else {
        Gender.UNKNOWN
    }
}

enum class MaritalStatus() {
    UNKNOWN(),
    SINGLE(),
    MARRIED(),
    SEPARATED(),
    DIVORCED(),
    WIDOWED();


    fun isUnknown() : Boolean = this == UNKNOWN
    fun isSingle() : Boolean = this == SINGLE
    fun isMarried() : Boolean = this == MARRIED
    fun isSeparated() : Boolean = this == SEPARATED
    fun isDivorced() : Boolean = this == DIVORCED
    fun isWidowed() : Boolean = this == WIDOWED
}

fun String?.toMaritalStatus(): MaritalStatus {
    return if (this != null) {
        MaritalStatus.valueOf(this)
    } else {
        MaritalStatus.UNKNOWN
    }
}

enum class GraduationType() {
    UNKNOWN(),
    MIDDLE_SCHOOL(),
    SECONDARY_SCHOOL(),
    HIGH_SCHOOL(),
    TECHNICAL_COLLEGE(),
    COLLEGE();


    fun isUnknown() : Boolean = this == UNKNOWN
    fun isMiddleSchool() : Boolean = this == MIDDLE_SCHOOL
    fun isSecondarySchool() : Boolean = this == SECONDARY_SCHOOL
    fun isHighSchool() : Boolean = this == HIGH_SCHOOL
    fun isTechnicalCollege() : Boolean = this == TECHNICAL_COLLEGE
    fun isCollege() : Boolean = this == COLLEGE
}

fun String?.toGraduationType(): GraduationType {
    return if (this != null) {
        GraduationType.valueOf(this)
    } else {
        GraduationType.UNKNOWN
    }
}

data class Education(var graduation: Graduation = Graduation.EMPTY, var profession: String = "") {
    companion object {
        val EMPTY = Education(Graduation.EMPTY, "")
    }



}

fun Education?.orEmpty(): Education {
    return if (this != null) this else Education.EMPTY
}

data class ChurchInfo(var name: String = "", var member: Boolean = false, var services: MutableList<String>  = arrayListOf(),
                      var church: String = "") {
    companion object {
        val EMPTY = ChurchInfo("", false, arrayListOf(), "")
    }



}

fun ChurchInfo?.orEmpty(): ChurchInfo {
    return if (this != null) this else ChurchInfo.EMPTY
}

data class Family(var maritalStatus: MaritalStatus = MaritalStatus.UNKNOWN, var childrenCount: Int = 0, 
                  var partner: PersonName = PersonName.EMPTY) {
    companion object {
        val EMPTY = Family(MaritalStatus.UNKNOWN, 0, PersonName.EMPTY)
    }



}

fun Family?.orEmpty(): Family {
    return if (this != null) this else Family.EMPTY
}

data class PersonName(var first: String = "", var last: String = "") {
    companion object {
        val EMPTY = PersonName("", "")
    }



}

fun PersonName?.orEmpty(): PersonName {
    return if (this != null) this else PersonName.EMPTY
}

data class Address(var street: String = "", var suite: String = "", var city: String = "", var code: String = "", var country: String = "") {
    companion object {
        val EMPTY = Address("", "", "", "", "")
    }



}

fun Address?.orEmpty(): Address {
    return if (this != null) this else Address.EMPTY
}

data class Contact(var phone: String = "", var email: String = "", var cellphone: String = "") {
    companion object {
        val EMPTY = Contact("", "", "")
    }



}

fun Contact?.orEmpty(): Contact {
    return if (this != null) this else Contact.EMPTY
}

open class User : SchkolaBase {
    companion object {
        val EMPTY = User("", Gender.UNKNOWN, PersonName.EMPTY, "", Date(), Address.EMPTY, Contact.EMPTY, "", Family.EMPTY, ChurchInfo.EMPTY, 
            Education.EMPTY)
    }
    var gender: Gender = Gender.UNKNOWN
    var name: PersonName = PersonName.EMPTY
    var birthName: String = ""
    var birthday: Date = Date()
    var address: Address = Address.EMPTY
    var contact: Contact = Contact.EMPTY
    var photo: String = ""
    var family: Family = Family.EMPTY
    var church: ChurchInfo = ChurchInfo.EMPTY
    var education: Education = Education.EMPTY

    constructor(id: String = "", gender: Gender = Gender.UNKNOWN, name: PersonName = PersonName.EMPTY, birthName: String = "", 
        birthday: Date = Date(), address: Address = Address.EMPTY, contact: Contact = Contact.EMPTY, photo: String = "", 
        family: Family = Family.EMPTY, church: ChurchInfo = ChurchInfo.EMPTY, education: Education = Education.EMPTY): super(id) {
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


}

fun User?.orEmpty(): User {
    return if (this != null) this else User.EMPTY
}

open class Church : SchkolaBase {
    companion object {
        val EMPTY = Church("", "", Address.EMPTY, PersonName.EMPTY, Contact.EMPTY)
    }
    var name: String = ""
    var address: Address = Address.EMPTY
    var pastor: PersonName = PersonName.EMPTY
    var contact: Contact = Contact.EMPTY

    constructor(id: String = "", name: String = "", address: Address = Address.EMPTY, pastor: PersonName = PersonName.EMPTY, 
        contact: Contact = Contact.EMPTY): super(id) {
        this.name = name
        this.address = address
        this.pastor = pastor
        this.contact = contact
    }


}

fun Church?.orEmpty(): Church {
    return if (this != null) this else Church.EMPTY
}

open class Graduation : SchkolaBase {
    companion object {
        val EMPTY = Graduation("", "", GraduationType.UNKNOWN)
    }
    var name: String = ""
    var type: GraduationType = GraduationType.UNKNOWN

    constructor(id: String = "", name: String = "", type: GraduationType = GraduationType.UNKNOWN): super(id) {
        this.name = name
        this.type = type
    }


}

fun Graduation?.orEmpty(): Graduation {
    return if (this != null) this else Graduation.EMPTY
}


