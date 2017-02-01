package ee.schkola.library

import ee.schkola.SchkolaBase
import ee.schkola.person.PersonName
import java.util.*

data class Location(var shelf: String = "", var fold: String = "") {
    companion object {
        val EMPTY = Location("", "")
    }



}

fun Location?.orEmpty(): Location {
    return if (this != null) this else Location.EMPTY
}

open class Book : SchkolaBase {
    companion object {
        val EMPTY = Book("", "", "", "", Date(), "", "", PersonName.EMPTY, Location.EMPTY)
    }
    var title: String = ""
    var description: String = ""
    var language: String = ""
    var releaseDate: Date = Date()
    var edition: String = ""
    var category: String = ""
    var author: PersonName = PersonName.EMPTY
    var location: Location = Location.EMPTY

    constructor(id: String = "", title: String = "", description: String = "", language: String = "", releaseDate: Date = Date(), 
        edition: String = "", category: String = "", author: PersonName = PersonName.EMPTY, 
        location: Location = Location.EMPTY): super(id) {
        this.title = title
        this.description = description
        this.language = language
        this.releaseDate = releaseDate
        this.edition = edition
        this.category = category
        this.author = author
        this.location = location
    }


}

fun Book?.orEmpty(): Book {
    return if (this != null) this else Book.EMPTY
}


