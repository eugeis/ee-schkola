import {Contact, PersonName, Profile} from "../person/PersonApiBase"
import {SchkolaBase, Trace} from "../SharedApiBase"


export class Attendance extends SchkolaBase {
    readonly student: Profile
    readonly date: Date
    readonly course: Course
    readonly hours: number
    readonly state: AttendanceState
    readonly stateTrace: Trace
    readonly token: string
    readonly tokenTrace: Trace

}


export class Course extends SchkolaBase {
    readonly name: string
    readonly begin: Date
    readonly end: Date
    readonly teacher: PersonName
    readonly schoolYear: SchoolYear
    readonly fee: number
    readonly description: string

}


export class Grade extends SchkolaBase {
    readonly student: Profile
    readonly course: Course
    readonly grade: number
    readonly gradeTrace: Trace
    readonly comment: string

}


export class Group extends SchkolaBase {
    readonly name: string
    readonly category: GroupCategory
    readonly schoolYear: SchoolYear
    readonly representative: Profile
    readonly students: Array<Profile>
    readonly courses: Array<Course>

}


export class SchoolApplication extends SchkolaBase {
    readonly profile: Profile
    readonly recommendationOf: PersonName
    readonly churchContactPerson: PersonName
    readonly churchContact: Contact
    readonly schoolYear: SchoolYear
    readonly group: string

}


export class SchoolYear extends SchkolaBase {
    readonly name: string
    readonly start: Date
    readonly end: Date
    readonly dates: Array<Date>

}


enum AttendanceState {
    REGISTERED,
    CONFIRMED,
    CANCELED,
    PRESENT
}


enum GroupCategory {
    COURSE_GROUP,
    YEAR_GROUP
}



