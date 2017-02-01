package ee.schkola.fx.view

import ee.schkola.fx.app.Styles.Companion.heading
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