import {Contact, PersonName, Profile} from '../person/PersonApiBase'

export class Attendance {
    student: Profile = new Profile()
    date: Date = new Date()
    course: Course = new Course()
    hours = 0
    state: AttendanceState = AttendanceState.REGISTERED
    token = ''
    id = ''
}

export class Course {
    name = ''
    begin: Date = new Date()
    end: Date = new Date()
    teacher: PersonName = new PersonName()
    schoolYear: SchoolYear = new SchoolYear()
    fee = 0
    description = ''
    id = ''
}

export class Grade {
    student: Profile = new Profile()
    course: Course = new Course()
    grade = 0
    comment = ''
    id = ''
}

export class Group {
    name = ''
    category: GroupCategory = GroupCategory.COURSE_GROUP
    schoolYear: SchoolYear = new SchoolYear()
    representative: Profile = new Profile()
    students: Array<Profile> = new Array()
    courses: Array<Course> = new Array()
    id = ''
}

export class SchoolApplication {
    profile: Profile = new Profile()
    churchContactPerson: PersonName = new PersonName()
    churchContact: Contact = new Contact()
    churchCommitment = false
    schoolYear: SchoolYear = new SchoolYear()
    group = ''
    id = ''
}

export class SchoolYear {
    name = ''
    start: Date = new Date()
    end: Date = new Date()
    dates: Array<Date> = new Array()
    id = ''
}







export enum AttendanceState {
    REGISTERED,
    CONFIRMED,
    CANCELED,
    PRESENT
}

export enum GroupCategory {
    COURSE_GROUP,
    YEAR_GROUP
}



