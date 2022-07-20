@file:Suppress("com.haulmont.jpb.DataClassEqualsAndHashCodeInspection")

package io.github.kotlin.player.api.model

import javax.persistence.*

@Entity
@Table(name = "tb_player")
data class Player (
    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    var id: Long,
    val name: String,
    val age: Int,
    val nationality: String
)