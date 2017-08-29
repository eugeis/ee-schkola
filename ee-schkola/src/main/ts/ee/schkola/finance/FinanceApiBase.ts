import {Profile} from "../person/PersonApiBase"
import {SchkolaBase} from "../SharedApiBase"


export class Expense extends SchkolaBase {
    readonly purpose: ExpensePurpose
    readonly amount: number
    readonly profile: Profile
    readonly date: Date

}


export class ExpensePurpose extends SchkolaBase {
    readonly name: string
    readonly description: string

}


export class Fee extends SchkolaBase {
    readonly student: Profile
    readonly amount: number
    readonly kind: FeeKind
    readonly date: Date

}


export class FeeKind extends SchkolaBase {
    readonly name: string
    readonly amount: number
    readonly description: string

}









