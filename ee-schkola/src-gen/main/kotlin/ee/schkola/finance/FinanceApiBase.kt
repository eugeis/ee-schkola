package ee.schkola.finance

import ee.schkola.SchkolaBase
import ee.schkola.person.Profile
import java.util.*


open class Expense : SchkolaBase {
    val purpose: ExpensePurpose
    val amount: Float
    val profile: Profile
    val date: Date


    constructor(id: String = "", purpose: ExpensePurpose = ExpensePurpose(), amount: Float = 0f,
        profile: Profile = Profile(), date: Date = Date()) : super(id) {
        this.purpose = purpose
        this.amount = amount
        this.profile = profile
        this.date = date
    }

    companion object {
        val EMPTY = Expense()
    }
}


open class ExpensePurpose : SchkolaBase {
    val name: String
    val description: String


    constructor(id: String = "", name: String = "", description: String = "") : super(id) {
        this.name = name
        this.description = description
    }

    companion object {
        val EMPTY = ExpensePurpose()
    }
}


open class Fee : SchkolaBase {
    val student: Profile
    val amount: Float
    val kind: FeeKind
    val date: Date


    constructor(id: String = "", student: Profile = Profile(), amount: Float = 0f, kind: FeeKind = FeeKind(),
        date: Date = Date()) : super(id) {
        this.student = student
        this.amount = amount
        this.kind = kind
        this.date = date
    }

    companion object {
        val EMPTY = Fee()
    }
}


open class FeeKind : SchkolaBase {
    val name: String
    val amount: Float
    val description: String


    constructor(id: String = "", name: String = "", amount: Float = 0f, description: String = "") : super(id) {
        this.name = name
        this.amount = amount
        this.description = description
    }

    companion object {
        val EMPTY = FeeKind()
    }
}



