import io.github.libvirgo.goctl.tasks.GoCtlApiTask
import io.github.libvirgo.goctl.tasks.GoCtlRpcTask

plugins {
    id("io.github.libvirgo.goctl")

}

goctl {
    goctlVersion.set("latest")
    apiSuffix.set("_api")
    rpcSuffix.set("_rpc")
}

subprojects {
    project.tasks.withType<GoCtlApiTask>().configureEach {
    }
    project.tasks.withType<GoCtlRpcTask>().configureEach {
    }
}
