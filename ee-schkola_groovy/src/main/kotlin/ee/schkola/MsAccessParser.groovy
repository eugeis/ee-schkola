package ee.schkola

import com.healthmarketscience.jackcess.Database
import ee.msaccess.MsAccess
import ee.schkola.finance.Expense
import ee.schkola.finance.ExpensePurpose
import ee.schkola.finance.Fee
import ee.schkola.finance.FeeKind
import ee.schkola.person.*
import ee.schkola.student.*

class MsAccessParser {
    private String url
    private Database database

    MsAccessParser(String url) {
        this.url = url
    }

    public readAll() {

        database = MsAccess.open(url)

        //def rollen = readRolle()
        def schuljahrs = readSchoolYear()
        //def ausgabenzwecks = readExpensenzweck()
        //def beitragtypes = readFeeart()
        def dienste = readDienst()
        //def dokumenttypes = readDokumenttype()
        def gemeinden = readChurch()
        def persons = readPerson(gemeinden)
        def termins = readDate(schuljahrs)
        def fachs = readCourse(persons, schuljahrs)
        def gruppen = readGroup(schuljahrs, persons)
        def anwesenheiten = readAttendance(persons, fachs, termins)
        def ausgaben = readExpense(schuljahrs, persons, ausgabenzwecks)

        readCourseDate(fachs, termins)

        def beitrags = readFee(schuljahrs, fachs, persons, beitragtypes, termins)
        //def dokuments = readDokument(persons, dokumenttypes)
        def noten = readPersonCourse_Grade(persons, fachs)

        readGroupCourse(gruppen, fachs)
        readGroupPerson(gruppen, persons)
        readPersonDienst(persons, dienste)
        readPersonRolle(persons, rollen)

        /*
        //fill pastor vor- and nachname
        gemeinden.values().each { Church value ->
            if (value.pastorVorname) {
                Person pastor = persons.get(value.pastorVorname)
                value.pastorVorname = pastor.vorname
                value.pastorNachname = pastor.nachname
            }
        }
        */

        /*
        //activate personen in aktiv jahr
        gruppen.values().each { Group gruppe ->
            gruppe.studenten.each { Person person ->
                if (gruppe.type.isJahresgruppe()) {
                    person.student = true
                    person.gast = false
                }
                addToAktivInSchoolYear(person, gruppe.schuljahr)
            }
        }

        fachs.values().each { Course fach ->
            if (fach.lehrer) {
                addToAktivInSchoolYear(fach.lehrer, fach.schuljahr)
            }
        }

        //set lehrer=true
        fachs.values().each { Course fach ->
            fach.lehrer?.lehrer = true
            fach.lehrer?.gast = false
        }

        //activate
        def aktivRolles = [
                'Administrator',
                'Sachbearbeiter',
                'Koch'] as Set


        persons.values().each { Person person ->
            if (person.rollen.find { Rolle rolle -> aktivRolles.contains(rolle.rolle) }) {
                schuljahrs.values().each { addToAktivInSchoolYear(person, it) }
            }
        }
*/
        database = null

        new SchoolData(schuljahrs: schuljahrs.values().toList(),
                dienste: dienste.values().toList(), persons: persons.values().toList(), gemeinden: gemeinden.values().toList(), termins: termins.values().toList(),
                fachs: fachs.values().toList(), gruppen: gruppen.values().toList(), anwesenheiten: anwesenheiten.values().toList(), ausgaben: ausgaben.values().toList(),
                beitrags: beitrags.values().toList(), noten: noten.values().toList(), errors: [])
    }

    private readSchoolYear() {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("schuljahr")) {
            def schuljahr = new SchoolYear(start: rp(row, 'Schuljahrstart'), end: rp(row, 'Schuljahrende'))
            ret.put rp(row, 'schuljahr_ssid'), schuljahr
        }
        ret
    }

    private readDate(schuljahrs) {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("termin")) {
            ret.put rp(row, 'termin_ssid'), new Date(datum: rp(row, 'Datum'), schuljahr: schuljahrs.get(rp(row, 'schuljahr_ssid')))
        }
        ret
    }

    private readExpensenzweck() {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("_ausgabenzweck")) {
            ret.put rp(row, 'ausgabenzweck_ssid'), new ExpensePurpose(zweck: rp(row, 'Expensenzweck'))
        }
        ret
    }

    private readFeeart() {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("_beitragart")) {
            FeeKind entity = new FeeKind(type: rp(row, 'Feeart'), betrag: rp(row, 'Betrag'), einmalig: rp(row, 'Einmalig'))
            Calendar cal = Calendar.instance
            if (entity.type.startsWith('Anmeldegeb')) {
                cal.set 2008, 8, 1
                entity.faelligkeitsmonat = cal.time
            } else if (entity.type.startsWith('1. Rate')) {
                cal.set 2008, 8, 1
                entity.faelligkeitsmonat = cal.time
            } else if (entity.type.startsWith('2. Rate')) {
                cal.set 2009, 0, 1
                entity.faelligkeitsmonat = cal.time
            } else if (entity.type.startsWith('3. Rate')) {
                cal.set 2009, 2, 1
                entity.faelligkeitsmonat = cal.time
            }
            ret.put rp(row, 'beitragart_ssid'), entity
        }
        ret
    }

    private readGroup(schuljahrs, persons) {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("gruppe")) {
            ret.put rp(row, 'gruppe_ssid'), new Group(name: rp(row, 'Group'), schuljahr: schuljahrs.get(rp(row, 'schuljahr_ssid')), sprecher: persons.get(rp(row, 'gruppensprecher_ssid')), type: parseGroupType(rp(row, 'gruppentype_ssid')), studenten: [], faecher: [])
        }
        ret
    }

    private readExpense(Map schuljahrs, persons, ausgabenzwecks) {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("ausgaben")) {
            Date datum = rp(row, 'Datum')
            SchoolYear schuljahr = schuljahrs.find { String key, SchoolYear schuljahr ->
                datum.compareTo(schuljahr.start) >= 0 && datum.compareTo(schuljahr.ende) <= 0
            }?.value
            if (!schuljahr) {
                //first schuljahr
                schuljahr = schuljahrs.values().iterator().next()
            }
            ret.put rp(row, 'ausgaben_id'), new Expense(datum: datum, betrag: rp(row, 'Betrag'), mitarbeiter: persons.get(rp(row, 'person_ssid')), zweck: ausgabenzwecks.get(rp(row, 'ausgabenzweck_ssid')), schuljahr: schuljahr)
        }
        ret
    }

    /*
    private readRolle() {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("_rolle")) {
            ret.put rp(row, 'rolle_ssid'), new Rolle(rolle: rp(row, 'Rolle'))
        }
        ret
    }
*/

    /*
    private readDokumenttype() {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("_dokumenttype")) {
            ret.put rp(row, 'dokumenttype_ssid'), new Dokumenttype(type: rp(row, 'Dokumenttype'))
        }
        ret
    }
*/

    private readDienst() {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("_dienst")) {
            ret.put rp(row, 'dienst_ssid'), rp(row, 'Dienst')
        }
        ret
    }

    private buildSchoolYearmonatToDates(schuljahrs, termins) {
        def schuljahrmonatToDates = [:]
        schuljahrs.each { String schuljahrId, SchoolYear sj ->
            Date termin_september = termins.find { String terminId, Date t ->
                t.schuljahr == sj && t.datum.month == 8
            }.value
            Date termin_januar = termins.find { String terminId, Date t ->
                t.schuljahr == sj && t.datum.month == 0
            }.value
            Date termin_maerz = termins.find { String terminId, Date t ->
                t.schuljahr == sj && t.datum.month == 2
            }.value
            int firstYear = sj.start.year
            int secondYear = firstYear + 1
            schuljahrmonatToDates.put "$firstYear 9", termin_september
            schuljahrmonatToDates.put "$secondYear 1", termin_januar
            schuljahrmonatToDates.put "$secondYear 3", termin_maerz
        }
        schuljahrmonatToDates
    }

    private readFee(schuljahrs, fachs, persons, beitragtypes, termins) {
        def schuljahrmonatToDates = buildSchoolYearmonatToDates(schuljahrs, termins)

        def ret = [:]
        for (Map<String, Object> row : database.getTable("beitrag")) {
            Fee entity = new Fee(eingangsdatum: rp(row, 'Eingangsdatum'), istBetrag: rp(row, 'Ist-Betrag'), sollBetrag: rp(row, 'Soll-Betrag'), fach: fachs.get(rp(row, 'fach_ssid')), type: beitragtypes.get(rp(row, 'beitragart_ssid')), schuljahr: schuljahrs.get(rp(row, 'schuljahr_ssid')), student: persons.get(rp(row, 'person_ssid')))

            int firstYear = entity.schuljahr.start.year
            int secondYear = firstYear + 1
            Date termin

            if (entity.type.type.startsWith('Anmeldegeb')) {
                termin = schuljahrmonatToDates.get("$firstYear 9")
            } else if (entity.type.type.startsWith('1. Rate')) {
                termin = schuljahrmonatToDates.get("$firstYear 9")
            } else if (entity.type.type.startsWith('2. Rate')) {
                termin = schuljahrmonatToDates.get("$secondYear 1")
            } else if (entity.type.type.startsWith('3. Rate')) {
                termin = schuljahrmonatToDates.get("$secondYear 3")
            } else if (entity.type.type.startsWith('Fee nach Course')) {
                Course fach = fachs.find { String fachId, Course f ->
                    f == entity.fach
                }.value
                termin = fach.termine.iterator().next()
            }

            entity.faelligkeitstermin = termin
            ret.put rp(row, 'beitrag_id'), entity
        }
        ret
    }

    private readChurch() {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("_gemeinde")) {
            def gemeindepastor_sid = rp(row, 'gemeindepastor_sid')
            ret.put rp(row, 'gemeinde_ssid'), new Church(name: rp(row, 'Gemeinde'), stadt: rp(row, 'Adresse'), pastorVorname: gemeindepastor_sid, pastorNachname: gemeindepastor_sid)
        }
        ret
    }

    private readCourse(persons, schuljahrs) {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("fach")) {
            ret.put rp(row, 'fach_ssid'), new Course(name: rp(row, 'Course'), gebuehr: rp(row, 'Gebuehr'), schuljahr: schuljahrs.get(rp(row, 'schuljahr_ssid')), lehrer: persons.get(rp(row, 'lehrer_ssid')), termine: [])
        }
        ret
    }

    private readPerson(gemeinden) {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("person")) {
            ret.put rp(row, 'person_ssid'), new Profile(anrede: parseGender(rp(row, 'anrede_ssid')), vorname: rp(row, 'Vorname'), nachname: rp(row, 'Name'),
                    strasse: rp(row, 'Strasse'), plz: rp(row, 'PLZ'), stadt: rp(row, 'Ort'), land: rp(row, 'land_ssid'), telefon: rp(row, 'Tel_Priv'), mobil: rp(row, 'Handy'), email: rp(row, 'E-Mail'),
                    familienstand: parseMartialStatus(rp(row, 'familienstand_ssid')), schulabschluss: parseGraduationType(rp(row, 'schulabschluss_ssid')),
                    gemeinde: gemeinden.get(rp(row, 'gemeinde_ssid')), student: false, lehrer: false, mitarbeiter: false, gast: true, geburtsname: rp(row, 'Geburtsname'),
                    beruf: rp(row, 'Beruf'), geburtsdatum: rp(row, 'Geburtsdatum'), fotoPfad: rp(row, 'fotoPfad'), kinderanzahl: rp(row, 'Kinderanzahl'), dienste: [], rollen: [])
        }
        ret
    }

    private readAttendance(persons, fachs, termins) {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("anwesenheit")) {
            int sumStunden = 0

            if (rp(row, 'stunde_1'))
                sumStunden++
            if (rp(row, 'stunde_2'))
                sumStunden++
            if (rp(row, 'stunde_3'))
                sumStunden++
            if (rp(row, 'stunde_4'))
                sumStunden++
            if (rp(row, 'stunde_5'))
                sumStunden++
            if (rp(row, 'stunde_6'))
                sumStunden++
            if (rp(row, 'stunde_7'))
                sumStunden++
            if (rp(row, 'stunde_8'))
                sumStunden++

            Date termin = termins.get(rp(row, 'termin_ssid'))
            ret.put rp(row, 'anwesenheit_id'), new Attendance(person: persons.get(rp(row, 'person_ssid')),
                    fach: fachs.get(rp(row, 'fach_ssid')), termin: termin, stunden: sumStunden)
        }
        ret
    }

    /*
    private readDokument(persons, dokumenttypes) {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("dokument")) {
            ret.put rp(row, 'dokument_id'), new Dokument(elName: rp(row, 'Dokument'), pfad: rp(row, 'Pfad'), beschreibung: rp(row, 'Beschreibung'), datum: rp(row, 'Datum'), type: dokumenttypes.get(rp(row, 'dokumenttype_ssid')), person: persons.get(rp(row, 'person_ssid')))
        }
        ret
    }
    */

    private readPersonDienst(persons, dienste) {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("person_dienst")) {
            persons.get(rp(row, 'person_ssid')).dienste.add(dienste.get(rp(row, 'dienst_ssid')))
        }
        ret
    }

    private readPersonCourse_Grade(persons, fachs) {
        def ret = [:]
        for (Map<String, Object> row : database.getTable("person_fach")) {
            ret.add new Grade(note: rp(row, 'Grade'), datum: rp(row, 'Abschlussdatum'), studenten: persons.get(rp(row, 'person_ssid')), fach: fachs.get(rp(row, 'fach_ssid')))
        }
        ret
    }

    private readPersonRolle(persons, rollen) {
        for (Map<String, Object> row : database.getTable("person_rolle")) {
            persons.get(rp(row, 'person_ssid')).rollen.add(rollen.get(rp(row, 'rolle_ssid')))
        }
    }

    private readGroupPerson(gruppen, persons) {
        for (Map<String, Object> row : database.getTable("gruppe_person")) {
            gruppen.get(rp(row, 'gruppe_ssid')).studenten.add(persons.get(rp(row, 'person_ssid')))
        }
    }

    private readGroupCourse(gruppen, fachs) {
        for (Map<String, Object> row : database.getTable("gruppe_fach")) {
            gruppen.get(rp(row, 'gruppe_ssid')).faecher.add(fachs.get(rp(row, 'fach_ssid')))
        }
    }

    private readCourseDate(fachs, termins) {
        for (Map<String, Object> row : database.getTable("fach_termin")) {
            fachs.get(rp(row, 'fach_ssid')).termine.add(termins.get(rp(row, 'termin_ssid')))
        }
    }

    private MaritalState parseMartialStatus(familienstand_ssid) {
        def ret
        if (familienstand_ssid == 'ledig') {
            ret = MartialStatus.SINGLE
        } else if (familienstand_ssid.contains('verheirat')) {
            ret = MartialStatus.MARRIED
        } else {
            ret = MartialStatus.UNKNOWN
        }
        ret
    }

    private GraduationLevel parseGraduationType(schulabschluss_ssid) {
        def ret
        if (schulabschluss_ssid == 'Wirtschaftsschule') {
            ret = GraduationType.TECHNICAL_COLLEGE
        } else if (schulabschluss_ssid == 'Abitur') {
            ret = GraduationType.HIGH_SCHOOL
        } else if (schulabschluss_ssid == 'Hauptschule') {
            ret = GraduationType.SECONDARY_SCHOOL
        } else if (schulabschluss_ssid == 'Mittlere Reife') {
            ret = GraduationType.MIDDLE_SCHOOL
        } else if (schulabschluss_ssid == 'Realschule') {
            ret = GraduationType.MIDDLE_SCHOOL
        } else if (schulabschluss_ssid == 'Master of Theologie') {
            ret = GraduationType.COLLEGE
        } else {
            ret = null
        }
        ret
    }

    private Gender parseGender(anrede_ssid) {
        def ret
        if (anrede_ssid == 'bruder') {
            ret = Gender.MALE
        } else if (anrede_ssid == 'schwester') {
            ret = Gender.FEMALE
        } else {
            ret = Gender.UNKNOWN
        }
        ret
    }

    private Group parseGroupType(gruppentype_ssid) {
        def ret
        if (gruppentype_ssid == 'frei') {
            ret = GroupType.Coursegruppe
        } else if (gruppentype_ssid == 'standard') {
            ret = GroupType.Jahresgruppe
        } else {
            ret = null
        }
        ret
    }

    /**
     * Red property
     */
    private rp(Map rs, String propertyName) {
        if (!rs.containsKey(propertyName)) {
            System.err.println(propertyName + " does not exists in " + rs)
        }
        def ret = rs.get(propertyName)
        ret
    }
}

