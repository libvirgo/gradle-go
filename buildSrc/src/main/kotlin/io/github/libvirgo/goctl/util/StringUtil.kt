package io.github.libvirgo.goctl.util

import org.gradle.api.Project
import java.io.OutputStream

val env = mapOf(
    "HOME" to System.getenv("HOME"),
    "PATH" to "${System.getenv("HOME")}/go/bin:${System.getenv("PATH")}",
    "GOPATH" to "${System.getenv("HOME")}/go"
)

fun String.run(project: Project) {
    project.exec {
        it.apply {
            environment = env
            commandLine("sh")
            args("-c", this@run)
        }
    }
}

fun String.runOnlyErr(project: Project) {
    project.exec {
        it.apply {
            environment = env
            commandLine("sh")
            args("-c", this@runOnlyErr)
            standardOutput = OutputStream.nullOutputStream()
            errorOutput = System.out
        }
    }
}

fun String.runWithExitCode(project: Project): Int {
    return project.exec {
        it.apply {
            environment = env
            commandLine("sh")
            args("-c", this@runWithExitCode)
            isIgnoreExitValue = true
            standardOutput = OutputStream.nullOutputStream()
            errorOutput = OutputStream.nullOutputStream()
        }
    }.exitValue
}

fun String.binaryExist(project: Project): Boolean {
    return "which $this".runWithExitCode(project) == 0
}

fun String.ifEndsWith(end: String, block: (String, String) -> Unit) {
    if (this.endsWith(end)) {
        block(this, end)
    }
}