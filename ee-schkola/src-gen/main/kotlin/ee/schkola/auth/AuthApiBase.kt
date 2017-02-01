package ee.schkola.auth

import ee.schkola.SchkolaBase
import ee.schkola.person.User
import java.util.*

open class Login : SchkolaBase {
    companion object {
        val EMPTY = Login("", "", "", User.EMPTY, false, Date())
    }
    var principal: String = ""
    var password: String = ""
    var person: User = User.EMPTY
    var disabled: Boolean = false
    var lastLoginAt: Date = Date()

    constructor(id: String = "", principal: String = "", password: String = "", person: User = User.EMPTY, disabled: Boolean = false, 
        lastLoginAt: Date = Date()): super(id) {
        this.principal = principal
        this.password = password
        this.person = person
        this.disabled = disabled
        this.lastLoginAt = lastLoginAt
    }


}

fun Login?.orEmpty(): Login {
    return if (this != null) this else Login.EMPTY
}


