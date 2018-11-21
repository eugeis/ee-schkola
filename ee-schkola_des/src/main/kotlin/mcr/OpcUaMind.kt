package mcr

import ee.design.Config
import ee.design.Module
import ee.lang.*

object OpcUaMind : Module({
    namespace("mcr.opcua.mind").artifact("mcr-opcua-mind").initObjectTree()
}) {
    object SyncConfig : Config() {
        val opcUa = prop { type(OpcUa.OpcUaConfig) }
        val syncNodePath = propS { value("Objects/DeviceSet") }
        val subscriptionConfig = prop(OpcUa.SubscriptionConfig)

        val assetId = propS()
        val location = prop { type(Mind.Location) }
        val idPrefix = propS { value("opcUa") }
        val typeNamePrefix = propS()
        val baseAssetTypeId = propS { value("core.basicasset") }
        val namePartReplacements = prop { type(n.Map).doc("Replace all parts at building of target names") }
        val nameReplacements = prop { type(n.Map).doc("Replace complete names at building of target names") }
        val aspectConcatenationDepth = propI {
            value(2)
                    .doc("""An assent has aspects. Aspects are actually child nodes of a node
                            In order to reduce depth of asset tree, it is possible to map more layers
                            of children nodes to an aspect. The parameter defines how much children layers are mapped
                            to an asset. If AspectConcatenationDepth > 1, then aspect names are concatenated.""")
        }
    }
}