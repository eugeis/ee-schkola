package com.siemens.mcr

import ee.design.*
import ee.lang.*


object Mcr : Comp({ artifact("mcr").namespace("com.siemens.mcr") }) {
    object Shared : Module() {}

    object Opcua : Module() {

        object IdType : EnumType() {
            val Numeric = lit()
            val String = lit()
            val Guid = lit()
            val Opaque = lit()
        }

        object NodeType : EnumType() {
            val Unspecified = lit()
            val Object = lit()
            val Variable = lit()
            val Method = lit()
            val ObjectType = lit()
            val VariableType = lit()
            val ReferenceType = lit()
            val DataType = lit()
            val View = lit()
        }

        object Id : Basic() {
            val ns = propI()
            val id = prop { type(n.Any) }
            val type = prop { type(IdType) }
        }

        object Node : Entity() {
            val id = prop { type(Id) }
            val name = propS({ replaceable(true) })
            val type = prop(NodeType)
            val nodes = prop { type(n.List.GT(Node)).mutable(true) }
        }
    }

    object Mindsphere : Module() {}
}
