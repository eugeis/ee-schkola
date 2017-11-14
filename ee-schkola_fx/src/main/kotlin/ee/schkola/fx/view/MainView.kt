package ee.schkola.fx.view

import tornadofx.*

class MainView : View("Bibelschule Stephanus") {
    override val root = hbox {
        label(title) {
            addClass(heading)
        }
    }
}