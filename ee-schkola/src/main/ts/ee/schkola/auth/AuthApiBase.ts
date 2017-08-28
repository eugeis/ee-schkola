import { Profile } from "./ee/schkola/person"
import { SchkolaBase } from "./ee/schkola"


export class Account extends SchkolaBase {
    readonly username: string
    readonly password: string
    readonly email: string
    readonly disabled: boolean
    readonly lastLoginAt: Date
    readonly profile: Profile

}









