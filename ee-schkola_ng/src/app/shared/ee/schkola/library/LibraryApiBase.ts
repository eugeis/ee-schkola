import {PersonName} from "../person/PersonApiBase"
import {SchkolaBase} from "../SharedApiBase"


export class Book extends SchkolaBase {
    readonly title: string
    readonly description: string
    readonly language: string
    readonly releaseDate: Date
    readonly edition: string
    readonly category: string
    readonly author: PersonName
    readonly location: Location


    findByAuthor(author: PersonName = new PersonName()) : Book {
        throw new ReferenceError("Not implemented yet.");
    }

    findByPattern(pattern: string = "") : Book {
        throw new ReferenceError("Not implemented yet.");
    }

    findByTitle(title: string = "") : Book {
        throw new ReferenceError("Not implemented yet.");
    }

}






export class Location {
    readonly shelf: string
    readonly fold: string

}





