package ee.schkola.fx

import ee.dsl.TypedComposite
import ee.dsl.fx.ModelApp
import ee.dsl.fx.view.ExplorerModel
import ee.schkola.model
import ee.dsl.system.dev.BuildToolFactory
import ee.dsl.system.dev.Gradle
import ee.dsl.system.dev.Maven
import ee.dsl.system.task.SystemTaskRegistry
import ee.dsl.task.PathResolver
import ee.dsl.task.TaskRepository
import java.nio.file.Paths

open class SchkolaModelApp() : ModelApp(ExplorerModel("SkolaModel", listOf("Components" to model())) {
    it is TypedComposite<*>
}, taskRepository())
/*
{
    companion object {
        @JvmStatic fun main(args: Array<String>) {
            Application.launch(SchkolaModelApp::class.java)
        }
    }
}
*/


private fun taskRepository(): TaskRepository {
    val ret = TaskRepository()
    val home = Paths.get("schkola")
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