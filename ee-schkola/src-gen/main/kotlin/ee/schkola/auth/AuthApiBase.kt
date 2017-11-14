package ee.schkola.auth

import ee.schkola.SchkolaBase
import ee.schkola.person.Profile
import java.util.*


open class Account : SchkolaBase {
    val username: String
    val password: String
    val email: String
    val disabled: Boolean
    val lastLoginAt: Date
    val profile: Profile


    constructor(id: String = "", username: String = "", password: String = "", email: String = "", disabled: Boolean = false, 
                lastLoginAt: Date = Date(), profile: Profile = Profile()) : super(id) {
        this.username = username
        this.password = password
        this.email = email
        this.disabled = disabled
        this.lastLoginAt = lastLoginAt
        this.profile = profile
    }

    companion object {
        val EMPTY = Account()
    }
}



