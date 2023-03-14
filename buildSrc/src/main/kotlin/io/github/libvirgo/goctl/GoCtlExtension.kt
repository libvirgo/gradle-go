package io.github.libvirgo.goctl

import org.gradle.api.provider.Property

interface GoCtlExtension {
    val goctlVersion: Property<String>
    val apiSuffix: Property<String>
    val rpcSuffix: Property<String>
}
