package io.github.libvirgo.goctl.tasks

import io.github.libvirgo.goctl.util.runOnlyErr
import org.gradle.api.provider.Property
import org.gradle.api.tasks.Input
import org.gradle.api.tasks.Optional
import org.gradle.api.tasks.TaskAction

abstract class GoCtlRpcTask : GoCtlBaseTask() {

    @get:Input
    abstract val rpcFileSrc: Property<String>

    @get:Input
    abstract val outputDirSrc: Property<String>

    @TaskAction
    fun rpc() {
        """goctl rpc protoc ${rpcFileSrc.get()} \
            --proto_path=${outputDirSrc.get()} \
            --go_out=${outputDirSrc.get()} \
            --go-grpc_out=${outputDirSrc.get()} \
            --zrpc_out=${outputDirSrc.get()} \
            --style=goZero
        """.trimMargin().runOnlyErr(project)
    }
}
