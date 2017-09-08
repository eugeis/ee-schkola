import {PersonName} from "../SharedApiBase"
import {Profile} from "../person/PersonApiBase"


export class Account {
    name: PersonName
    username: string
    password: string
    email: string
    disabled: boolean
    roles: Array<string>
    profile: Profile
    id: string

}




export class UserCredentials {
    username: string
    password: string

}







