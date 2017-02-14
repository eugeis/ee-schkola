package ee.schkola

import ee.design.gen.DesignKotlinGenerator
import ee.lang.integ.eePath

fun main(args: Array<String>) {
    val generator = DesignKotlinGenerator(Schkola)
    generator.generate(eePath)
}