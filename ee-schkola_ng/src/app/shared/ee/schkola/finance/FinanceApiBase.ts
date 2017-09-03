import {Profile} from "../person/PersonApiBase"


export class Expense {
    purpose: ExpensePurpose
    amount: number
    profile: Profile
    date: Date
    id: string

}


export class ExpensePurpose {
    name: string
    description: string
    id: string

}


export class Fee {
    student: Profile
    amount: number
    kind: FeeKind
    date: Date
    id: string

}


export class FeeKind {
    name: string
    amount: number
    description: string
    id: string

}









