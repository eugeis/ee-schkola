package com.siemens.mcr

import ee.common.ext.exists
import ee.design.gen.kt.DesignKotlinGenerator
import org.slf4j.LoggerFactory
import java.nio.file.Path
import java.nio.file.Paths

class McrGenerator {
    private val log = LoggerFactory.getLogger(javaClass)

    fun generate(target: Path = Paths.get("../MindConnectRail").toAbsolutePath()) {
        val kotlinGenerator = DesignKotlinGenerator(Mcr, false)
        if (target.exists()) {
            kotlinGenerator.generate(target)
        } else {
            log.warn("can't generate, because target not found: {}", target)
        }
    }
}

fun main(args: Array<String>) {
    McrGenerator().generate()
}