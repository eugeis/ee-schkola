package ee.schkola

import ee.design.gen.go.DesignGoGenerator
import ee.design.gen.kt.DesignKotlinGenerator
import ee.lang.integ.dPath
import ee.lang.integ.eePath

fun main(args: Array<String>) {
    val generator = DesignKotlinGenerator(Schkola)
    //val generator = DesignGoGenerator(Schkola)
    generator.generate(dPath)
}