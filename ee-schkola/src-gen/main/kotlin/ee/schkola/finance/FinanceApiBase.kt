package ee.schkola.Finance

import ee.schkola.Person.User
import ee.schkola.SchkolaBase




open class Expense : SchkolaBase {
    val purpose: ExpenseType
    val amount: Float
    val person: User
    val date: 


    constructor(purpose: ExpenseType = ExpenseType.EMPTY, amount: Float = 0f, person: User = User.EMPTY, date:  = ()) {
        this.purpose = purpose
        this.amount = amount
        this.person = person
        this.date = date
    }

    companion object {
        val EMPTY = Expense()
    }
}


open class ExpenseType : SchkolaBase {
    val name: String
    val description: String


    constructor(name: String = "", description: String = "") {
        this.name = name
        this.description = description
    }

    companion object {
        val EMPTY = ExpenseType()
    }
}


open class Fee : SchkolaBase {
    val student: User
    val amount: Float
    val type: FeeType
    val date: 


    constructor(student: User = User.EMPTY, amount: Float = 0f, type: FeeType = FeeType.EMPTY, date:  = ()) {
        this.student = student
        this.amount = amount
        this.type = type
        this.date = date
    }

    companion object {
        val EMPTY = Fee()
    }
}


open class FeeType : SchkolaBase {
    val name: String
    val amount: Float
    val description: String


    constructor(name: String = "", amount: Float = 0f, description: String = "") {
        this.name = name
        this.amount = amount
        this.description = description
    }

    companion object {
        val EMPTY = FeeType()
    }
}



