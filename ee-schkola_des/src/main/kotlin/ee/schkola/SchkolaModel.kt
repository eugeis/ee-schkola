package ee.schkola

import ee.lang.*
import ee.design.*


object Schkola : Comp({ artifact("ee-schkola").namespace("ee.schkola") }) {
    object Shared : Module() {
        object SchkolaBase : Entity({ virtual(true) }) {
            val id = id()
            val trace = prop { type(Trace).meta(true) }
        }

        object Trace : Basic() {
            val createdAt = prop(n.Date)
            val updatedAt = prop(n.Date)
            val modifiedBy = prop()
        }
    }

    object Auth : Module() {
        object Account : Entity({ superUnit(Shared.SchkolaBase) }) {
            val username = prop()
            val password = prop()
            val email = prop()
            val disabled = prop { type(n.Boolean) }
            val lastLoginAt = prop { type(n.Date) }
            val profile = prop { type(Person.Profile) }
            val trace = prop { type(Shared.Trace).meta(true) }

            object commands : CommandController() {
                val register = command(username, email, password)
                val enable = command()
                val disable = command()
            }
        }

    }

    object Person : Module() {

        object Profile : Entity({ superUnit(Shared.SchkolaBase) }) {
            val gender = prop(Gender)
            val name = prop(PersonName)
            val birthName = prop()
            val birthday = prop(n.Date)
            val address = prop(Address)
            val contact = prop(Contact)
            val photo = prop(n.Blob)
            val family = prop(Family)
            val church = prop(ChurchInfo)
            val education = prop(Education)

            object queries : QueryController() {
                val findByName = findBy(name)
                val findByEmail = findBy(contact.sub { email })
            }
        }

        object Church : Entity({ superUnit(Shared.SchkolaBase) }) {
            val name = prop()
            val address = prop(Address)
            val pastor = prop(PersonName)
            val contact = prop(Contact)
        }

        object Education : Basic() {
            val graduation = prop(Graduation)
            val profession = prop()
        }

        object ChurchInfo : Basic() {
            val church = prop()
            val association = prop()
            val member = prop(n.Boolean)
            val services = prop(n.List)
        }

        object Family : Basic() {
            val maritalState = prop(MaritalState)
            val childrenCount = prop(n.Int)
            val partner = prop(PersonName)
        }

        object PersonName : Basic() {
            val first = prop()
            val last = prop()
        }

        object Address : Basic() {
            val street = prop()
            val suite = prop()
            val city = prop()
            val code = prop()
            val country = prop()
        }

        object Contact : Basic() {
            val phone = prop()
            val email = prop()
            val cellphone = prop()

            object queries : QueryController() {
                val findByPhone = findBy(phone)
                val findByEmail = findBy(email)
            }
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
            val name = prop()
            val level = prop(GraduationLevel)
        }
    }

    object Finance : Module() {
        object Expense : Entity({ superUnit(Shared.SchkolaBase) }) {
            val purpose = prop(ExpensePurpose)
            val amount = prop(n.Float)
            val profile = prop(Person.Profile)
            val date = prop(n.Date)
        }

        object ExpensePurpose : Entity({ superUnit(Shared.SchkolaBase) }) {
            val name = prop()
            val description = prop()
        }

        object Fee : Entity({ superUnit(Shared.SchkolaBase) }) {
            val student = prop(Person.Profile)
            val amount = prop { type(n.Float).doc("Negative values are charge a fee and positive values are paying of it.") }
            val kind = prop(FeeKind)
            val date = prop { type(n.Date).doc("Deadline of a fee or paying date of it.") }
        }

        object FeeKind : Entity({ superUnit(Shared.SchkolaBase) }) {
            val name = prop()
            val amount = prop(n.Float)
            val description = prop()
        }

    }

    object Student : Module() {
        object SchoolApplication : Entity({ superUnit(Shared.SchkolaBase) }) {
            val profile = prop(Person.Profile)
            val recommendationOf = prop(Person.PersonName)
            val churchContactPerson = prop(Person.PersonName)
            val churchContact = prop(Person.Contact)
            val schoolYear = prop(SchoolYear)
            val group = prop()
        }

        object SchoolYear : Entity({ superUnit(Shared.SchkolaBase) }) {
            val name = prop()
            val start = prop(n.Date)
            val end = prop(n.Date)
            val dates = prop(n.List.GT(n.Date))
        }

        object Course : Entity({ superUnit(Shared.SchkolaBase) }) {
            val name = prop()
            val begin = prop(n.Date)
            val end = prop(n.Date)
            val teacher = prop(Person.PersonName)
            val schoolYear = prop(SchoolYear)
            val fee = prop(n.Float)
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
            val name = prop { nullable(false) }
            val category = prop(GroupCategory)
            val schoolYear = prop(SchoolYear)
            val representative = prop(Person.Profile)
            val students = prop(n.List.GT(Person.Profile))
            val courses = prop(n.List.GT(Course))
        }

        object Grade : Entity({ superUnit(Shared.SchkolaBase) }) {
            val student = prop { type(Person.Profile).nullable(false) }
            val course = prop { type(Course).nullable(false) }
            val grade = prop(n.Float)
            val gradeTrace = prop(Shared.Trace)
            val comment = prop()
        }

        object Attendance : Entity({ superUnit(Shared.SchkolaBase) }) {
            val student = prop { type(Person.Profile).nullable(false) }
            val date = prop { type(n.Date).nullable(false) }
            val course = prop { type(Course).nullable(false) }
            val hours = prop(n.Int)
            val state = prop(AttendanceState)
            val stateTrace = prop(Shared.Trace)
            val token = prop()
            val tokenTrace = prop(Shared.Trace)

            object commands : CommandController() {
                val register = command(student, course)
                val confirm = command()
                val cancel = command()
            }
        }
    }

    object Library : Module() {
        object Location : Basic() {
            val shelf = prop()
            val fold = prop()
        }

        object Book : Entity({ superUnit(Shared.SchkolaBase) }) {
            val title = prop { nullable(false) }
            val description = prop(n.Text)
            val language = prop()
            val releaseDate = prop(n.Date)
            val edition = prop()
            val category = prop()
            val author = prop(Person.PersonName)
            val location = prop(Location)

            object queries : QueryController() {
                val findByTitle = findBy(title)
                val findByAuthor = findBy(author)
                val findByPattern = findBy(p("pattern"))
            }

            object commands : CommandController() {
                val unregister = deleteBy(Shared.SchkolaBase.id)
                val register = createBy(title, description,
                        language, releaseDate, edition, category, author)
                val changeDetails = updateBy(title, description,
                        language, releaseDate, edition, category, author)
                val changeLocation = updateBy(Location.shelf, Location.fold)
            }
        }
    }
}
