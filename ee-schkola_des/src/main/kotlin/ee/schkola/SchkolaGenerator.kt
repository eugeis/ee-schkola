package ee.schkola

import ee.design.gen.go.DesignGoGenerator
import ee.design.gen.kt.DesignKotlinGenerator
import ee.lang.integ.dPath

fun main(args: Array<String>) {
    //val generator = DesignKotlinGenerator(Schkola)
    val generator = DesignGoGenerator(Schkola)
    generator.generate(dPath)
    //val typeScriptGenerator = DesignTypeScriptGenerator(Schkola)
    //typeScriptGenerator.generate(dPath)
}