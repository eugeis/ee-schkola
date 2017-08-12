package ee.schkola.fx

import ee.lang.MultiHolder
import ee.lang.fx.ModelApp
import ee.lang.fx.view.ExplorerModel
import ee.lang.initObjectTrees
import ee.schkola.Schkola
import ee.system.dev.BuildToolFactory
import ee.system.dev.Gradle
import ee.system.dev.Maven
import ee.system.task.SystemTaskRegistry
import ee.task.PathResolver
import ee.task.TaskRepository
import java.nio.file.Paths

open class SchkolaDesignApp() : ModelApp(ExplorerModel("SchkolaDesign",
        listOf("Component" to Schkola.initObjectTrees())) {
    it is MultiHolder<*>
}, taskRepository())
/*
{
    companion object {
        @JvmStatic fun main(args: Array<String>) {
            Application.launch(SchkolaDesignApp::class.java)
        }
    }
}
*/


private fun taskRepository(): TaskRepository {
    val ret = TaskRepository()
    val home = Paths.get("Schkola")
    val pathResolver = PathResolver(home, hashMapOf("SkolaModel" to "src"))
    val maven = Maven(home = home.resolve("maven"), plugins = arrayListOf("eclipse"),
            defaultProfiles = listOf(),
            defaultParams = hashMapOf())
    val gradle = Gradle(home = home.resolve("gradle"))
    val buildToolFactory = BuildToolFactory(maven, gradle)
    val systemTaskRegistry = SystemTaskRegistry(pathResolver, buildToolFactory)
    systemTaskRegistry.register(ret)
    return ret
}