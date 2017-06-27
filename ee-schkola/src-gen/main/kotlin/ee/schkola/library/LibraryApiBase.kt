package ee.schkola.Library

import ee.schkola.Person.PersonName
import ee.schkola.SchkolaBase




open class Book : SchkolaBase {
    val title: String
    val description: String
    val language: String
    val releaseDate: 
    val edition: String
    val category: String
    val author: PersonName
    val location: Location


    constructor(title: String = "", description: String = "", language: String = "", releaseDate:  = (), edition: String = "", 
                category: String = "", author: PersonName = PersonName.EMPTY, location: Location = Location.EMPTY) {
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



