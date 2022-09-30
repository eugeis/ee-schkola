package ee.schkola

import ee.design.gen.angular.DesignAngularGenerator
import ee.design.gen.go.DesignGoGenerator
import ee.design.gen.kt.DesignKotlinGenerator
import ee.design.gen.ts.DesignTypeScriptGenerator
import ee.lang.integ.dPath

fun main(args: Array<String>) {
    //val generator = DesignKotlinGenerator(Schkola)
    // val generator = DesignGoGenerator(Schkola)
    // generator.generate(dPath.toRealPath())

    val generatorTypeScript = DesignTypeScriptGenerator(Schkola)
    val generatorAngular = DesignAngularGenerator(Schkola)
    generatorTypeScript.generate(dPath)
    generatorAngular.generate(dPath)
    //val typeScriptGenerator = DesignTypeScriptGenerator(Schkola)
    //typeScriptGenerator.generate(dPath)
}
