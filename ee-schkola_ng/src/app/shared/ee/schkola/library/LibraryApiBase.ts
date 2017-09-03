import {PersonName} from "../SharedApiBase"


export class Book {
    title: string
    description: string
    language: string
    releaseDate: Date
    edition: string
    category: string
    author: PersonName
    location: Location
    id: string


    findByPattern(pattern: string = "") : Book {
        throw new ReferenceError("Not implemented yet.");
    }

    findByTitle(title: string = "") : Book {
        throw new ReferenceError("Not implemented yet.");
    }

}






export class Location {
    shelf: string
    fold: string

}





