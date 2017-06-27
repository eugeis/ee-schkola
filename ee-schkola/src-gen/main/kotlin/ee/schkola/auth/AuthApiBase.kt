package ee.schkola.Auth

import ee.schkola.Person.User
import ee.schkola.Trace




open class Login {
    val principal: String
    val password: String
    val person: User
    val disabled: Boolean
    val lastLoginAt: 
    val trace: Trace


    constructor(principal: String = "", password: String = "", person: User = User.EMPTY, disabled: Boolean = false, lastLoginAt:  = ()) {
        this.principal = principal
        this.password = password
        this.person = person
        this.disabled = disabled
        this.lastLoginAt = lastLoginAt
    }

    companion object {
        val EMPTY = Login()
    }
}



