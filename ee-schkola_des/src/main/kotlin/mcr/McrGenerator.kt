package mcr

import ee.common.ext.exists
import ee.design.ModuleI
import ee.design.gen.kt.DesignKotlinGenerator
import ee.lang.initObjectTree
import org.slf4j.LoggerFactory
import java.nio.file.Path
import java.nio.file.Paths

class McrGenerator {
    private val log = LoggerFactory.getLogger(javaClass)

    fun generate(target: Path = Paths.get("../MindConnectRail").toAbsolutePath()) {
        milo.initObjectTree()
        val kotlinGenerator = DesignKotlinGenerator(Mcr, false)
        if (target.exists()) {
            val generator = kotlinGenerator.generatorFactory.pojoAndBuildersKt()
            val builders = setOf("BuilderIfcBase", "BuilderApiBase")
            kotlinGenerator.generate(target, generator) { model ->
                builders.contains(name()) && model is ModuleI<*> && model.name() == "Mindsphere"
            }
        } else {
            log.warn("can't generate, because target not found: {}", target)
        }
    }
}

fun main(args: Array<String>) {
    McrGenerator().generate()
}