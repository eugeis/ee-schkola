package ee.schkola.student

import ee.schkola.SchkolaBase
import ee.schkola.Trace
import ee.schkola.person.Contact
import ee.schkola.person.PersonName
import ee.schkola.person.Profile
import java.util.*


enum class AttendanceState {
    REGISTERED,
    CONFIRMED,
    CANCELED,
    PRESENT;

    fun isRegistered() : Boolean = this == REGISTERED
    fun isConfirmed() : Boolean = this == CONFIRMED
    fun isCanceled() : Boolean = this == CANCELED
    fun isPresent() : Boolean = this == PRESENT
}


fun String?.toAttendanceState(): AttendanceState {
    return if (this != null) AttendanceState.valueOf(this) else AttendanceState.REGISTERED
}


enum class GroupCategory {
    COURSE_GROUP,
    YEAR_GROUP;

    fun isCourseGroup() : Boolean = this == COURSE_GROUP
    fun isYearGroup() : Boolean = this == YEAR_GROUP
}


fun String?.toGroupCategory(): GroupCategory {
    return if (this != null) GroupCategory.valueOf(this) else GroupCategory.COURSE_GROUP
}




open class Attendance : SchkolaBase {
    val student: Profile
    val date: Date
    val course: Course
    val hours: Int
    val state: AttendanceState
    val stateTrace: Trace
    val token: String
    val tokenTrace: Trace


    constructor(id: String = "", student: Profile = Profile(), date: Date = Date(), course: Course = Course(), hours: Int = 0, 
                state: AttendanceState = AttendanceState.REGISTERED, stateTrace: Trace = Trace(), token: String = "", 
                tokenTrace: Trace = Trace()) : super(id) {
        this.student = student
        this.date = date
        this.course = course
        this.hours = hours
        this.state = state
        this.stateTrace = stateTrace
        this.token = token
        this.tokenTrace = tokenTrace
    }

    companion object {
        val EMPTY = Attendance()
    }
}


open class Course : SchkolaBase {
    val name: String
    val begin: Date
    val end: Date
    val teacher: PersonName
    val schoolYear: SchoolYear
    val fee: Float
    val description: String


    constructor(id: String = "", name: String = "", begin: Date = Date(), end: Date = Date(), teacher: PersonName = PersonName(), 
                schoolYear: SchoolYear = SchoolYear(), fee: Float = 0f, description: String = "") : super(id) {
        this.name = name
        this.begin = begin
        this.end = end
        this.teacher = teacher
        this.schoolYear = schoolYear
        this.fee = fee
        this.description = description
    }

    companion object {
        val EMPTY = Course()
    }
}


open class Grade : SchkolaBase {
    val student: Profile
    val course: Course
    val grade: Float
    val gradeTrace: Trace
    val comment: String


    constructor(id: String = "", student: Profile = Profile(), course: Course = Course(), grade: Float = 0f, gradeTrace: Trace = Trace(), 
                comment: String = "") : super(id) {
        this.student = student
        this.course = course
        this.grade = grade
        this.gradeTrace = gradeTrace
        this.comment = comment
    }

    companion object {
        val EMPTY = Grade()
    }
}


open class Group : SchkolaBase {
    val name: String
    val category: GroupCategory
    val schoolYear: SchoolYear
    val representative: Profile
    val students: MutableList<Course>
    val courses: MutableList<Course>


    constructor(id: String = "", name: String = "", category: GroupCategory = GroupCategory.COURSE_GROUP, 
                schoolYear: SchoolYear = SchoolYear(), representative: Profile = Profile(), 
                students: MutableList<Course> = arrayListOf(), courses: MutableList<Course> = arrayListOf()) : super(id) {
        this.name = name
        this.category = category
        this.schoolYear = schoolYear
        this.representative = representative
        this.students = students
        this.courses = courses
    }

    companion object {
        val EMPTY = Group()
    }
}


open class SchoolApplication : SchkolaBase {
    val profile: Profile
    val recommendationOf: PersonName
    val churchContactPerson: PersonName
    val churchContact: Contact
    val schoolYear: SchoolYear
    val group: String


    constructor(id: String = "", profile: Profile = Profile(), recommendationOf: PersonName = PersonName(), 
                churchContactPerson: PersonName = PersonName(), churchContact: Contact = Contact(), 
                schoolYear: SchoolYear = SchoolYear(), group: String = "") : super(id) {
        this.profile = profile
        this.recommendationOf = recommendationOf
        this.churchContactPerson = churchContactPerson
        this.churchContact = churchContact
        this.schoolYear = schoolYear
        this.group = group
    }

    companion object {
        val EMPTY = SchoolApplication()
    }
}


open class SchoolYear : SchkolaBase {
    val name: String
    val start: Date
    val end: Date
    val dates: MutableList<Course>


    constructor(id: String = "", name: String = "", start: Date = Date(), end: Date = Date(), dates: MutableList<Course> = arrayListOf()) : super(id) {
        this.name = name
        this.start = start
        this.end = end
        this.dates = dates
    }

    companion object {
        val EMPTY = SchoolYear()
    }
}



