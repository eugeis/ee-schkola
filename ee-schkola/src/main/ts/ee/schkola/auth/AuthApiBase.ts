import {Profile} from "../person/PersonApiBase"
import {SchkolaBase} from "../SharedApiBase"


export class Account extends SchkolaBase {
    readonly username: string
    readonly password: string
    readonly email: string
    readonly disabled: boolean
    readonly lastLoginAt: Date
    readonly profile: Profile

}









