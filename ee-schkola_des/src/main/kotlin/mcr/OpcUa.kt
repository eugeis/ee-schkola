package mcr

import ee.design.*
import ee.lang.*

object OpcUa : Module({ artifact("mcr-opcua").initObjectTree() }) {
    val nodeIdEmpty = p("nodeId", milo.core.NodeId).value("NodeIdEmpty")
    val nodeIdEmptyAsId = p("id", milo.core.NodeId).value("NodeIdEmpty")

    object OpcUa : Controller({ ifc() }) {
        val info = op { ret(n.String) }
        val browser = op { ret(OpcUaBrowser) }
        val reader = op { ret(OpcUaReader) }
        val browseReaderBlock = op { ret(OpcUaBrowseReaderBlocking) }
        val browseReader = op { ret(OpcUaBrowseReader) }
        val subscriber = op { ret(OpcUaSubscriber) }
        val close = op()
    }

    object OpcUaBrowser : Controller({ ifc() }) {
        val root = op {
            suspend()
            ret(Node)
        }
        val findNode = op(p("nodeId", milo.core.NodeId)) {
            suspend()
            ret { type(Node).nullable() }
        }
        val findNodes = op(p("nodeIds", n.Collection.GT(milo.core.NodeId))) {
            suspend()
            ret(n.Map.GT(milo.core.NodeId, Node))
        }
        val findNodesRecursive = op(p("nodeId", milo.core.NodeId), p("depth", n.Int).nullable()) {
            suspend()
            ret(n.Map.GT(milo.core.NodeId, Node))
        }
        val findNodesRecursiveAll = op(p("nodeId", milo.core.NodeId)) {
            suspend()
            ret(n.Map.GT(milo.core.NodeId, Node))
        }
    }

    object OpcUaReader : Controller({ ifc() }) {
        val readValue = op(p("variableNodeId", milo.core.NodeId)) {
            suspend()
            ret { type(Shared.ValueDate).nullable() }
        }

        val readValues = op(p("variableNodeIds", n.Collection.GT(milo.core.NodeId))) {
            suspend()
            ret(n.Map.GT(milo.core.NodeId, Shared.ValueDate))
        }

        val readValuesTo = op(p("variableNodeIds", n.Collection.GT(milo.core.NodeId)),
                p("destination", lambda(p("nodeId", milo.core.NodeId),
                        p("value", n.Any), p("time", n.Date)))).suspend()

        val readHistoricalValues = op(p("variableNodeIds", n.Collection.GT(milo.core.NodeId)),
                p("range", Shared.FromTo)) {
            suspend()
            ret(n.Map.GT(milo.core.NodeId, n.List.GT(Shared.ValueDate)))
        }
        val readHistoricalValuesTo = op(p("variableNodeIds", n.Collection.GT(milo.core.NodeId)),
                p("range", Shared.FromTo),
                p("destination", lambda(p("nodeId", milo.core.NodeId),
                        p("valueDates", n.List.GT(Shared.ValueDate))))).suspend()
    }

    object OpcUaSubscriber : Controller({ ifc() }) {
        val subscribe = op(p("variableNodeId", milo.core.NodeId), p("config", SubscriptionConfig),
                p("callback", lambda(p("nodeId", milo.core.NodeId),
                        p("value", n.Any), p("time", n.Date)) )) {
            suspend()
            ret(Subscription)
        }

        val subscribeAll = op(p("variableNodeIds", n.Collection.GT(milo.core.NodeId)),
                p("config", SubscriptionConfig),
                p("callback", lambda(p("nodeId", milo.core.NodeId),
                        p("value", n.Any), p("time", n.Date)) )) {
            suspend()
            ret(n.List.GT(Subscription))
        }
    }

    object OpcUaBrowseReaderBlocking : Controller({ ifc() }) {
        val findValuesByRootPath = op(p("paths", n.Collection),
                p("namespaceMapping", n.Map.GT(n.String, n.Int))) {
            ret(n.Map.GT(n.String, n.Any))
        }

        val findValuesByPath = op(p("startingNode", milo.core.NodeId), p("paths", n.Collection),
                p("namespaceMapping", n.Map.GT(n.String, n.Int))) {
            ret(n.Map.GT(n.String, n.Any))
        }
    }

    object OpcUaBrowseReader : Controller({ ifc() }) {
        val findValuesByRootPath = op(p("paths", n.Collection),
                p("namespaceMapping", n.Map.GT(n.String, n.Int))) {
            suspend()
            ret(n.Map.GT(n.String, n.Any))
        }

        val findValuesByPath = op(p("startingNode", milo.core.NodeId), p("paths", n.Collection),
                p("namespaceMapping", n.Map.GT(n.String, n.Int))) {
            suspend()
            ret(n.Map.GT(n.String, n.Any))
        }
    }

    object OpcUaConfig : Values() {
        val uri = propS { value("opc.tcp://localhost:4840") }
        val timeout = 1.minutes.doc("used for connection and " +
                "all other requests like read, subscription")
        val maxTimeout = 2.minutes.doc("in case if the timeout reached," +
                " the reconnection process doubles the current timeout but only until the max timeout")
        val maxBrowseDescriptionsInRequest = propI { value(5000) }
        val maxReadValuesInRequest = propI { value(10000) }
        val maxItemsInSubscription = propI { value(10000) }
        val acknowledgeTimeout = propI { value(10000) }
    }

    object IdNameType : Basic() {
        val id = props().addItem(nodeIdEmptyAsId)
        val name = propS()
        val type = prop(milo.core.NodeClass)
    }

    object Node : Entity() {
        val me = prop(IdNameType)
        val description = propS()
        val parent = prop(IdNameType).nullable()
        val dataType = prop(IdNameType).nullable()
        val typeDef = prop(IdNameType).nullable()
        val nodes = prop { type(n.List.GT(IdNameType)) }
    }

    object NodeEager : Entity() {
        val node = prop(Node)
        val parent = prop(NodeEager).replaceable().nullable()
        val dataType = prop(NodeEager).replaceable().nullable()
        val typeDef = prop(NodeEager).replaceable().nullable()
        val baseType = prop(NodeEager).replaceable().nullable()
        val nodes = prop { type(n.List.GT(NodeEager)).mutable() }
        val nodesInHierarchy = prop { type(n.List.GT(NodeEager)).mutable() }
        val maxChildDepth = propI().replaceable()
        val path = propS().replaceable()
    }

    object NodeValues : Values() {
        val nodeId = props().addItem(nodeIdEmpty)
        val values = prop { type(n.List.GT(Shared.ValueDate)).mutable() }
    }

    object Subscription : Values({ ifc(true) }) {
        val variableIds = op { ret(n.Collection.GT(milo.core.NodeId)) }
        val close = op().suspend()
    }

    object SubscriptionConfig : Values() {
        val samplingInterval = prop(n.Double).replaceable().value(200.0)
        val publishingInterval = prop(n.Double).replaceable().value(500.0)
        val queueSize = propI().replaceable().value(10)
        val deadbandType = prop(milo.core.DeadbandType).replaceable().value(milo.core.DeadbandType.None)
        val deadbandTypeValue = prop(n.Double).replaceable().value(0.0)
        val dataChangeTrigger = prop(milo.core.DataChangeTrigger).replaceable()
                .value(milo.core.DataChangeTrigger.StatusValueTimestamp)
    }


    object Synchronizer : Values() {
        val node = prop(Node)
        val externalCustomId = propS()
        val externalUniqueId = propS()
        val subscriptions = prop(n.List.GT(Subscription)).nullable().replaceable()
        val historicalSync = propDT().replaceable()
    }
}