package ee.schkola

import ee.common.ext.logger
import ee.common.ext.toIntOr0
import ee.schkola.ImportFields.*
import ee.schkola.ImportFields.Church
import ee.schkola.person.*
import ee.schkola.student.SchoolApplication
import ee.schkola.student.SchoolYear
import java.nio.file.Path
import java.text.SimpleDateFormat
import java.util.*


enum class ImportFields(val regex: Regex) {
    SubmitDate(Regex(".*Submit date")), Ip(Regex("Ip")), Name(Regex("Name:")), Geschlecht(Regex("Geschlecht:")),
    Geburtsdatum(Regex("Geburtsdatum:")), Adresse(Regex("Adresse")), PLZ(Regex("PLZ")), Ort(Regex("Ort")),
    TelefonHandy(Regex("Telefon / Handy:")), Email(Regex("E-mail:")), Passbild(Regex("Passbild:")),
    Schulabschluss(Regex("Schulabschluss:")), Beruf(Regex("Beruf:")), Familienstand(Regex("Familienstand:")),
    Partner(Regex("Name des Ehepartners:")), ChildrenCount(Regex("Anzahl der Kinder:")),
    MemberOfChurch(Regex("Mitglied in Ortsgemeinde:")), Church(Regex("Name der Ortsgemeinde:")),
    ChurchServices(Regex("Gemeindedienste:")), Mentor(Regex("Mentor:")),
    MentorTelefon(Regex("Telefonnummer deines Pastors oder Mentors:")),
    Recommendation(Regex("Mündliche Empfehlung:")), RecommendationOf(Regex("Empfehlung von:")),
    ApplicationFor(Regex("Bewerbung für:"));

    companion object {
        fun findByCode(code: String): ImportFields? {
            return values().find { it.regex.matches(code) }
        }
    }
}

class SchkolaCsvParser {
    val log = logger()

    val dateFormatEn = SimpleDateFormat("yyyy-MM-dd")
    val dateTimeFormatEn = SimpleDateFormat("yyyy-MM-dd hh:mm:ss")
    val dateFormatDe = SimpleDateFormat("dd.MM.yyyy")
    val dateTimeFormatDe = SimpleDateFormat("dd.MM.yyyy hh:mm:ss")

    val f = HashMap<ImportFields, Int>()

    fun List<String>.f(field: ImportFields): String {
        return get(f[field]!!)
    }

    fun String.toPersonName(): PersonName {
        val parts = split(" ")
        return PersonName(parts.first(), parts.last())
    }

    fun String.toGender(): Gender {
        return if (equals("Weiblich")) Gender.FEMALE else Gender.MALE
    }

    fun String.toDate(): Date {
        try {
            return dateFormatEn.parse(this)
        } catch (e: Exception) {
            try {
                return dateFormatDe.parse(this)
            } catch (e: Exception) {
                return Date()
            }
        }
    }

    fun String.toDateTime(): Date {
        try {
            return dateTimeFormatEn.parse(this)
        } catch (e: Exception) {
            try {
                return dateTimeFormatDe.parse(this)
            } catch (e: Exception) {
                return Date()
            }
        }
    }

    fun String.toMaritalState(): MaritalState {
        return when (this) {
            "Verheiratet" -> MaritalState.MARRIED
            "ledig" -> MaritalState.SINGLE
            else -> MaritalState.SINGLE
        }
    }

    fun String.split(): List<String> {
        try {
            return trim().split(",").map { it.trim().substring(1, it.length - 1).trim() }
        } catch (e: Exception) {
            return emptyList()
        }
    }

    val graduation = hashMapOf<String, Graduation>()

    fun String.toGraduation(): Graduation {
        return graduation.getOrPut(this, { Graduation(id = this, level = GraduationLevel.UNKNOWN) })
    }

    fun load(schoolYearName: String, source: Path): List<SchoolApplication> {
        val schoolYear = SchoolYear(name = schoolYearName)
        val applications = arrayListOf<SchoolApplication>()
        f.clear()

        source.toFile().readLines().forEachIndexed { i, s ->
            log.info("parse: $s")
            try {
                if (i == 0) {
                    s.split().forEachIndexed { i, s ->
                        val field = Companion.findByCode(s)
                        if (field != null) f.put(field, i) else log.warn("Unknown code $s")
                    }
                } else {
                    val l = s.split()
                    if (l.isNotEmpty()) {
                        val person = Profile(gender = l.f(Geschlecht).toGender(), name = l.f(Name).toPersonName(),
                                birthday = l.f(Geburtsdatum).toDate(), address = Address(street = l.f(Adresse),
                                code = l.f(PLZ), city = l.f(Ort)), contact = Contact(phone = l.f(TelefonHandy), email = l.f(Email)),
                                photo = l.f(Passbild),
                                family = Family(l.f(Familienstand).toMaritalState(), childrenCount = l.f(ChildrenCount).toIntOr0(),
                                        partner = l.f(Partner).toPersonName()),
                                church = ChurchInfo(church = l.f(Church), member = l.f(MemberOfChurch).toBoolean(),
                                        services = l.f(ChurchServices).split(",").toMutableList()),
                                education = Education(graduation = l.f(Schulabschluss).toGraduation(), profession = l.f(Beruf)),
                                id = "2016_" + i
                        )
                        person.trace.createdAt = l.f(SubmitDate).toDateTime()

                        val application = SchoolApplication(recommendationOf = l.f(RecommendationOf).toPersonName(),
                                churchContactPerson = l.f(Mentor).toPersonName(), churchContact = Contact(phone = l.f(MentorTelefon)),
                                group = l.f(ApplicationFor), schoolYear = schoolYear, profile = person)

                        applications.add(application)
                    }
                }
            } catch (e: Exception) {
                log.error("Parse problem of $i: $s", e)
            }
        }
        return applications
    }
}