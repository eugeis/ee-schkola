package ee.schkola.student

import ee.schkola.SchkolaBase
import ee.schkola.Trace
import ee.schkola.person.Contact
import ee.schkola.person.PersonName
import ee.schkola.person.User
import java.util.*

enum class GroupType() {
    COURSE_GROUP(),
    YEAR_GROUP();


    fun isCourseGroup() : Boolean = this == COURSE_GROUP
    fun isYearGroup() : Boolean = this == YEAR_GROUP
}

fun String?.toGroupType(): GroupType {
    return if (this != null) {
        GroupType.valueOf(this)
    } else {
        GroupType.COURSE_GROUP
    }
}

enum class AttendanceStatus() {
    REGISTERED(),
    CONFIRMED(),
    CANCELED(),
    PRESENT();


    fun isRegistered() : Boolean = this == REGISTERED
    fun isConfirmed() : Boolean = this == CONFIRMED
    fun isCanceled() : Boolean = this == CANCELED
    fun isPresent() : Boolean = this == PRESENT
}

fun String?.toAttendanceStatus(): AttendanceStatus {
    return if (this != null) {
        AttendanceStatus.valueOf(this)
    } else {
        AttendanceStatus.REGISTERED
    }
}

open class SchoolApplication : SchkolaBase {
    companion object {
        val EMPTY = SchoolApplication("", User.EMPTY, false, PersonName.EMPTY, PersonName.EMPTY, Contact.EMPTY, SchoolYear.EMPTY, "")
    }
    var person: User = User.EMPTY
    var hasRecommendation: Boolean = false
    var recommendationOf: PersonName = PersonName.EMPTY
    var mentor: PersonName = PersonName.EMPTY
    var mentorContact: Contact = Contact.EMPTY
    var schoolYear: SchoolYear = SchoolYear.EMPTY
    var group: String = ""

    constructor(id: String = "", person: User = User.EMPTY, hasRecommendation: Boolean = false, 
        recommendationOf: PersonName = PersonName.EMPTY, mentor: PersonName = PersonName.EMPTY, 
        mentorContact: Contact = Contact.EMPTY, schoolYear: SchoolYear = SchoolYear.EMPTY, group: String = ""): super(id) {
        this.person = person
        this.hasRecommendation = hasRecommendation
        this.recommendationOf = recommendationOf
        this.mentor = mentor
        this.mentorContact = mentorContact
        this.schoolYear = schoolYear
        this.group = group
    }


}

fun SchoolApplication?.orEmpty(): SchoolApplication {
    return if (this != null) this else SchoolApplication.EMPTY
}

open class SchoolYear : SchkolaBase {
    companion object {
        val EMPTY = SchoolYear("", "", Date(), Date(), arrayListOf())
    }
    var name: String = ""
    var start: Date = Date()
    var end: Date = Date()
    var dates: MutableList<Date> = arrayListOf()

    constructor(id: String = "", name: String = "", start: Date = Date(), end: Date = Date(), dates: MutableList<Date> = arrayListOf()): super(id) {
        this.name = name
        this.start = start
        this.end = end
        this.dates = dates
    }


}

fun SchoolYear?.orEmpty(): SchoolYear {
    return if (this != null) this else SchoolYear.EMPTY
}

open class Course : SchkolaBase {
    companion object {
        val EMPTY = Course("", "", Date(), Date(), PersonName.EMPTY, SchoolYear.EMPTY, 0f, "")
    }
    var name: String = ""
    var begin: Date = Date()
    var end: Date = Date()
    var teacher: PersonName = PersonName.EMPTY
    var schoolYear: SchoolYear = SchoolYear.EMPTY
    var fee: Float = 0f
    var description: String = ""

    constructor(id: String = "", name: String = "", begin: Date = Date(), end: Date = Date(), teacher: PersonName = PersonName.EMPTY, 
        schoolYear: SchoolYear = SchoolYear.EMPTY, fee: Float = 0f, description: String = ""): super(id) {
        this.name = name
        this.begin = begin
        this.end = end
        this.teacher = teacher
        this.schoolYear = schoolYear
        this.fee = fee
        this.description = description
    }


}

fun Course?.orEmpty(): Course {
    return if (this != null) this else Course.EMPTY
}

open class Group : SchkolaBase {
    companion object {
        val EMPTY = Group("", "", GroupType.COURSE_GROUP, SchoolYear.EMPTY, User.EMPTY, arrayListOf(), arrayListOf())
    }
    var name: String = ""
    var type: GroupType = GroupType.COURSE_GROUP
    var schoolYear: SchoolYear = SchoolYear.EMPTY
    var representative: User = User.EMPTY
    var students: MutableList<User> = arrayListOf()
    var courses: MutableList<Course> = arrayListOf()

    constructor(id: String = "", name: String = "", type: GroupType = GroupType.COURSE_GROUP, schoolYear: SchoolYear = SchoolYear.EMPTY, 
        representative: User = User.EMPTY, students: MutableList<User> = arrayListOf(), 
        courses: MutableList<Course> = arrayListOf()): super(id) {
        this.name = name
        this.type = type
        this.schoolYear = schoolYear
        this.representative = representative
        this.students = students
        this.courses = courses
    }


}

fun Group?.orEmpty(): Group {
    return if (this != null) this else Group.EMPTY
}

open class Grade : SchkolaBase {
    companion object {
        val EMPTY = Grade("", User.EMPTY, Course.EMPTY, 0f, Trace.EMPTY, "")
    }
    var student: User = User.EMPTY
    var course: Course = Course.EMPTY
    var grade: Float = 0f
    var gradeTrace: Trace = Trace.EMPTY
    var comment: String = ""

    constructor(id: String = "", student: User = User.EMPTY, course: Course = Course.EMPTY, grade: Float = 0f, 
        gradeTrace: Trace = Trace.EMPTY, comment: String = ""): super(id) {
        this.student = student
        this.course = course
        this.grade = grade
        this.gradeTrace = gradeTrace
        this.comment = comment
    }


}

fun Grade?.orEmpty(): Grade {
    return if (this != null) this else Grade.EMPTY
}

open class Attendance : SchkolaBase {
    companion object {
        val EMPTY = Attendance("", User.EMPTY, Date(), Course.EMPTY, 0, AttendanceStatus.REGISTERED, Trace.EMPTY, "", Trace.EMPTY)
    }
    var student: User = User.EMPTY
    var date: Date = Date()
    var course: Course = Course.EMPTY
    var hours: Int = 0
    var status: AttendanceStatus = AttendanceStatus.REGISTERED
    var statusTrace: Trace = Trace.EMPTY
    var token: String = ""
    var tokenTrace: Trace = Trace.EMPTY

    constructor(id: String = "", student: User = User.EMPTY, date: Date = Date(), course: Course = Course.EMPTY, hours: Int = 0, 
        status: AttendanceStatus = AttendanceStatus.REGISTERED, statusTrace: Trace = Trace.EMPTY, token: String = "", 
        tokenTrace: Trace = Trace.EMPTY): super(id) {
        this.student = student
        this.date = date
        this.course = course
        this.hours = hours
        this.status = status
        this.statusTrace = statusTrace
        this.token = token
        this.tokenTrace = tokenTrace
    }


}

fun Attendance?.orEmpty(): Attendance {
    return if (this != null) this else Attendance.EMPTY
}


