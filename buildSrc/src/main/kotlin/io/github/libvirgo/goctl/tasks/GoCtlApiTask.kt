package io.github.libvirgo.goctl.tasks

import io.github.libvirgo.goctl.util.runOnlyErr
import org.gradle.api.provider.Property
import org.gradle.api.tasks.Input
import org.gradle.api.tasks.TaskAction

abstract class GoCtlApiTask : GoCtlBaseTask() {
    @get:Input
    abstract val apiFileSrc: Property<String>

    @get:Input
    abstract val outputDirSrc: Property<String>

    @TaskAction
    fun api() {
        """goctl api go -api ${apiFileSrc.get()}\
            -dir ${outputDirSrc.get()} \
            --style goZero
        """.trimMargin().runOnlyErr(project)
    }
}