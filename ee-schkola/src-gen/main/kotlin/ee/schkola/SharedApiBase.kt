package ee.schkola

import java.util.*

data class Trace(var createdAt: Date = Date(), var updatedAt: Date = Date(), var modifiedBy: String = "") {
    companion object {
        val EMPTY = Trace(Date(), Date(), "")
    }



}

fun Trace?.orEmpty(): Trace {
    return if (this != null) this else Trace.EMPTY
}

abstract class SchkolaBase {

    var id: String = ""
    var trace: Trace = Trace.EMPTY
    constructor(id: String = "") {
        this.id = id
    }


}



