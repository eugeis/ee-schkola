package ee.schkola

import ee.schkola.student.SchoolApplication
import java.nio.file.Path
import java.nio.file.Paths
import java.util.*

fun main(args: Array<String>) {
    val importer = SchkolaCsvParser()
    val basePath = Paths.get("/Users/eugeis/Documents/BSS-Verwaltung/")

    val applications = parseApplicationForms(basePath, importer)
    applications.forEach {
        println("${it.size}: ${it.joinToString(", ") { it.person.name.toString() }}")
    }

}

fun parseApplicationForms(basePath: Path, importer: SchkolaCsvParser): ArrayList<List<SchoolApplication>> {
    val applications = arrayListOf(importer.load("2014-2016",
            basePath.resolve("BSS-2014-2015/Bewerbung2014_20160908.csv")),
            importer.load("2015-2016",
                    basePath.resolve("BSS-2015-2016/Bewerbung2015_20160908.csv")),
            importer.load("2016-2017",
                    basePath.resolve("BSS-2016-2017/Bewerbung2016_20160905.csv")))
    return applications
}