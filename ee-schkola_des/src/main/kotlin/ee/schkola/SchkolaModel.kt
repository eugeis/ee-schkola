package ee.schkola

import ee.lang.*
import ee.design.*


object Schkola : Comp({ artifact("ee-schkola").namespace("ee.schkola") }) {
    object Shared : Module() {
        object SchkolaBase : Entity({ virtual(true) }) {
            val id = buildId()
            val trace = prop { type(Trace).meta(true) }
            val constr = constr { params(p(id)) }
        }

        object Trace : Basic() {
            val createdAt = propDT().replaceable(true)
            val updatedAt = propDT().replaceable(true)
            val modifiedBy = propS().replaceable(true)
        }
    }

    object Auth : Module() {
        object Account : Entity({ superUnit(Shared.SchkolaBase) }) {
            val username = propS()
            val password = propS()
            val email = propS()
            val disabled = propB()
            val lastLoginAt = propDT()
            val profile = prop { type(Person.Profile) }

            object commands : Commands() {
                val register = createBy(username, email, password)
                val enable = command()
                val disable = command()
            }
        }
    }

    object Person : Module() {

        object Profile : Entity({ superUnit(Shared.SchkolaBase) }) {
            val gender = prop(Gender)
            val name = prop(PersonName)
            val birthName = propS()
            val birthday = propDT()
            val address = prop(Address)
            val contact = prop(Contact)
            val photoData = prop(n.Blob)
            val photo = propS()
            val family = prop(Family)
            val church = prop(ChurchInfo)
            val education = prop(Education)

            object queries : Queries() {
                val findByName = findBy(name)
                val findByEmail = findBy(contact.sub { email })
                val findByPhone = findBy(contact.sub { phone })
            }
        }

        object Church : Entity({ superUnit(Shared.SchkolaBase) }) {
            val name = propS()
            val address = prop(Address)
            val pastor = prop(PersonName)
            val contact = prop(Contact)
        }

        object Education : Basic() {
            val graduation = prop(Graduation)
            val profession = propS()
        }

        object ChurchInfo : Basic() {
            val church = propS()
            val association = propS()
            val member = propB()
            val services = prop(n.List)
        }

        object Family : Basic() {
            val maritalState = prop(MaritalState)
            val childrenCount = propI()
            val partner = prop(PersonName)
        }

        object PersonName : Basic() {
            val first = propS()
            val last = propS()
        }

        object Address : Basic() {
            val street = propS()
            val suite = propS()
            val city = propS()
            val code = propS()
            val country = propS()
        }

        object Contact : Basic() {
            val phone = propS()
            val email = propS()
            val cellphone = propS()
        }

        object Gender : EnumType() {
            val Unknown = lit()
            val Male = lit()
            val Female = lit()
        }

        object MaritalState : EnumType() {
            val Unknown = lit()
            val Single = lit()
            val Married = lit()
            val Separated = lit()
            val Divorced = lit()
            val Widowed = lit()
        }

        object GraduationLevel : EnumType() {
            val Unknown = lit()
            val MiddleSchool = lit({ doc("Hauptschule") })
            val SecondarySchool = lit({ doc("Realschule") })
            val HighSchool = lit({ doc("Abitur") })
            val TechnicalCollege = lit({ doc("Fachschule") })
            val College = lit({ doc("Hochschule") })
        }

        object Graduation : Entity({ superUnit(Shared.SchkolaBase) }) {
            val name = propS()
            val level = prop(GraduationLevel)
        }
    }

    object Finance : Module() {
        object Expense : Entity({ superUnit(Shared.SchkolaBase) }) {
            val purpose = prop(ExpensePurpose)
            val amount = propF()
            val profile = prop(Person.Profile)
            val date = propDT()
        }

        object ExpensePurpose : Entity({ superUnit(Shared.SchkolaBase) }) {
            val name = propS()
            val description = propS()
        }

        object Fee : Entity({ superUnit(Shared.SchkolaBase) }) {
            val student = prop(Person.Profile)
            val amount = propF { doc("Negative values are charge a fee and positive values are paying of it.") }
            val kind = prop(FeeKind)
            val date = propDT { doc("Deadline of a fee or paying date of it.") }
        }

        object FeeKind : Entity({ superUnit(Shared.SchkolaBase) }) {
            val name = propS()
            val amount = propF()
            val description = propS()
        }

    }

    object Student : Module() {
        object SchoolApplication : Entity({ superUnit(Shared.SchkolaBase) }) {
            val profile = prop(Person.Profile)
            val recommendationOf = prop(Person.PersonName)
            val churchContactPerson = prop(Person.PersonName)
            val churchContact = prop(Person.Contact)
            val schoolYear = prop(SchoolYear)
            val group = propS()
        }

        object SchoolYear : Entity({ superUnit(Shared.SchkolaBase) }) {
            val name = propS()
            val start = propDT()
            val end = propDT()
            val dates = prop(n.List.GT(n.Date))
        }

        object Course : Entity({ superUnit(Shared.SchkolaBase) }) {
            val name = propS()
            val begin = propDT()
            val end = propDT()
            val teacher = prop(Person.PersonName)
            val schoolYear = prop(SchoolYear)
            val fee = propF()
            val description = prop(n.Text)
        }

        object GroupCategory : EnumType() {
            val CourseGroup = lit()
            val YearGroup = lit()
        }

        object AttendanceState : EnumType() {
            val Registered = lit()
            val Confirmed = lit()
            val Canceled = lit()
            val Present = lit()
        }

        object Group : Entity({ superUnit(Shared.SchkolaBase) }) {
            val name = propS { nullable(false) }
            val category = prop(GroupCategory)
            val schoolYear = prop(SchoolYear)
            val representative = prop(Person.Profile)
            val students = prop(n.List.GT(Person.Profile))
            val courses = prop(n.List.GT(Course))
        }

        object Grade : Entity({ superUnit(Shared.SchkolaBase) }) {
            val student = prop { type(Person.Profile).nullable(false) }
            val course = prop { type(Course).nullable(false) }
            val grade = propF()
            val gradeTrace = prop(Shared.Trace)
            val comment = propS()
        }

        object Attendance : Entity({ superUnit(Shared.SchkolaBase) }) {
            val student = prop { type(Person.Profile).nullable(false) }
            val date = propDT { nullable(false) }
            val course = prop { type(Course).nullable(false) }
            val hours = propI()
            val state = prop(AttendanceState)
            val stateTrace = prop(Shared.Trace)
            val token = propS()
            val tokenTrace = prop(Shared.Trace)

            object commands : Commands() {
                val register = createBy(student, course)
                val confirm = command()
                val cancel = command()
            }
        }
    }

    object Library : Module() {
        object Location : Basic() {
            val shelf = propS()
            val fold = propS()
        }

        object Book : Entity({ superUnit(Shared.SchkolaBase) }) {
            val title = prop { nullable(false) }
            val description = prop(n.Text)
            val language = propS()
            val releaseDate = propDT()
            val edition = propS()
            val category = propS()
            val author = prop(Person.PersonName)
            val location = prop(Location)

            object queries : Queries() {
                val findByTitle = findBy(title)
                val findByAuthor = findBy(author)
                val findByPattern = findBy(p("pattern"))
            }

            object commands : Commands() {
                val unregister = deleteBy(id())
                val register = createBy(title, description, language, releaseDate, edition, category, author)
                val change = updateBy(title, description, language, releaseDate, edition, category, author)
                val changeLocation = updateBy(Location.shelf, Location.fold)
            }
        }
    }
}
