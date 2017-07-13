package ee.schkola.library

import ee.schkola.SchkolaBase
import ee.schkola.person.PersonName
import java.util.Date




open class Book : SchkolaBase {
    val title: String
    val description: String
    val language: String
    val releaseDate: Date
    val edition: String
    val category: String
    val author: PersonName
    val location: Location


    constructor(id: String = "", title: String = "", description: String = "", language: String = "", releaseDate: Date = Date(), 
                edition: String = "", category: String = "", author: PersonName = PersonName(), 
                location: Location = Location()) : super(id) {
        this.title = title
        this.description = description
        this.language = language
        this.releaseDate = releaseDate
        this.edition = edition
        this.category = category
        this.author = author
        this.location = location
    }

    companion object {
        val EMPTY = Book()
    }
}


open class Location {
    val shelf: String
    val fold: String


    constructor(shelf: String = "", fold: String = "") {
        this.shelf = shelf
        this.fold = fold
    }

    companion object {
        val EMPTY = Location()
    }
}



