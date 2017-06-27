package ee.schkola




open class SchkolaBase {
    val trace: Trace

    companion object {
        val EMPTY = SchkolaBase()
    }
}


open class Trace {
    val createdAt: 
    val updatedAt: 
    val modifiedBy: String


    constructor(createdAt:  = (), updatedAt:  = (), modifiedBy: String = "") {
        this.createdAt = createdAt
        this.updatedAt = updatedAt
        this.modifiedBy = modifiedBy
    }

    companion object {
        val EMPTY = Trace()
    }
}



