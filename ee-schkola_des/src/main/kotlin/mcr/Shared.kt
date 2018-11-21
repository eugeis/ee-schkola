package mcr

import ee.design.Module
import ee.lang.*

object Shared : Module({ initObjectTree() }) {
    object ValueDate : Basic() {
        val value = prop { type(n.Any) }
        val date = prop { type(n.Date) }
    }

    object NameValueDate : Basic() {
        val name = propS()
        val value = prop { type(n.Any) }
        val date = prop { type(n.Date) }
    }

    object FromTo : Values() {
        val from = prop { type(n.Date) }
        val to = prop { type(n.Date) }
    }
}