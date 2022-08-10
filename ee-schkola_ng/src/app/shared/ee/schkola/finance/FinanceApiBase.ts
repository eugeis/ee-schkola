import {Profile} from '../person/PersonApiBase'

export class Expense {
    purpose: ExpensePurpose = new ExpensePurpose()
    amount = 0
    profile: Profile = new Profile()
    date: Date = new Date()
    id = ''
}

export class ExpensePurpose {
    name = ''
    description = ''
    id = ''
}

export class Fee {
    student: Profile = new Profile()
    amount = 0
    kind: FeeKind = new FeeKind()
    date: Date = new Date()
    id = ''
}

export class FeeKind {
    name = ''
    amount = 0
    description = ''
    id = ''
}









