package mcr

import ee.design.Comp

object Mcr : Comp({ name("mcr").artifact("mcr").namespace("mcr").modules(Shared, OpcUa, Mind, OpcUaMind) }) {

}
