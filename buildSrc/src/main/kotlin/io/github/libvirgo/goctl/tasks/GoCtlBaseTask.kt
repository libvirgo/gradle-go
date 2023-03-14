package io.github.libvirgo.goctl.tasks

import org.gradle.api.DefaultTask

abstract class GoCtlBaseTask : DefaultTask() {
    init {
        group = "goctl"
    }

}
