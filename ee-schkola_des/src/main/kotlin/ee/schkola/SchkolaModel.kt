package ee.schkola

import ee.lang.*
import ee.design.*


object Schkola : Comp({ artifact("ee-schkola").namespace("ee.schkola") }) {
    object Shared : Module() {
        object PersonName : Basic() {
            val first = propS()
            val last = propS()
        }
    }

    object Auth : Module() {
        object UserCredentials : Values() {
            val username = propS()
            val password = propS()
        }

        object Account : Entity() {
            val name = prop(Shared.PersonName)
            val username = propS { unique(true) }
            val password = propS { hidden(true) }
            val email = propS { unique(true) }
            val disabled = propB()
            val roles = propListT(n.String)
            val profile = prop { type(Person.Profile) }

            val login = command(username, email, password)
            val enable = updateBy(p(disabled, { value(false) }))
            val disable = updateBy(p(disabled, { value(true) }))

            val sendEnabledConfirmation = command()


            object AccountConfirmation : ProcessManager() {
                object Created : State({
                    handle(enable).to(Enabled).produce(sendEnabledConfirmation)
                })

                object Enabled : State() {

                }
            }
        }
    }

    object Person : Module() {

        object Profile : Entity() {
            val gender = prop(Gender)
            val name = prop(Shared.PersonName)
            val birthName = propS()
            val birthday = propDT()
            val address = prop(Address)
            val contact = prop(Contact)
            val photoData = prop(n.Blob)
            val photo = propS()
            val family = prop(Family)
            val church = prop(ChurchInfo)
            val education = prop(Education)

            //val findByName = findBy(name)
            val findByEmail = findBy(contact.sub { email })
            val findByPhone = findBy(contact.sub { phone })
        }

        object Church : Entity() {
            val name = propS()
            val address = prop(Address)
            val pastor = prop(Shared.PersonName)
            val contact = prop(Contact)
            val association = propS()
        }

        object Education : Basic() {
            val graduation = prop(Graduation)
            val other = propS()
            val profession = propS()
        }

        object ChurchInfo : Basic() {
            val church = propS()
            val member = propB()
            val services = prop(n.String)
        }

        object Family : Basic() {
            val maritalState = prop(MaritalState)
            val childrenCount = propI()
            val partner = prop(Shared.PersonName)
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
            val MiddleSchool = lit { doc("Hauptschule") }
            val SecondarySchool = lit { doc("Realschule") }
            val HighSchool = lit { doc("Abitur") }
            val TechnicalCollege = lit { doc("Fachschule") }
            val College = lit { doc("Hochschule") }
        }

        object Graduation : Entity() {
            val name = propS()
            val level = prop(GraduationLevel)
        }
    }

    object Finance : Module() {
        object Expense : Entity() {
            val purpose = prop(ExpensePurpose)
            val amount = propF()
            val profile = prop(Person.Profile)
            val date = propDT()
        }

        object ExpensePurpose : Entity() {
            val name = propS()
            val description = propS()
        }

        object Fee : Entity() {
            val student = prop(Person.Profile)
            val amount = propF { doc("Negative values are charge a fee and positive values are paying of it.") }
            val kind = prop(FeeKind)
            val date = propDT { doc("Deadline of a fee or paying date of it.") }
        }

        object FeeKind : Entity() {
            val name = propS()
            val amount = propF()
            val description = propS()
        }

    }

    object Student : Module() {
        object SchoolApplication : Entity() {
            val profile = prop(Person.Profile)
            val churchContactPerson = prop(Shared.PersonName)
            val churchContact = prop(Person.Contact)
            val churchCommitment = propB { doc("Do the responsible parties agree?") }

            val schoolYear = prop(SchoolYear)
            val group = propS()
        }

        object SchoolYear : Entity() {
            val name = propS()
            val start = propDT()
            val end = propDT()
            val dates = prop(n.List.GT(n.Date))
        }

        object Course : Entity() {
            val name = propS()
            val begin = propDT()
            val end = propDT()
            val teacher = prop(Shared.PersonName)
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

        object Group : Entity() {
            val name = propS { nullable(false) }
            val category = prop(GroupCategory)
            val schoolYear = prop(SchoolYear)
            val representative = prop(Person.Profile)
            val students = prop(n.List.GT(Person.Profile))
            val courses = prop(n.List.GT(Course))
        }

        object Grade : Entity() {
            val student = prop { type(Person.Profile).nullable(false) }
            val course = prop { type(Course).nullable(false) }
            val grade = propF()
            val comment = propS()
        }

        object Attendance : Entity() {
            val student = prop { type(Person.Profile).nullable(false) }
            val date = propDT { nullable(false) }
            val course = prop { type(Course).nullable(false) }
            val hours = propI()
            val state = prop(AttendanceState)
            val token = propS()

            val register = createBy(student, course, p(state, { value(AttendanceState.Registered) }))
            val confirm = updateBy(p(state, { value(AttendanceState.Confirmed) }))
            val cancel = updateBy(p(state, { value(AttendanceState.Canceled) }))
        }
    }

    object Library : Module() {
        object Location : Basic() {
            val shelf = propS()
            val fold = propS()
        }

        object Book : Entity() {
            val title = propS { nullable(false) }
            val description = prop(n.Text)
            val language = propS()
            val releaseDate = propDT()
            val edition = propS()
            val category = propS()
            val author = prop(Shared.PersonName)
            val location = prop(Location)

            val findByTitle = findBy(title)
            //val findByAuthor = findBy(author)
            val findByPattern = findBy(p("pattern"))

            val changeLocation = updateBy(location)
        }
    }
}
