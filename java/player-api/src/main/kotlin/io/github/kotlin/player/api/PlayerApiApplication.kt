package io.github.kotlin.player.api

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication

/**
 * https://manuelernest0.medium.com/ Kotlin & Spring Boot: Building a Rest API
 */
@SpringBootApplication
class PlayerApiApplication

fun main(args: Array<String>) {
	runApplication<PlayerApiApplication>(*args)
}
