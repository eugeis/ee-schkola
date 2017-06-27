package ee.schkola.Student

import ee.schkola.Person.Contact
import ee.schkola.Person.PersonName
import ee.schkola.Person.User
import ee.schkola.SchkolaBase
import ee.schkola.Trace


enum class AttendanceStatus {
    REGISTERED,
    CONFIRMED,
    CANCELED,
    PRESENT;

    fun isRegistered() : Boolean = this == REGISTERED
    fun isConfirmed() : Boolean = this == CONFIRMED
    fun isCanceled() : Boolean = this == CANCELED
    fun isPresent() : Boolean = this == PRESENT
}


fun String?.toAttendanceStatus(): AttendanceStatus {
    return if (this != null) AttendanceStatus.valueOf(this) else AttendanceStatus.REGISTERED
}


enum class GroupType {
    COURSE_GROUP,
    YEAR_GROUP;

    fun isCourseGroup() : Boolean = this == COURSE_GROUP
    fun isYearGroup() : Boolean = this == YEAR_GROUP
}


fun String?.toGroupType(): GroupType {
    return if (this != null) GroupType.valueOf(this) else GroupType.COURSE_GROUP
}




open class Attendance : SchkolaBase {
    val student: User
    val date: 
    val course: Course
    val hours: Int
    val status: AttendanceStatus
    val statusTrace: Trace
    val token: String
    val tokenTrace: Trace


    constructor(student: User = User.EMPTY, date:  = (), course: Course = Course.EMPTY, hours: Int = 0, 
                status: AttendanceStatus = AttendanceStatus.EMPTY, statusTrace: Trace = Trace.EMPTY, token: String = "", 
                tokenTrace: Trace = Trace.EMPTY) {
        this.student = student
        this.date = date
        this.course = course
        this.hours = hours
        this.status = status
        this.statusTrace = statusTrace
        this.token = token
        this.tokenTrace = tokenTrace
    }

    companion object {
        val EMPTY = Attendance()
    }
}


open class Course : SchkolaBase {
    val name: String
    val begin: 
    val end: 
    val teacher: PersonName
    val schoolYear: SchoolYear
    val fee: Float
    val description: String


    constructor(name: String = "", begin:  = (), end:  = (), teacher: PersonName = PersonName.EMPTY, 
                schoolYear: SchoolYear = SchoolYear.EMPTY, fee: Float = 0f, description: String = "") {
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
    val student: User
    val course: Course
    val grade: Float
    val gradeTrace: Trace
    val comment: String


    constructor(student: User = User.EMPTY, course: Course = Course.EMPTY, grade: Float = 0f, gradeTrace: Trace = Trace.EMPTY, 
                comment: String = "") {
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
    val type: GroupType
    val schoolYear: SchoolYear
    val representative: User
    val students: MutableList<User>
    val courses: MutableList<Course>


    constructor(name: String = "", type: GroupType = GroupType.EMPTY, schoolYear: SchoolYear = SchoolYear.EMPTY, 
                representative: User = User.EMPTY, students: MutableList<User> = arrayListOf(), 
                courses: MutableList<Course> = arrayListOf()) {
        this.name = name
        this.type = type
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
    val person: User
    val hasRecommendation: Boolean
    val recommendationOf: PersonName
    val mentor: PersonName
    val mentorContact: Contact
    val schoolYear: SchoolYear
    val group: String


    constructor(person: User = User.EMPTY, hasRecommendation: Boolean = false, recommendationOf: PersonName = PersonName.EMPTY, 
                mentor: PersonName = PersonName.EMPTY, mentorContact: Contact = Contact.EMPTY, 
                schoolYear: SchoolYear = SchoolYear.EMPTY, group: String = "") {
        this.person = person
        this.hasRecommendation = hasRecommendation
        this.recommendationOf = recommendationOf
        this.mentor = mentor
        this.mentorContact = mentorContact
        this.schoolYear = schoolYear
        this.group = group
    }

    companion object {
        val EMPTY = SchoolApplication()
    }
}


open class SchoolYear : SchkolaBase {
    val name: String
    val start: 
    val end: 
    val dates: MutableList<>


    constructor(name: String = "", start:  = (), end:  = (), dates: MutableList<> = arrayListOf()) {
        this.name = name
        this.start = start
        this.end = end
        this.dates = dates
    }

    companion object {
        val EMPTY = SchoolYear()
    }
}



