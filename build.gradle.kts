import java.io.OutputStream

val env = mapOf(
    "HOME" to System.getenv("HOME"),
    "PATH" to "${System.getenv("HOME")}/go/bin:${System.getenv("PATH")}",
    "GOPATH" to "${System.getenv("HOME")}/go",
    "https_proxy" to "http://127.0.0.1:7890"
)

tasks.register("goctl-cli") {
    group = "dependencies"
    doLast {
        if (!"goctl".binaryExist()) {
            logger.warn("goctl not found, install latest version.")
            "go install github.com/zeromicro/go-zero/tools/goctl@latest".runOnlyErr()
        }
    }
}


subprojects {
    if (project.name.endsWith("api")) {
        tasks.register("goctl-api") {
            group = "goctl"
            dependsOn(":goctl-cli")
            doLast {
                "goctl api go -api $projectDir/${project.name}.api -dir $projectDir --style=goZero".runOnlyErr()
            }
        }
    }
    if (project.name.endsWith("rpc")) {
        tasks.register("goctl-rpc") {
            group = "goctl"
            dependsOn(":goctl-cli")
            doLast {
                "goctl rpc protoc $projectDir/${project.name}.proto --proto_path=$projectDir --go_out=$projectDir --go-grpc_out=$projectDir --zrpc_out=$projectDir --style=goZero".runOnlyErr()
            }
        }
    }
}

fun String.run() {
    exec {
        environment = env
        commandLine("sh")
        args("-c", this@run)
    }
}

fun String.runOnlyErr() {
    exec {
        environment = env
        commandLine("sh")
        args("-c", this@runOnlyErr)
        standardOutput = java.io.OutputStream.nullOutputStream()
        errorOutput = System.out
    }
}

fun String.binaryExist(): Boolean {
    return exec {
        environment = env
        commandLine("sh")
        args("-c", "which " + this@binaryExist)
        isIgnoreExitValue = true
        standardOutput = OutputStream.nullOutputStream()
        errorOutput = OutputStream.nullOutputStream()
    }.exitValue == 0
}
