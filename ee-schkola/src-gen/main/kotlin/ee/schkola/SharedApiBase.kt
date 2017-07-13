package ee.schkola

import java.util.Date




open class SchkolaBase {
    val trace: Trace
    val id: String


    constructor(id: String = "") {
        this.id = id
        this.trace = Trace()
    }

    companion object {
        val EMPTY = SchkolaBase()
    }
}


open class Trace {
    var createdAt: Date
    var updatedAt: Date
    var modifiedBy: String


    constructor(createdAt: Date = Date(), updatedAt: Date = Date(), modifiedBy: String = "") {
        this.createdAt = createdAt
        this.updatedAt = updatedAt
        this.modifiedBy = modifiedBy
    }

    companion object {
        val EMPTY = Trace()
    }
}



