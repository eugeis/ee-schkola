package com.siemens.mcr

import ee.design.*
import ee.lang.*


object Mcr : Comp({ artifact("mcr").namespace("com.siemens.mcr") }) {
    object Shared : Module() {
        object ValueDate : Basic() {
            val value = prop { type(n.Any) }
            val date = prop { type(n.Date) }
        }

        object NameValueDate : Basic() {
            val name = propS()
            val value = prop { type(n.Any) }
            val date = prop { type(n.Date) }
        }

        object FromTo : Values() {
            val from = prop { type(n.Date) }
            val to = prop { type(n.Date) }
        }
    }

    object Gateway : Module() {
        object SyncConfig : Config() {
            val opcUa = prop { type(Opcua.OpcUaConfig) }
            val syncNodePath = propS { value("Objects/DeviceSet") }
            val subscribePublishingInterval = prop { type(n.Double).value("2000.0") }

            val assetId = propS()
            val location = prop { type(Mindsphere.Location) }
            val idPrefix = propS({ value("opcUa") })
            val typeNamePrefix = propS()
            val baseAssetTypeId = propS()
        }
    }

    object Opcua : Module() {

        object OpcUaConfig : Values() {
            val uri = propS()
            val timeout = 10.seconds.doc("used for connection and " +
                    "all other requests like read, subscription")
            val maxTimeout = 1.minutes.doc("in case if the timeout reached," +
                    " the reconnection process doubles the current timeout but only until the max timeout")
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

        object IdNameType : Basic() {
            val id = propS()
            val name = propS()
            val type = prop(NodeType)
        }

        object Node : Entity() {
            val me = prop(IdNameType)
            val description = propS()
            val parent = prop(IdNameType).nullable(true)
            val dataType = prop(IdNameType).nullable(true)
            val nodes = prop { type(n.List.GT(IdNameType)) }
        }

        object Subscription : Values({ ifc(true) }) {
            val close = op()
        }

        object Synchronizer : Values() {
            val node = prop(Node)
            val externalCustomId = propS()
            val externalUniqueId = propS()
            val subscription = prop(Subscription).nullable(true).replaceable(true)
            val historicalSync = propDT().replaceable(true)
        }
    }


    object Mindsphere : Module() {

        object ProxyConfig : Values() {
            val host = propS()
            val port = propI()
            var useProxy = propB { replaceable(true).value(true) }
        }

        object TokenExpiration : Values() {
            val token = propS()
            val expiration = prop { type(n.Date) }
        }

        object ClientConfig : Values() {
            val id = propS { replaceable(true) }
            val secret = propS { replaceable(true) }
            val token = prop { type(TokenExpiration).replaceable(true) }
        }

        object AgentConfig : Values({ superUnit(ClientConfig) }) {
            val jwtToken = prop { type(TokenExpiration).replaceable(true) }
            val registrationToken = prop { type(TokenExpiration).replaceable(true) }
            val audience = propS { replaceable(true) }
            val schema = propS { replaceable(true) }
            val iat = propS { nullable(true) }
        }

        object MindSphereConfig : Config() {
            val tenant = propS { replaceable(true) }
            val defaultEndpoint = propS {}
            val apiEndpoints = prop { type(n.Map) }
            val agent = prop { type(AgentConfig).replaceable(true).doc("MindconnectLib Agent") }
            val client = prop { type(ClientConfig).replaceable(true).doc("Technical User") }
            val proxy = prop { type(ProxyConfig).nullable(true) }
        }

        object ScopedRequest : Values() {
            val scope = prop { type(n.String).isNullable() }
        }

        object OAuthTokenRequest : Basic({ superUnit(ScopedRequest) }) {
            val grantType = prop {
                type(GrantTypes).externalName("grant_type").value(GrantTypes.clientCredentials).doc(
                        "The type of authentication being used to obtain the token")
            }
            val clientAssertionType = prop {
                type(n.String).externalName("client_assertion_type").value(
                        "urn:ietf:params:oauth:client-assertion-type:jwt-bearer")
            }
            val clientAssertion = prop {
                type(n.String).externalName("client_assertion").doc(
                        "Token is signed (by the client) with keys (which is provided at '/register' | '/register/<client_id>' endpoint) or client_secret (which is provided by '/register' | '/register/<client_id>' endpoint) in jwt format.")
            }
        }

        object Errors : Values() {
            val errors = prop { type(n.List.GT(ErrorcodeLogrefMessage)) }
        }

        object LimitQuerySelect : Basic({ superUnit(Shared.FromTo) }) {
            val limit = propI()
            val query = propS()
            val select = propS()
        }

        object AssetIdAspectData : Values() {
            val assetId = propS()
            val aspectName = propS()
            val aspectType = prop { type(AspectTypeResource) }
            val nodeIdData = prop { type(n.Map.GT(n.String, VariableData)).mutable(true) }
        }

        object VariableData : Values() {
            val variable = prop { type(AspectVariable) }
            val data = prop { type(n.List.GT(Shared.ValueDate)).mutable(true) }
            val aspect = prop(AssetIdAspectData)
        }

        //generated

        //Asset Management API

        object Aspect : Values() {
            val name = prop { type(n.String).nullable(true) }
            val variables = prop { type(n.List.GT(Variable)).nullable(true) }
        }

        object AspectLinks : Values() {
            val self = prop { type(Link).nullable(true) }
            val asset = prop { type(Link).nullable(true) }
            val aspectType = prop { type(Link).nullable(true) }
        }

        object AspectResource : Values() {
            val aspectTypeId = prop { type(n.String).doc("ID of the Aspect type") }
            val holderAssetId = prop { type(n.String).doc("Unique identifier") }
            val name = prop { type(n.String).doc("Name of the aspect") }
            val category = prop { type(Category).nullable(true) }
            val description = prop { type(n.String).nullable(true).doc("The description of the aspect") }
            val variables = prop { type(n.List.GT(AspectVariable)).nullable(true) }
            val links = prop { externalName("_links").type(AspectLinks).nullable(true) }
        }

        object AspectTypeDto : Values() {
            val name = prop { type(n.String).doc("Name of the aspect type") }
            val category = prop { type(Category) }
            val scope = prop { type(Scope).nullable(true).value(Scope.private).doc("Visibility of aspecttype. Setting this property to public makes it available to other tenants. Private types are only isVisible to the user's own tenant. Currently only private types can be created.") }
            val description = prop { type(n.String).nullable(true).doc("The description of the aspect type") }
            val variables = prop { type(n.List.GT(AspectVariable)) }
        }

        object AspectTypeLinks : Values() {
            val self = prop { type(Link).nullable(true) }
        }

        object AspectTypeResource : Values({ superUnit(AspectTypeDto) }) {
            val id = prop { type(n.String) }
            val tenantId = prop { type(n.String).nullable(true) }
            val etag = prop { type(n.String).nullable(true) }
            val links = prop { externalName("_links").type(AspectTypeLinks).nullable(true) }
        }

        object AspectVariable : Values({ superUnit(VariableDefinition) }) {
            val qualityCode = prop { type(n.Boolean).nullable(true).doc("Indicates whether the variable has quality code") }
        }

        object AssetDto : Values({ superUnit(AssetUpdateDto) }) {
            val typeId = prop { type(n.String).doc("""The type's id is a unique identifier. The id's length must be between 1 and 128 characters and matches the following symbols "A-Z", "a-z", "0-9", "_" and "." beginning with the tenant prefix what has a maximum of 8 characters. (e.g . ten_pref.type_id)""") }
            val t2Tenant = prop { type(n.String).nullable(true).doc("The id of the end-customer") }
            val parentId = prop { type(n.String).doc("Unique identifier") }
        }

        object AssetLinks : Values() {
            val self = prop { type(Link).nullable(true) }
            val parent = prop { type(Link).nullable(true) }
            val children = prop { type(Link).nullable(true) }
            val variables = prop { type(Link).nullable(true) }
            val aspects = prop { type(Link).nullable(true) }
            val t2Tenant = prop { type(Link).nullable(true) }
            val location = prop { type(Link).nullable(true) }
        }

        object AssetMoveDto : Values() {
            val newParentId = prop { type(n.String).doc("Unique identifier") }
        }

        object AssetResource : Values({ superUnit(AssetDto) }) {
            val tenantId = prop { type(n.String).nullable(true) }
            val assetId = prop { type(n.String).doc("Unique identifier") }
            val deleted = prop { type(n.Date).nullable(true) }
            val etag = prop { type(n.String).nullable(true) }
            val links = prop { externalName("_links").type(AssetLinks).nullable(true) }
        }

        object AssetTypeBase : Values() {
            val name = prop { type(n.String).doc("The type's name.") }
            val description = prop { type(n.String).nullable(true).doc("description") }
            val parentTypeId = prop { type(n.String) }
            val scope = prop { type(Scope).nullable(true).value(Scope.private).doc("Visibility of the assettype. Setting this property to public makes it available to other tenants. Private types are only isVisible to the user's own tenant.") }
            val variables = prop { type(n.List.GT(VariableDefinition)).nullable(true) }
            val instantiable = prop { type(n.Boolean).nullable(true).doc("The type's name.") }
        }

        object AssetTypeDto : Values({ superUnit(AssetTypeBase) }) {
            val aspects = prop { type(n.List.GT(NameAspectTypeId)).nullable(true) }
        }

        object AssetTypeLinks : Values() {
            val self = prop { type(Link).nullable(true) }
            val parent = prop { type(Link).nullable(true) }
        }

        object AssetTypeResource : Values({ superUnit(AssetTypeBase) }) {
            val id = prop { type(n.String) }
            val tenantId = prop { type(n.String).nullable(true) }
            val etag = prop { type(n.String).nullable(true) }
            val aspects = prop { type(n.List.GT(AspectTypeResource)).nullable(true) }
            val links = prop { externalName("_links").type(AssetTypeLinks).nullable(true) }
        }

        object AssetUpdateDto : Values() {
            val name = prop { type(n.String).doc("Name of the asset") }
            val externalId = prop { type(n.String).nullable(true).doc("The id given by the user") }
            val description = prop { type(n.String).nullable(true).doc("The description of the asset") }
            val location = prop { type(Location).nullable(true) }
            val variables = prop { type(n.List.GT(Variable)).nullable(true) }
            val aspects = prop { type(n.List.GT(Aspect)).nullable(true) }
        }

        object Category : EnumType() {
            val value = prop { type(n.String).key(true) }
            val dynamic = lit(value, "dynamic")
            val static = lit(value, "static")
        }

        object DataType : EnumType() {
            val value = prop { type(n.String).key(true) }
            val bOOLEAN = lit(value, "BOOLEAN")
            val iNT = lit(value, "INT")
            val lONG = lit(value, "LONG")
            val dOUBLE = lit(value, "DOUBLE")
            val sTRING = lit(value, "STRING")
            val tIMESTAMP = lit(value, "TIMESTAMP")
            val bigString = lit(value, "BIG_STRING")

        }

        object ErrorcodeLogrefMessage : Values() {
            val errorcode = prop { type(n.Double).nullable(true) }
            val logref = prop { type(n.String).nullable(true) }
            val message = prop { type(n.String).nullable(true) }
        }

        object Link : Values() {
            val href = prop { type(n.String).nullable(true) }
        }

        object Location : Values() {
            val country = prop { type(n.String).nullable(true) }
            val region = prop { type(n.String).nullable(true).doc("County or other region code or name") }
            val locality = prop { type(n.String).nullable(true) }
            val streetAddress = prop { type(n.String).nullable(true) }
            val postalcode = prop { type(n.String).nullable(true) }
            val longitude = prop { type(n.Double).nullable(true).doc("The longitude part of the geographic coordinates") }
            val latitude = prop { type(n.Double).nullable(true).doc("The latitude part of the geographic coordinates") }
        }

        object NameAspectTypeId : Values() {
            val name = prop { type(n.String).nullable(true) }
            val aspectTypeId = prop { type(n.String).nullable(true) }
        }

        object Page : Values() {
            val size = prop { type(n.Int).nullable(true) }
            val totalElements = prop { type(n.Int).nullable(true) }
            val totalPages = prop { type(n.Int).nullable(true) }
            val number = prop { type(n.Int).nullable(true) }
        }

        object PagingLinks : Values() {
            val self = prop { type(Link).nullable(true) }
            val first = prop { type(Link).nullable(true) }
            val prev = prop { type(Link).nullable(true) }
            val next = prop { type(Link).nullable(true) }
            val last = prop { type(Link).nullable(true) }
        }

        object Scope : EnumType() {
            val value = prop { type(n.String).key(true) }
            val public = lit(value, "public")
            val private = lit(value, "private")
        }

        object Variable : Values() {
            val name = prop { type(n.String).nullable(true) }
            val value = prop { type(n.String).replaceable(true).nullable(true) }
        }

        object VariableDefinition : Values() {
            val name = prop { type(n.String).doc("Name of the variable") }
            val dataType = prop { type(DataType).doc("Data type of the variable") }
            val unit = prop { type(n.String).nullable(true).doc("Unit of measurement") }
            val searchable = prop { type(n.Boolean).nullable(true).doc("Indicates whether sorting and filtering is allowed on this variable. Only usable for static properties.") }
            val length = prop { type(n.Int).nullable(true).doc("The max length of the variable's value. The length field is only used for variables of string dataType?.") }
        }
//MindConnect API

        object Badrequest : Values() {
            val id = prop { type(n.String).nullable(true) }
            val message = prop { type(n.String).nullable(true) }
        }

        object Conflict : Values() {
            val id = prop { type(n.String).nullable(true) }
            val message = prop { type(n.String).nullable(true) }
        }

        object DataPointType : EnumType() {
            val value = prop { type(n.String).key(true) }
            val iNT = lit(value, "INT")
            val lONG = lit(value, "LONG")
            val dOUBLE = lit(value, "DOUBLE")
            val bOOLEAN = lit(value, "BOOLEAN")
            val sTRING = lit(value, "STRING")
        }

        object Direction : EnumType() {
            val value = prop { type(n.String).key(true) }
            val aSC = lit(value, "ASC")
            val dESC = lit(value, "DESC")
        }

        object Error : Values() {
            val id = prop { type(n.String).nullable(true) }
            val message = prop { type(n.String).nullable(true) }
        }

        object Forbidden : Values() {
            val id = prop { type(n.String).nullable(true) }
            val message = prop { type(n.String).nullable(true) }
        }

        object Mapping : Values() {
            val id = prop { type(n.String).nullable(true) }
            val agentId = prop { type(n.String) }
            val dataPointId = prop { type(n.String) }
            val dataPointUnit = prop { type(n.String).nullable(true) }
            val dataPointType = prop { type(DataPointType).nullable(true) }
            val entityId = prop { type(n.String) }
            val propertySetName = prop { type(n.String) }
            val propertyName = prop { type(n.String) }
            val propertyUnit = prop { type(n.String).nullable(true) }
            val propertyType = prop { type(PropertyType).nullable(true) }
            val qualityEnabled = prop { type(n.Boolean).nullable(true) }
        }

        object Notfound : Values() {
            val id = prop { type(n.String).nullable(true) }
            val message = prop { type(n.String).nullable(true) }
        }

        object NullHandling : EnumType() {
            val value = prop { type(n.String).key(true) }
            val nATIVE = lit(value, "NATIVE")
            val nullsFirst = lit(value, "NULLS_FIRST")
            val nullsLast = lit(value, "NULLS_LAST")
        }

        object Order : Values() {
            val direction = prop { type(Direction).nullable(true).doc("The order the property shall be sorted for.") }
            val property = prop { type(n.String).nullable(true).doc("The property to order for.") }
            val ignoreCase = prop { type(n.Boolean).nullable(true).doc("Whether or not the sort will be case sensitive.") }
            val nullHandling = prop { type(NullHandling).nullable(true) }
            val descending = prop { type(n.Boolean).nullable(true).doc("Whether sorting for this property shall be descending.") }
            val ascending = prop { type(n.Boolean).nullable(true).doc("Whether sorting for this property shall be ascending.") }
        }

        object PagedMapping : Values() {
            val content = prop { type(n.List.GT(Mapping)) }
            val last = prop { type(n.Boolean).doc("Whether the current item is the last one.") }
            val totalPages = prop { type(n.Int).doc("The number of total pages.") }
            val totalElements = prop { type(n.Int).doc("The total amount of elements.") }
            val numberOfElements = prop { type(n.Int).doc("The number of elements currently on this page.") }
            val first = prop { type(n.Boolean).doc("Whether the current item is the first one.") }
            val sort = prop { type(n.List.GT(Order)).doc("The sorting parameters for the page.") }
            val size = prop { type(n.Int).doc("The size of the page.") }
            val number = prop { type(n.Int).doc("The number of the current item.") }
        }

        object PayLoadTooLarge : Values() {
            val id = prop { type(n.String).nullable(true) }
            val message = prop { type(n.String).nullable(true) }
        }

        object PropertyType : EnumType() {
            val value = prop { type(n.String).key(true) }
            val iNT = lit(value, "INT")
            val lONG = lit(value, "LONG")
            val dOUBLE = lit(value, "DOUBLE")
            val bOOLEAN = lit(value, "BOOLEAN")
            val sTRING = lit(value, "STRING")
        }

        object Unauthorized : Values() {
            val id = prop { type(n.String).nullable(true) }
            val message = prop { type(n.String).nullable(true) }
        }
//Agent Management API

        object AccessToken : Values() {
            val accessToken = prop { externalName("access_token").type(n.String).nullable(true).doc("The access token to be used in calls to MindSphereCommon with Bearer authentication scheme.") }
            val tokenType = prop { externalName("token_type").type(n.String).nullable(true).doc("The type of the access token issued.") }
            val expiresIn = prop { externalName("expires_in").type(n.Int).nullable(true).doc("Number of seconds before this token expires from the time of issuance.") }
            val jti = prop { type(n.String).nullable(true) }
            val scope = prop { type(n.String).nullable(true).doc("Agent permissions list.") }
        }

        object Agent : Values({ superUnit(AgentUpdate) }) {
            val id = prop { type(n.String).nullable(true).doc("Unique identifier of the agent") }
            val eTag = prop { type(n.String).nullable(true) }
            val entityId = prop { type(n.String).nullable(true).doc("Unique identifier of the entity") }
        }

        object AgentUpdate : Values() {
            val name = prop { type(n.String).doc("Name must be unique per tenant.") }
            val securityProfile = prop { type(SecurityProfile) }
        }

        object BadrequestIAM : Values() {
            val id = prop { type(n.String).nullable(true) }
            val error = prop { type(n.String).nullable(true) }
            val errorDescription = prop { externalName("error_description").type(n.String).nullable(true).doc("An error message with Correlation-ID value.") }
        }

        object ClientIdentifier : Values() {
            val clientId = prop { externalName("client_id").type(n.String) }
            val clientSecret = prop { externalName("client_secret").type(n.String).nullable(true).doc("Server generated client secret. Required if security profile is SHARED_SECRET.") }
            val clientSecretExpiresAt = prop { externalName("client_secret_expires_at").type(n.Int).nullable(true).doc("Epoch time in seconds which client secret expires at.") }
            val grantTypes = prop { externalName("grant_types").type(n.List.GT(GrantTypes)).nullable(true) }
            val registrationAccessToken = prop { externalName("registration_access_token").type(n.String).doc("The access token to be used at the client configuration endpoint to perform subsequent operations upon the client registration.") }
            val registrationClientUri = prop { externalName("registration_client_uri").type(n.String).doc("The fully qualified URL of the client configuration endpoint for this client.") }
            val tokenEndpointAuthMethod = prop { externalName("token_endpoint_auth_method").type(TokenEndpointAuthMethod).nullable(true).doc("The client authentication method.") }
        }

        object Configuration : Values() {
            val content = prop { type(OnboardingConfigurationContent).nullable(true) }
            val expiration = prop { type(n.Date).nullable(true) }
        }

        object DataPoint : Values() {
            val id = prop {
                type(n.String).doc("""Identifier of this data point. This id needs to be unique per data source configuration. Agents expected to upload timeseries value with this id, enabling backend services to match data with this data point.
This is NOT an auto generated field, enabling agents to specify it before uploading matching timeseries value.
""")
            }
            val name = prop { type(n.String) }
            val description = prop { type(n.String).nullable(true) }
            val type = prop { type(Type) }
            val unit = prop {
                type(n.String).doc("""Unit of data point. Can be empty.
""")
            }
            val customData = prop { type(n.Map).nullable(true).doc("A list of string tuples. Max 5 tuples allowed.") }
        }

        object DataSource : Values() {
            val name = prop { type(n.String) }
            val description = prop { type(n.String).nullable(true) }
            val dataPoints = prop { type(n.List.GT(DataPoint)) }
            val customData = prop { type(n.Map).nullable(true).doc("A list of string tuples. Max 5 tuples allowed.") }
        }

        object DataSourceConfiguration : Values() {
            val id = prop { type(n.String).nullable(true) }
            val eTag = prop { type(n.String).nullable(true) }
            val configurationId = prop { type(n.String).doc("Unique identifier of the datasource configuration.") }
            val dataSources = prop { type(n.List.GT(DataSource)) }
        }

        object GrantTypes : EnumType() {
            val value = prop { type(n.String).key(true) }
            val sharedSecret = lit(value, "SHARED_SECRET")
            val rsa3072 = lit(value, "RSA_3072")
            val clientCredentials = lit(value, "client_credentials")
        }

        object Jwks : Values() {
            val keys = prop { type(n.List.GT(Key)).nullable(true) }
        }

        object Key : Values() {
            val e = prop { type(ee.lang.n.String) }
            val n = prop { type(ee.lang.n.String) }
            val kty = prop { type(ee.lang.n.String) }
            val kid = prop { type(ee.lang.n.String) }
        }

        object Keys : Values() {
            val jwks = prop { type(Jwks).nullable(true) }
        }

        object OnboardingConfigurationContent : Values() {
            val baseUrl = prop { type(n.String).nullable(true) }
            val iat = prop { type(n.String).nullable(true) }
            val clientCredentialProfile = prop { type(n.List.GT(n.String)).nullable(true) }
            val clientId = prop { type(n.String).nullable(true) }
            val tenant = prop { type(n.String).nullable(true) }
        }

        object OnboardingStatus : Values() {
            val status = prop { type(Status).nullable(true) }
        }

        object PagedAgent : Values() {
            val content = prop { type(n.List.GT(Agent)) }
            val last = prop { type(n.Boolean).doc("Whether the current item is the last one.") }
            val totalPages = prop { type(n.Int).doc("The number of total pages.") }
            val totalElements = prop { type(n.Int).doc("The total amount of elements.") }
            val numberOfElements = prop { type(n.Int).doc("The number of elements currently on this page.") }
            val first = prop { type(n.Boolean).doc("Whether the current item is the first one.") }
            val sort = prop { type(n.List.GT(Order)).doc("The sorting parameters for the page.") }
            val size = prop { type(n.Int).doc("The size of the page.") }
            val number = prop { type(n.Int).doc("The number of the current item.") }
        }

        object Preconditionfailed : Values() {
            val id = prop { type(n.String).nullable(true) }
            val message = prop { type(n.String).nullable(true) }
        }

        object RotationKeys : Values() {
            val clientId = prop { externalName("client_id").type(n.String).nullable(true) }
            val jwks = prop { type(Jwks).nullable(true) }
        }

        object SecurityProfile : EnumType() {
            val value = prop { type(n.String).key(true) }
            val sharedSecret = lit(value, "SHARED_SECRET")
            val rsa3072 = lit(value, "RSA_3072")
        }

        object Status : EnumType() {
            val value = prop { type(n.String).key(true) }
            val notOnboarded = lit(value, "NOT_ONBOARDED")
            val oNBOARDING = lit(value, "ONBOARDING")
            val oNBOARDED = lit(value, "ONBOARDED")
        }

        object TokenEndpointAuthMethod : EnumType() {
            val value = prop { type(n.String).key(true) }
            val clientSecretJwt = lit(value, "client_secret_jwt")
            val privateKeyJwt = lit(value, "private_key_jwt")
        }

        object TokenKey : Values({ superUnit(Key) }) {
            val alg = prop { type(n.String).nullable(true) }
            val use = prop { type(n.String).nullable(true) }
            val value = prop { type(n.String).nullable(true) }
        }

        object Type : EnumType() {
            val value = prop { type(n.String).key(true) }
            val iNT = lit(value, "INT")
            val lONG = lit(value, "LONG")
            val dOUBLE = lit(value, "DOUBLE")
            val bOOLEAN = lit(value, "BOOLEAN")
            val sTRING = lit(value, "STRING")
        }

        object UnauthorizedIAM : Values() {
            val id = prop { type(n.String).nullable(true) }
            val error = prop { type(n.String).nullable(true) }
            val errorDescription = prop { externalName("error_description").type(n.String).nullable(true).doc("An error message with Correlation-ID value.") }
        }
//IoT Time Series API

        object Timeseries : Values() {
            val time = prop { externalName("_time").type(n.Date).doc("Timestamp of the data points.") }
            val exampleproperty0 = prop { type(n.String).nullable(true).doc("Property as specified in the property set type configuration") }
            val exampleproperty0Qc = prop { externalName("exampleproperty0_qc").type(n.Int).nullable(true).doc("Quality code for a property as specified in the property set type configuration") }
            val exampleproperty1 = prop { type(n.Int).nullable(true).doc("Property as specified in the property set type configuration") }
        }
//OAuth Authorization Server

        object AccessTokenResponse : Values() {
            val accessToken = prop { externalName("access_token").type(n.String).nullable(true) }
            val tokenType = prop { externalName("token_type").type(n.String).nullable(true) }
            val expiresIn = prop { externalName("expires_in").type(n.Int).nullable(true) }
            val scope = prop { type(n.String).nullable(true) }
            val jti = prop { type(n.String).nullable(true) }
        }

        object JsonWebKeyIncl0 : Values() {
            val kty = prop { type(ee.lang.n.String).nullable(true) }
            val e = prop { type(ee.lang.n.String).nullable(true) }
            val use = prop { type(ee.lang.n.String).nullable(true) }
            val kid = prop { type(ee.lang.n.String).nullable(true) }
            val n = prop { type(ee.lang.n.String).nullable(true) }
        }

        object TokenKeysResponse : Values() {
            val keys = prop { type(n.List.GT(JsonWebKeyIncl0)).nullable(true) }
        }
    }
}
