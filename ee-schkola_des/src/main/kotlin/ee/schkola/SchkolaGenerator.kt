package ee.schkola

import ee.design.gen.go.DesignGoGenerator
import ee.lang.integ.dPath

fun main(args: Array<String>) {
    //val generator = DesignKotlinGenerator(Schkola)
    val generator = DesignGoGenerator(Schkola)
    generator.generate(dPath)
}