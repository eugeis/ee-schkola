package mcr

import ee.design.*
import ee.lang.*

object OpcUa : Module({artifact("mcr-opcua").initObjectTree()  }) {

    val nodeIdEmpty = p("nodeId", milo.core.NodeId).value(NodeIdEmpty)
    val nodeIdEmptyAsId = p("id", milo.core.NodeId).value(NodeIdEmpty)
    val nodeIdEmptyAsTypeId = p("typeId", milo.core.NodeId).value(NodeIdEmpty)

    object OpcUa : Controller({ ifc().doc("The main factory for all OPC UA APIs") }) {
        val info = op { ret(n.String) }
        val browser = op { ret(OpcUaBrowser) }
       // val reader = op { ret(OpcUaReader) }
        val browseReader = op { ret(OpcUaBrowseReader) }
        val subscriber = op { ret(OpcUaSubscriber) }
        val close = op()
    }

    object OpcUaBrowser : Controller({ ifc().nonBlocking().doc("The browser OPC API") }) {
        val root = op {
            ret(Node)
        }
        val findNode = op(p("nodeId", milo.core.NodeId)) {
            ret { type(Node).nullable() }
        }
        val findNodes = op(p("nodeIds", n.Collection.GT(milo.core.NodeId))) {
            ret(n.Map.GT(milo.core.NodeId, Node))
        }
        val findNodesRecursive = op(p("nodeId", milo.core.NodeId), p("depth", n.Int).nullable()) {
            ret(NodeSet)
        }
        val findNodesRecursiveAll = op(p("nodeId", milo.core.NodeId)) {
            ret(NodeSet)
        }
    }
/*
    object OpcUaReader : Controller({ ifc().nonBlocking() }) {
        val readValue = op(p("variableNodeId", milo.core.NodeId)) {
            ret { type(Shared.ValueDate).nullable() }
        }

        val readValues = op(p("variableNodeIds", n.Collection.GT(milo.core.NodeId))) {
            ret(n.Map.GT(milo.core.NodeId, Shared.ValueDate))
        }

        val readValuesTo = op(p("variableNodeIds", n.Collection.GT(milo.core.NodeId)),
                p("destination", lambda(p("nodeId", milo.core.NodeId),
                        p("value", n.Any), p("time", n.Date))))

        val readHistoricalValues = op(p("variableNodeIds", n.Collection.GT(milo.core.NodeId)),
                p("range", Shared.FromTo)) {
            ret(n.Map.GT(milo.core.NodeId, n.List.GT(Shared.ValueDate)))
        }
        val readHistoricalValuesTo = op(p("variableNodeIds", n.Collection.GT(milo.core.NodeId)),
                p("range", Shared.FromTo),
                p("destination", lambda(p("nodeId", milo.core.NodeId),
                        p("valueDates", n.List.GT(Shared.ValueDate)))))
    }
*/
    object OpcUaSubscriber : Controller({ ifc().nonBlocking() }) {
        val subscribe = op(p("variableNodeId", milo.core.NodeId), p("config", SubscriptionConfig),
                p("callback", lambda(p("nodeId", milo.core.NodeId),
                        p("value", n.Any), p("time", n.Date)))) {
            ret(Subscription)
        }

        val subscribeAll = op(p("variableNodeIds", n.Collection.GT(milo.core.NodeId)),
                p("config", SubscriptionConfig),
                p("callback", lambda(p("nodeId", milo.core.NodeId),
                        p("value", n.Any), p("time", n.Date)))) {
            ret(n.List.GT(Subscription))
        }
    }

    object OpcUaBrowseReader : Controller({ ifc().nonBlocking() }) {
        val findValuesByRootPath = op(p("paths", n.Collection),
                p("namespaceMapping", n.Map.GT(n.String, n.Int))) {
            ret(n.Map.GT(n.String, n.Any))
        }

        val findValuesByPath = op(p("startingNode", milo.core.NodeId), p("paths", n.Collection),
                p("namespaceMapping", n.Map.GT(n.String, n.Int))) {
            ret(n.Map.GT(n.String, n.Any))
        }

        val findDataValuesByPath = op(p("startingNode", milo.core.NodeId), p("paths", n.Collection),
                p("namespaceMapping", n.Map.GT(n.String, n.Int))) {
            ret(n.Map.GT(n.String, milo.core.DataValue))
        }

        val findNodeIdsByRootPath = op(p("paths", n.Collection),
                p("namespaceMapping", n.Map.GT(n.String, n.Int))) {
            ret(n.Map.GT(n.String, milo.core.NodeId))
        }

        val findNodeIdsByPath = op(p("startingNode", milo.core.NodeId), p("paths", n.Collection),
                p("namespaceMapping", n.Map.GT(n.String, n.Int))) {
            ret(n.Map.GT(n.String, milo.core.NodeId))
        }
    }

    object OpcUaNamespacesClient : Controller({ ifc().nonBlocking() }) {
        val findValuesByRootPath = op(p("paths", n.Collection)) {
            ret(n.Map.GT(n.String, n.Any))
        }

        val findValuesByPath = op(p("startingNodeId", milo.core.NodeId), p("paths", n.Collection)) {
            ret(n.Map.GT(n.String, n.Any))
        }

        val findDataValuesByPath = op(p("startingNodeId", milo.core.NodeId), p("paths", n.Collection)) {
            ret(n.Map.GT(n.String, milo.core.DataValue))
        }
    }

    object IndexUri : Basic({
        doc("Data type for OPC UA namespace definition")
    }) {
        val index = propI()
        val uri = propS()
    }

    val maxValue2M: Int = 2 * 1024 * 1024
    val maxValue20M: Int = 20 * 1024 * 1024

    object OpcUaConfig : Basic() {
        val uri = propS { value("opc.tcp://localhost:4840") }.doc("DEPRECATED")
        val timeout = 1.minutes.doc("used for connection and " +
                "all other requests like read, subscription")
        val maxTimeout = 2.minutes.doc("in case if the timeout reached," +
                " the reconnection process doubles the current timeout but only until the max timeout")
        val maxBrowseDescriptionsInRequest = propI().replaceable().value(5000)
        val maxReadValuesInRequest = propI().replaceable().value(5000)
        val maxItemsInSubscription = propI().replaceable().value(10000)
        val acknowledgeTimeout = propI().replaceable().value(10000)

        val keepAliveFailuresAllowed = propI().replaceable().value(3)
        val keepAliveInterval = propI().replaceable().value(5000)
        val keepAliveTimeout = propI().replaceable().value(5000)

        val lookupByIp = propB().replaceable().value(false).doc("try to lookup by ip if endpoint can't be found")
        val reconnectByErrors = propB().replaceable().value(false).doc("try to reconnect by errors")

        val encodingMaxArrayLength = propI().replaceable().value(maxValue2M)
        val encodingMaxStringLength = propI().replaceable().value(maxValue20M)
        val encodingMaxRecursionDepth = propI().replaceable().value(128)
    }

    //sync config
    object OpcUaSync : Config() {
        val nodeSets = prop(n.List).doc("NodeSet file names")
        val instanceCreations = prop(n.List.GT(InstanceCreation))
        val endpoints = prop(n.List.GT(EndpointSync))
        val opcUaConfigs = prop(n.Map.GT(n.String, OpcUaConfig))
        val subscriptionConfigs = prop(n.Map.GT(n.String, SubscriptionConfig))
        val nodesSyncs = prop(n.Map.GT(n.String, n.List.GT(NodeSync)))
    }

    object EndpointSync : Config() {
        val endpointUri = propS().value("opc.tcp://localhost:4840")
        val nodeSyncs = prop(n.List.GT(NodeSync)).nullable()
        val opcUaConfigRef = propS().nullable()
        val subscriptionConfigRef = propS().nullable()
    }

    object NodeSync : Basic() {
        val sourcePath = propS()
        val targetPath = propS()
        val clone = propB().value(true)
        val keepId = propB().value(false)
        val recursive = prop(Recursive).nullable()
        val nodeSyncRefs = prop(n.List).nullable()
    }

    object Recursive : Basic() {
        val enabled = propB().value(true)
        val depth = propI().nullable()
        val keepIdPaths = prop(n.List.GT(n.String)).nullable()
        val excludePaths = prop(n.List.GT(n.String)).nullable()
        val clone = propB().value(true)
        val keepId = propB().value(false)
    }

    object InstanceCreation : Basic() {
        val typePath = propS()
        val instanceNames = prop(n.List.GT(n.String))
        val targetPath = propS().value("DeviceSet")
    }

    //sync resolved types
    object NodeSyncMeta : Basic() {
        val id = props().addItem(nodeIdEmptyAsId)
        val name = propS()
        val type = prop(n.Class.GTStar()).replaceable()
        val path = propS().doc("optional, path of the node")
        val valueRank = propI().nullable()
        val arrayDimensions = prop(n.List.GT(milo.core.UInteger)).nullable()
    }

    object SourceTargetSync : Basic({doc("source and target sync")}) {
        val source = prop(NodeSyncMeta)
        val target = prop(NodeSyncMeta)
    }

    object SyncPoint : Basic() {
        val uri = propS()
        val namespaces = prop(n.List.GT(IndexUri))
        val namespaceMapping = prop(n.Map.GT(milo.core.UShort, milo.core.UShort))
        val idPrefix = propS()
    }

    //model config
    object UrlRange : Values() {
        val uri = propS().replaceable(true).value("opc.tcp://localhost")
        val startPort = propI().replaceable(true).value(6001)
        val serverCount = propI().replaceable(true).value(10)
        val portOffset = propI().replaceable(true).value(1)
        val enabled = propB().replaceable(true).value(false)
    }

    object IdNameType : Basic() {
        val id = props().addItem(nodeIdEmptyAsId)
        val name = propS()
        val type = prop(milo.core.NodeClass)
    }

    object Ref : Basic() {
        val me = prop(IdNameType).doc("target id")
        val referenceTypeId = prop(milo.core.NodeId).value("NodeIdEmpty")
    }

    object RefEager : Values() {
        val node = prop(NodeEager)
        val ref = prop(Ref)
    }

    object Node : Entity() {
        val me = prop(IdNameType)
        val displayName = propS().nullable()
        val description = propS().nullable()
        val writeMask = prop { type(milo.core.UInteger).nullable() }
        val userWriteMask = prop { type(milo.core.UInteger).nullable() }
        val dataType = prop(IdNameType).nullable()
        val typeDef = prop(IdNameType).nullable()
        val parent = prop(IdNameType).nullable()
        val nodes = prop { type(n.List.GT(IdNameType)) }
        val inverses = prop { type(n.List.GT(Ref)) }
        val forwards = prop { type(n.List.GT(Ref)) }
    }

    object TypeNode : Entity({ superUnit(Node) }) {
        val subType = prop(IdNameType).nullable()
        val abstract = propB().nullable()
    }

    object DataTypeNode : Entity({ superUnit(TypeNode) })

    object MethodNode : Entity({ superUnit(Node) }) {
        val executable = propB().nullable()
        val userExecutable = propB().nullable()
    }

        object ViewNode : Entity({ superUnit(Node) }) {
            val containsNoLoops = propB().nullable()
            val eventNotifier = prop(milo.core.UByte).nullable()
        }

        object ObjectNode : Entity({ superUnit(Node) }) {
            val eventNotifier = prop(milo.core.UByte).nullable()
        }

        object VariableTypeNode : Entity({ superUnit(TypeNode) }) {
            val value = prop(milo.core.DataValue).nullable()
            //val dataType = prop(IdNameType).nullable()
            val valueRank = propI().nullable()
            val arrayDimensions = prop(n.List.GT(milo.core.UInteger)).nullable()
            val accessLevel = prop { type(milo.core.UByte).nullable() }
            val userAccessLevel = prop { type(milo.core.UByte).nullable() }
        }

        object ObjectTypeNode : Entity({ superUnit(TypeNode) })

        object ReferenceTypeNode : Entity({ superUnit(Node) }) {
            val abstract = propB().nullable()
            val symmetric = propB().nullable()
            val inverseName = propS().nullable()
        }

        object VariableNode : Entity({ superUnit(Node) }) {
            val value = prop(milo.core.DataValue).nullable()
            //val dataType = prop(IdNameType).nullable()
            val valueRank = propI().nullable()
            val arrayDimensions = prop(n.List.GT(milo.core.UInteger)).nullable()
            val accessLevel = prop { type(milo.core.UByte).nullable() }
            val userAccessLevel = prop { type(milo.core.UByte).nullable() }
            val minimumSamplingInterval = prop(n.Double).nullable()
            val historizing = propB().nullable()
        }

        object NodeEager : Entity() {
            val node = prop(Node)
            val parent = prop(NodeEager).replaceable().nullable()
            val dataType = prop(NodeEager).replaceable().nullable()
            val typeDef = prop(NodeEager).replaceable().nullable()
            val baseType = prop(NodeEager).replaceable().nullable()
            val nodes = prop { type(n.List.GT(NodeEager)).mutable() }
            val hierarchy = prop { type(n.List.GT(NodeEager)).mutable() }
            val maxChildDepth = propI().replaceable()
            val path = propS().replaceable()
            val inverses = prop { type(n.List.GT(RefEager)).mutable() }
            val forwards = prop { type(n.List.GT(RefEager)).mutable() }
        }

        object NodeSet : Values() {
            val startingNode = prop(Node)
            val nodes = prop(n.Map.GT(milo.core.NodeId, Node))
            val namespaces = prop(n.List.GT(IndexUri))
        }

        object NodeValues : Values() {
            val nodeId = props().addItem(nodeIdEmpty)
            //val values = prop { type(n.List.GT(Shared.ValueDate)).mutable() }
        }

        object Subscription : Values({ ifc(true) }) {
            val variableIds = op { ret(n.Collection.GT(milo.core.NodeId)) }
            val close = op().nonBlocking()
        }

    object SubscriptionConfig : Basic({doc("Data type for OPC UA subscription configration")}) {

        val samplingInterval = prop(n.Double).replaceable().value(200.0)
        val publishingInterval = prop(n.Double).replaceable().value(500.0)
        val queueSize = propI().replaceable().value(10)
        val deadbandType = prop(milo.core.DeadbandType).replaceable().value(milo.core.DeadbandType.None)
        val deadbandTypeValue = prop(n.Double).replaceable().value(0.0)
        val dataChangeTrigger = prop(milo.core.DataChangeTrigger).replaceable()
                .value(milo.core.DataChangeTrigger.StatusValueTimestamp)
    }

    object OpcUaLimits : Basic() {
        val maxSessionCount = propI().replaceable().value(550)
        val minSupportedSampleRate = prop(n.Double).replaceable().value(200.0)
        val maxSupportedSampleRate = prop(n.Double).replaceable().value(86400000.0)//24 hours in millis
        val maxBrowseContinuationPoints = prop(n.Short).replaceable().value(Short.MAX_VALUE)
        val maxQueryContinuationPoints = prop(n.Short).replaceable().value(Short.MAX_VALUE)
        val maxHistoryContinuationPoints = prop(n.Short).replaceable().value(550)
        val maxArrayLength = prop(n.Int).replaceable().value(maxValue2M)
        val maxStringLength = prop(n.Int).replaceable().value(maxValue20M)
        val maxNodesPerRead = prop(n.Int).replaceable().value(maxValue2M)
        val maxNodesPerHistoryReadData = prop(n.Int).replaceable().value(maxValue2M)
        val maxNodesPerHistoryReadEvents = prop(n.Int).replaceable().value(maxValue2M)
        val maxNodesPerWrite = prop(n.Int).replaceable().value(maxValue2M)
        val maxNodesPerHistoryUpdateData = prop(n.Int).replaceable().value(maxValue2M)
        val maxNodesPerHistoryUpdateEvents = prop(n.Int).replaceable().value(maxValue2M)
        val maxNodesPerMethodCall = prop(n.Int).replaceable().value(maxValue2M)
        val maxNodesPerBrowse = prop(n.Int).replaceable().value(maxValue2M)
        val maxNodesPerRegisterNodes = prop(n.Int).replaceable().value(maxValue2M)
        val maxNodesPerTranslateBrowsePathsToNodeIds = prop(n.Int).replaceable().value(maxValue2M)
        val maxNodesPerNodeManagement = prop(n.Int).replaceable().value(maxValue2M)
        val maxMonitoredItemsPerCall = prop(n.Int).replaceable().value(maxValue2M)
    }

    object Synchronizer : Values() {
        val node = prop(Node)
        val externalCustomId = propS()
        val externalUniqueId = propS()
        val subscriptions = prop(n.List.GT(Subscription)).nullable().replaceable()
        val historicalSync = propDT().replaceable()
    }

}

val NodeIdEmpty = ExternalType { name("NodeIdEmpty").namespace("mcr.opcua") }