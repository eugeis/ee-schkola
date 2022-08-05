import {PersonName} from '../shared/SharedApiBase'
import {Profile} from '../person/PersonApiBase'


export class Account {
    name: PersonName = new PersonName()
    username = ''
    password = ''
    email = ''
    disabled = false
    roles: Array<string> = new Array()
    profile: Profile = new Profile()
    id = ''

}




export class UserCredentials {
    username = ''
    password = ''

}
