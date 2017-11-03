import {Contact, Profile} from "../person/PersonApiBase"
import {PersonName} from "../SharedApiBase"


export class Attendance {
    student: Profile
    date: Date
    course: Course
    hours: number
    state: AttendanceState
    token: string
    id: string

}


export class Course {
    name: string
    begin: Date
    end: Date
    teacher: PersonName
    schoolYear: SchoolYear
    fee: number
    description: string
    id: string

}


export class Grade {
    student: Profile
    course: Course
    grade: number
    comment: string
    id: string

}


export class Group {
    name: string
    category: GroupCategory
    schoolYear: SchoolYear
    representative: Profile
    students: Array<Profile>
    courses: Array<Course>
    id: string

}


export class SchoolApplication {
    profile: Profile
    churchContactPerson: PersonName
    churchContact: Contact
    churchCommitment: boolean
    schoolYear: SchoolYear
    group: string
    id: string

}


export class SchoolYear {
    name: string
    start: Date
    end: Date
    dates: Array<Date>
    id: string

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



