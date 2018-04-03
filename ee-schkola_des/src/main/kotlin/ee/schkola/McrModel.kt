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
            val id = propS()
            val name = propS { replaceable(true) }
            val description = propS { replaceable(true) }
            val type = prop(NodeType)
            val dataTypeId = propS { replaceable(true) }
            val dataType = propS { replaceable(true) }
            val nodes = prop { type(n.List.GT(Node)).mutable(true) }
        }

        object Object : Entity({
            superUnit(Node)
            constrSuper(Node.type.asParam(NodeType.Object))
        })
    }

    object Mindsphere : Module() {
        object Test1 : EnumType() {
            val Numeric = lit()
            val Guid = lit()
        }

        object Test2 : EnumType() {
            val r = propS()
            val Numeric = lit(r, "1")
            val String2 = lit(r, "2")
            val Guid = lit(r, "3")
            val Opaque = lit(r, "4")
        }
    }
}
