package io.github.libvirgo.goctl

import io.github.libvirgo.goctl.tasks.GoCtlApiTask
import io.github.libvirgo.goctl.tasks.GoCtlDepTask
import io.github.libvirgo.goctl.tasks.GoCtlRpcTask
import io.github.libvirgo.goctl.util.ifEndsWith
import org.gradle.api.Plugin
import org.gradle.api.Project

class GoCtlPlugin : Plugin<Project> {
    override fun apply(project: Project) {
        val goctlExtension = project.extensions.create("goctl", GoCtlExtension::class.java)
        project.tasks.register("goctl-cli", GoCtlDepTask::class.java) {
            it.goctlVersion.set(goctlExtension.goctlVersion)
        }

        project.subprojects { sub ->
            sub.afterEvaluate { subproject ->
                subproject.name.ifEndsWith(goctlExtension.apiSuffix.get()) { _, _ ->
                    subproject.tasks.register("goctl-api", GoCtlApiTask::class.java) {
                        it.apply {
                            dependsOn(":goctl-cli")
                            apiFileSrc.set("${subproject.projectDir}/${subproject.name}.api")
                            outputDirSrc.set("${subproject.projectDir}")
                        }
                    }
                }
                subproject.name.ifEndsWith(goctlExtension.rpcSuffix.get()) { _, _ ->
                    subproject.tasks.register("goctl-rpc", GoCtlRpcTask::class.java) {
                        it.apply {
                            rpcFileSrc.set("${subproject.projectDir}/${subproject.name}.proto")
                            outputDirSrc.set("${subproject.projectDir}")
                        }
                    }
                }
            }
        }
    }
}