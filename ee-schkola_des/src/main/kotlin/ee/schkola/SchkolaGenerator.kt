package ee.schkola

import ee.design.gen.ts.DesignTypeScriptGenerator
import ee.lang.integ.dPath

fun main(args: Array<String>) {
    //val kotlinGenerator = DesignKotlinGenerator(Schkola)
    //val goGenerator = DesignGoGenerator(Schkola)
    //goGenerator.generate(dPath)
    val typeScriptGenerator = DesignTypeScriptGenerator(Schkola)
    typeScriptGenerator.generate(dPath)
}