package ee.schkola.fx.view

import tornadofx.View
import tornadofx.addClass
import tornadofx.hbox
import tornadofx.label

class MainView : View("Bibelschule Stephanus") {
    override val root = hbox {
        label(title) {
            addClass(heading)
        }
    }
}