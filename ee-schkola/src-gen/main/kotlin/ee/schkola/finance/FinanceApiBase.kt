package ee.schkola.finance

import ee.schkola.SchkolaBase
import ee.schkola.person.User
import java.util.*

open class Expense : SchkolaBase {
    companion object {
        val EMPTY = Expense("", ExpenseType.EMPTY, 0f, User.EMPTY, Date())
    }
    var purpose: ExpenseType = ExpenseType.EMPTY
    var amount: Float = 0f
    var person: User = User.EMPTY
    var date: Date = Date()

    constructor(id: String = "", purpose: ExpenseType = ExpenseType.EMPTY, amount: Float = 0f, person: User = User.EMPTY, 
        date: Date = Date()): super(id) {
        this.purpose = purpose
        this.amount = amount
        this.person = person
        this.date = date
    }


}

fun Expense?.orEmpty(): Expense {
    return if (this != null) this else Expense.EMPTY
}

open class ExpenseType : SchkolaBase {
    companion object {
        val EMPTY = ExpenseType("", "", "")
    }
    var name: String = ""
    var description: String = ""

    constructor(id: String = "", name: String = "", description: String = ""): super(id) {
        this.name = name
        this.description = description
    }


}

fun ExpenseType?.orEmpty(): ExpenseType {
    return if (this != null) this else ExpenseType.EMPTY
}

open class Fee : SchkolaBase {
    companion object {
        val EMPTY = Fee("", User.EMPTY, 0f, FeeType.EMPTY, Date())
    }
    var student: User = User.EMPTY
    var amount: Float = 0f
    var type: FeeType = FeeType.EMPTY
    var date: Date = Date()

    constructor(id: String = "", student: User = User.EMPTY, amount: Float = 0f, type: FeeType = FeeType.EMPTY, date: Date = Date()): super(id) {
        this.student = student
        this.amount = amount
        this.type = type
        this.date = date
    }


}

fun Fee?.orEmpty(): Fee {
    return if (this != null) this else Fee.EMPTY
}

open class FeeType : SchkolaBase {
    companion object {
        val EMPTY = FeeType("", "", 0f, "")
    }
    var name: String = ""
    var amount: Float = 0f
    var description: String = ""

    constructor(id: String = "", name: String = "", amount: Float = 0f, description: String = ""): super(id) {
        this.name = name
        this.amount = amount
        this.description = description
    }


}

fun FeeType?.orEmpty(): FeeType {
    return if (this != null) this else FeeType.EMPTY
}


