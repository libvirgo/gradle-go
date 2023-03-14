package io.github.libvirgo.goctl.tasks

import io.github.libvirgo.goctl.util.binaryExist
import io.github.libvirgo.goctl.util.runOnlyErr
import org.gradle.api.provider.Property
import org.gradle.api.tasks.Input
import org.gradle.api.tasks.TaskAction

abstract class GoCtlDepTask : GoCtlBaseTask() {
    @get:Input
    abstract val goctlVersion: Property<String>

    @TaskAction
    fun goctlCli() {
        if (!"goctl".binaryExist(project)) {
            logger.warn("goctl not exist, while install goctl@${goctlVersion.get()}")
            "go install github.com/zeromicro/go-zero/tools/goctl@latest".runOnlyErr(project)
        }
    }
}