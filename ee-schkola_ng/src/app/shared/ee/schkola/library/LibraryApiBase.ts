import {PersonName} from '../person/PersonApiBase'

export class Book {
    title = ''
    description = ''
    language = ''
    releaseDate: Date = new Date()
    edition = ''
    category = ''
    author: PersonName = new PersonName()
    location: Location = new Location()
    id = ''

    findByPattern(pattern = ''): Book {
        throw new ReferenceError('Not implemented yet.');
    }

    findByTitle(title = ''): Book {
        throw new ReferenceError('Not implemented yet.');
    }
}





export class Location {
    shelf = ''
    fold = ''
}





