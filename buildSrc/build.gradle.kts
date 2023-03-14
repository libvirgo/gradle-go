plugins {
    kotlin("jvm") version "1.8.0"
    `java-gradle-plugin`
}
repositories {
    mavenCentral()
}
gradlePlugin {
    plugins {
        create("goctlPlugin") {
            id = "io.github.libvirgo.goctl"
            implementationClass = "io.github.libvirgo.goctl.GoCtlPlugin"
        }
    }
}
kotlin {
    jvmToolchain(17)
}
dependencies {
    implementation(gradleApi())
}
