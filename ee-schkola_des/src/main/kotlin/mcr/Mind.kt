package mcr

import ee.design.Config
import ee.design.Controller
import ee.design.Module
import ee.design.op
import ee.lang.*

object Mind : Module({ initObjectTree() }) {

    object MindCommon : Controller({ ifc() }) {
        val info = op { ret(n.String) }
        val tenant = op { ret(n.String) }
        val close = op()
    }

    object Mind : Controller({ ifc().superUnits(MindConnector, MindReader, MindWriter) })

    object MindConnector : Controller({ ifc().superUnit(MindCommon) }) {
        val agent = op { ret(MindAgent) }
        val oauth = op { ret(MindOauth) }
    }

    object MindReader : Controller({ ifc().superUnit(MindCommon) }) {
        val timeSeries = op { ret(MindTimeSeriesReader) }
        val asset = op { ret(MindAssetReader) }
    }

    object MindWriter : Controller({ ifc().superUnit(MindCommon) }) {
        val timeSeriesWriter = op { ret(MindTimeSeriesWriter) }
        val exchange = op { ret(MindExchange) }
        val assetWriter = op { ret(MindAssetWriter) }
    }

    object MindAgent : Controller({ ifc().superUnit(MindCommon) }) {
        val register = op(p("iat")) { ret(ClientIdentifier) }
        val rotateKey = op(p("clientToken"), p("clientId")) { ret(ClientIdentifier) }
        val accessToken = op(p("jwtClientToken")) { ret(AccessToken) }
    }

    object MindOauth : Controller({ ifc().superUnit(MindCommon) }) {
        val accessToken = op(p("clientId"), p("clientSecret")) { ret(AccessToken) }
    }

    object MindExchange : Controller({ ifc().superUnit(MindCommon) }) {
        val exchange = op(p("content"), p("boundaryId")) { ret(n.String) }
    }

    object MindTimeSeriesWriter : Controller({ ifc().superUnit(MindCommon) }) {
        val createTimeSeries = op(p("assetId"), p("aspectName"), p("variableData", n.Map.GT(n.String, n.List.GT(Shared.ValueDate)))) { ret(n.Boolean) }
        val deleteTimeSeries = op(p("assetId"), p("aspectName"), p("range", Shared.FromTo)) { ret(n.Boolean) }
    }

    object MindTimeSeriesReader : Controller({ ifc().superUnit(MindCommon) }) {
        val timeSeries = op(p("assetId"), p("aspectName"), p("range", LimitQuerySelect))
    }

    object MindAssetWriter : Controller({ ifc().superUnit(MindCommon) }) {
        val createAsset = op(p("item", AssetDto)) { ret(AssetResource) }
        val createAssetType = op(p("id"), p("item", AssetTypeDto)) { ret(AssetTypeResource) }
        val createAspectType = op(p("id"), p("item", AspectTypeDto)) { ret(AspectTypeResource) }

        val updateAsset = op(p("id"), p("item", AssetDto), p("version")) { ret(AssetResource) }
        val updateAssetType = op(p("id"), p("item", AssetTypeDto), p("version")) { ret(AssetTypeResource) }
        val updateAspectType = op(p("id"), p("item", AspectTypeDto), p("version")) { ret(AspectTypeResource) }

        val deleteAsset = op(p("id"), p("version")) { ret(n.Boolean) }
        val deleteAssetType = op(p("id"), p("version")) { ret(n.Boolean) }
        val deleteAspectType = op(p("id"), p("version")) { ret(n.Boolean) }
    }

    object MindAssetReader : Controller({ ifc().superUnit(MindCommon) }) {
        val assets = op { ret(n.List.GT(AssetResource)) }
        val assetTypes = op { ret(n.List.GT(AssetTypeResource)) }
        val aspectTypes = op { ret(n.List.GT(AspectTypeResource)) }
        val assetRoot = op { ret(AssetResource) }

        val findAspectTypeById = op(AspectTypeResource.id) { retA(AspectTypeResource).nullable() }
    }


    object ProxyConfig : Values() {
        val host = propS()
        val port = propI()
        var useProxy = propB { replaceable().value(true) }
    }

    object TokenExpiration : Values() {
        val token = propS()
        val expiration = prop { type(n.Date) }
    }

    object ClientConfig : Values() {
        val id = propS { replaceable() }
        val secret = propS { replaceable() }
        val token = prop { type(TokenExpiration).replaceable() }
    }

    object AgentConfig : Values({ superUnit(ClientConfig) }) {
        val jwtToken = prop { type(TokenExpiration).replaceable() }
        val registrationToken = prop { type(TokenExpiration).replaceable() }
        val audience = propS { replaceable() }
        val schema = propS { replaceable() }
        val iat = propS { nullable() }
    }

    object MindConfig : Config() {
        val tenant = propS { replaceable() }
        val defaultEndpoint = propS {}
        val apiEndpoints = prop { type(n.Map) }
        val agent = prop { type(AgentConfig).replaceable().doc("MindconnectLib Agent") }
        val client = prop { type(ClientConfig).replaceable().doc("Technical User") }
        val proxy = prop { type(ProxyConfig).nullable() }
    }

    object ScopedRequest : Values() {
        val scope = prop { type(n.String).nullable() }
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

    object AspectData : Values() {
        val assetId = propS()
        val aspectName = propS()
        val aspectType = prop { type(AspectTypeResource) }
        val nodeIdData = prop { type(n.Map.GT(n.Any, VariableData)).mutable() }
    }

    object VariableData : Values() {
        val variable = prop { type(AspectVariable) }
        val data = prop { type(n.List.GT(Shared.ValueDate)).mutable() }
        val aspect = prop(AspectData)
    }

    //generated

    //Asset Management API

    object Aspect : Values() {
        val name = prop { type(n.String).nullable() }
        val variables = prop { type(n.List.GT(Variable)).nullable() }
    }

    object AspectLinks : Values() {
        val self = prop { type(Link).nullable() }
        val asset = prop { type(Link).nullable() }
        val aspectType = prop { type(Link).nullable() }
    }

    object AspectResource : Values() {
        val aspectTypeId = prop { type(n.String).doc("ID of the Aspect type") }
        val holderAssetId = prop { type(n.String).doc("Unique identifier") }
        val name = prop { type(n.String).doc("Name of the aspect") }
        val category = prop { type(Category).nullable() }
        val description = prop { type(n.String).nullable().doc("The description of the aspect") }
        val variables = prop { type(n.List.GT(AspectVariable)).nullable() }
        val links = prop { externalName("_links").type(AspectLinks).nullable() }
    }

    object AspectTypeDto : Values() {
        val name = prop { type(n.String).doc("Name of the aspect type. It has to be unique inside the tenant and cannot be changed later.") }
        val category = prop { type(Category).doc("If the aspect-type is used for static data or time-series. Cannot be changed once the aspect-type is created.") }
        val scope = prop { type(Scope).nullable().value(Scope.private).doc("Visibility of aspecttype. Setting this property to public makes it available to other tenants. Private types are only visible to the user's own tenant. Currently only private types can be created.") }
        val description = prop { type(n.String).nullable().doc("The description of the aspect type") }
        val variables = prop { type(n.List.GT(AspectVariable)).doc("Variables of the aspect-type. Variable names should be unique inside an aspect-type. Once a variable is added to the aspect that it cannot be renamed or removed.") }
    }

    object AspectTypeLinks : Values() {
        val self = prop { type(Link).nullable() }
    }

    object AspectTypeResource : Values({ superUnit(AspectTypeDto) }) {
        val id = prop { type(n.String) }
        val tenantId = prop { type(n.String).nullable() }
        val etag = prop { type(n.String).nullable() }
        val links = prop { externalName("_links").type(AspectTypeLinks).nullable() }
    }

    object AspectVariable : Values({ superUnit(VariableDefinition) }) {
        val qualityCode = prop { type(n.Boolean).nullable().doc("Indicates whether the variable has quality code. Cannot be changed.") }
    }

    object AssetDto : Values({ superUnit(AssetUpdateDto) }) {
        val typeId = prop { type(n.String).doc("""The type's id is a unique identifier. The id's length must be between 1 and 128 characters and matches the following symbols "A-Z", "a-z", "0-9", "_" and "." beginning with the tenant prefix what has a maximum of 8 characters. (e.g . ten_pref.type_id)""") }
        val t2Tenant = prop { type(n.String).nullable().doc("The id of the end-customer") }
        val parentId = prop { type(n.String).doc("Unique identifier") }
    }

    object AssetLinks : Values() {
        val self = prop { type(Link).nullable() }
        val parent = prop { type(Link).nullable() }
        val children = prop { type(Link).nullable() }
        val variables = prop { type(Link).nullable() }
        val aspects = prop { type(Link).nullable() }
        val t2Tenant = prop { type(Link).nullable() }
        val location = prop { type(Link).nullable() }
    }

    object AssetMoveDto : Values() {
        val newParentId = prop { type(n.String).doc("Unique identifier") }
    }

    object AssetResource : Values({ superUnit(AssetDto) }) {
        val tenantId = prop { type(n.String).nullable() }
        val assetId = prop { type(n.String).doc("Unique identifier") }
        val fileAssignments = prop { type(n.List.GT(FileAssignmentResource)).nullable() }
        val deleted = prop { type(n.Date).nullable() }
        val etag = prop { type(n.String).nullable() }
        val links = prop { externalName("_links").type(AssetLinks).nullable() }
    }

    object AssetTypeBase : Values() {
        val name = prop { type(n.String).doc("The type's name.") }
        val description = prop { type(n.String).nullable().doc("description") }
        val parentTypeId = prop { type(n.String) }
        val instantiable = prop { type(n.Boolean).nullable().doc("If instances can be created from this type. A non-instantiable type could be changed to be instantiable but not the other way around.") }
        val scope = prop { type(Scope).nullable().value(Scope.private).doc("Visibility of the assettype. Setting this property to public makes it available to other tenants. Private types are only visible to the user's own tenant. Currently only private types could be created.") }
        val variables = prop { type(n.List.GT(VariableDefinition)).nullable().doc("Direct variables of the asset-type. Variable names has to be unique inside the whole type-family (ancestors and descendants). Once added variables cannot be changed or removed.") }
    }

    object AssetTypeDto : Values({ superUnit(AssetTypeBase) }) {
        val aspects = prop { type(n.List.GT(NameAspectTypeId)).nullable().doc("Aspects of the asset-type. Once added aspects cannot be removed.") }
    }

    object AssetTypeLinks : Values() {
        val self = prop { type(Link).nullable() }
        val parent = prop { type(Link).nullable() }
    }

    object AssetTypeResource : Values({ superUnit(AssetTypeBase) }) {
        val id = prop { type(n.String) }
        val tenantId = prop { type(n.String).nullable() }
        val etag = prop { type(n.String).nullable() }
        val aspects = prop { type(n.List.GT(AspectTypeResource)).nullable() }
        val fileAssignments = prop { type(n.List.GT(FileAssignmentResource)).nullable() }
        val links = prop { externalName("_links").type(AssetTypeLinks).nullable() }
    }

    object AssetUpdateDto : Values() {
        val name = prop { type(n.String).doc("Name of the asset") }
        val externalId = prop { type(n.String).nullable().doc("The id given by the user") }
        val description = prop { type(n.String).nullable().doc("The description of the asset") }
        val location = prop { type(Location).nullable() }
        val variables = prop { type(n.List.GT(Variable)).nullable() }
        val aspects = prop { type(n.List.GT(Aspect)).nullable() }
    }

    object Category : EnumType() {
        val value = prop { type(n.String).key() }
        val dynamic = lit(value, "dynamic")
        val static = lit(value, "static")
    }

    object DataType : EnumType() {
        val value = prop { type(n.String).key() }
        val bOOLEAN = lit(value, "BOOLEAN")
        val iNT = lit(value, "INT")
        val lONG = lit(value, "LONG")
        val dOUBLE = lit(value, "DOUBLE")
        val sTRING = lit(value, "STRING")
        val tIMESTAMP = lit(value, "TIMESTAMP")
        val bigString = lit(value, "BIG_STRING")
    }

    object ErrorcodeLogrefMessage : Values() {
        val errorcode = prop { type(n.Double).nullable() }
        val logref = prop { type(n.String).nullable() }
        val message = prop { type(n.String).nullable() }
    }

    object FileAssignmentDto : Values() {
        val fileId = prop { type(n.String).nullable().doc("The id of the file to be assigned") }
    }

    object FileAssignmentLinks : Values() {
        val download = prop { type(Link).nullable() }
        val metadata = prop { type(Link).nullable() }
    }

    object FileAssignmentResource : Values() {
        val key = prop { type(n.String).nullable().doc("The key of assignment") }
        val fileId = prop { type(n.String).nullable().doc("The identifier of the file") }
        val links = prop { externalName("_links").type(FileAssignmentLinks).nullable() }
    }

    object FileMetadataResource : Values() {
        val id = prop { type(n.String).doc("Unique identifier") }
        val name = prop { type(n.String).nullable().doc("File name given by the user") }
        val originalFileName = prop { type(n.String).nullable().doc("Original filename of the file") }
        val description = prop { type(n.String).nullable().doc("File description") }
        val tenantId = prop { type(n.String).nullable() }
        val t2Tenant = prop { type(n.String).nullable().doc("The id of the end-customer") }
        val uploaded = prop { type(n.Date).nullable().doc("The time of the file upload") }
        val lastModified = prop { type(n.Date).nullable().doc("The time of the latest modification of the file") }
        val etag = prop { type(n.String).nullable() }
        val links = prop { externalName("_links").type(SelfDownload).nullable() }
    }

    object Link : Values() {
        val href = prop { type(n.String).nullable() }
    }

    object Location : Values() {
        val country = prop { type(n.String).nullable() }
        val region = prop { type(n.String).nullable().doc("County or other region code or name") }
        val locality = prop { type(n.String).nullable() }
        val streetAddress = prop { type(n.String).nullable() }
        val postalCode = prop { type(n.String).nullable() }
        val longitude = prop { type(n.Double).nullable().doc("The longitude part of the geographic coordinates") }
        val latitude = prop { type(n.Double).nullable().doc("The latitude part of the geographic coordinates") }
    }

    object NameAspectTypeId : Values() {
        val name = prop { type(n.String).nullable().doc("Name of the aspect. It has to be unique inside the type-family (ancestors and descendants).Cannot be changed. Reserved words (id, name, description, tenant, etag, scope, properties, propertySets, extends, variables, aspects, parentTypeId) cannot be used as aspect names.") }
        val aspectTypeId = prop { type(n.String).nullable() }
    }

    object Page : Values() {
        val size = prop { type(n.Int).nullable() }
        val totalElements = prop { type(n.Int).nullable() }
        val totalPages = prop { type(n.Int).nullable() }
        val number = prop { type(n.Int).nullable() }
    }

    object PagingLinks : Values() {
        val self = prop { type(Link).nullable() }
        val first = prop { type(Link).nullable() }
        val prev = prop { type(Link).nullable() }
        val next = prop { type(Link).nullable() }
        val last = prop { type(Link).nullable() }
    }

    object Scope : EnumType() {
        val value = prop { type(n.String).key() }
        val public = lit(value, "public")
        val private = lit(value, "private")
    }

    object SelfDownload : Values() {
        val self = prop { type(Link).nullable() }
        val download = prop { type(Link).nullable() }
    }

    object Variable : Values() {
        val name = prop { type(n.String).nullable() }
        val value = prop { type(n.String).nullable() }
    }

    object VariableDefinition : Values() {
        val name = prop { type(n.String).doc("Name of the variable. Once set cannot be changed. Reserved words (id, name, description, tenant, etag, scope, properties, propertySets, extends, variables, aspects, parentTypeId) cannot be used as variable names.") }
        val dataType = prop { type(DataType).doc("Data type of the variable. BIG_STRING could only be used by variables in dynamic aspect-types. Cannot be changed.") }
        val unit = prop { type(n.String).nullable().doc("Unit of measurement. Cannot be changed.") }
        val searchable = prop { type(n.Boolean).nullable().doc("Indicates whether sorting and filtering is allowed on this variable. Only usable for static properties. Cannot be changed.") }
        val length = prop { type(n.Int).nullable().doc("The max length of the variable's value. The length field is only used for variables of string or big_string dataType. Max length for string is 255 and max length for big_string 100000. Cannot be changed.") }
    }
//MindConnect API

    object Badrequest : Values() {
        val id = prop { type(n.String).nullable() }
        val message = prop { type(n.String).nullable() }
    }

    object Conflict : Values() {
        val id = prop { type(n.String).nullable() }
        val message = prop { type(n.String).nullable() }
    }

    object DataPointType : EnumType() {
        val value = prop { type(n.String).key() }
        val iNT = lit(value, "INT")
        val lONG = lit(value, "LONG")
        val dOUBLE = lit(value, "DOUBLE")
        val bOOLEAN = lit(value, "BOOLEAN")
        val sTRING = lit(value, "STRING")
    }

    object Direction : EnumType() {
        val value = prop { type(n.String).key() }
        val aSC = lit(value, "ASC")
        val dESC = lit(value, "DESC")
    }

    object Error : Values() {
        val id = prop { type(n.String).nullable() }
        val message = prop { type(n.String).nullable() }
    }

    object Forbidden : Values() {
        val id = prop { type(n.String).nullable() }
        val message = prop { type(n.String).nullable() }
    }

    object Mapping : Values() {
        val id = prop { type(n.String).nullable() }
        val agentId = prop { type(n.String) }
        val dataPointId = prop { type(n.String) }
        val dataPointUnit = prop { type(n.String).nullable() }
        val dataPointType = prop { type(DataPointType).nullable() }
        val entityId = prop { type(n.String) }
        val propertySetName = prop { type(n.String) }
        val propertyName = prop { type(n.String) }
        val propertyUnit = prop { type(n.String).nullable() }
        val propertyType = prop { type(PropertyType).nullable() }
        val qualityEnabled = prop { type(n.Boolean).nullable() }
    }

    object Notfound : Values() {
        val id = prop { type(n.String).nullable() }
        val message = prop { type(n.String).nullable() }
    }

    object NullHandling : EnumType() {
        val value = prop { type(n.String).key() }
        val nATIVE = lit(value, "NATIVE")
        val nullsFirst = lit(value, "NULLS_FIRST")
        val nullsLast = lit(value, "NULLS_LAST")
    }

    object Order : Values() {
        val direction = prop { type(Direction).nullable().doc("The order the property shall be sorted for.") }
        val property = prop { type(n.String).nullable().doc("The property to order for.") }
        val ignoreCase = prop { type(n.Boolean).nullable().doc("Whether or not the sort will be case sensitive.") }
        val nullHandling = prop { type(NullHandling).nullable() }
        val descending = prop { type(n.Boolean).nullable().doc("Whether sorting for this property shall be descending.") }
        val ascending = prop { type(n.Boolean).nullable().doc("Whether sorting for this property shall be ascending.") }
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
        val id = prop { type(n.String).nullable() }
        val message = prop { type(n.String).nullable() }
    }

    object PropertyType : EnumType() {
        val value = prop { type(n.String).key() }
        val iNT = lit(value, "INT")
        val lONG = lit(value, "LONG")
        val dOUBLE = lit(value, "DOUBLE")
        val bOOLEAN = lit(value, "BOOLEAN")
        val sTRING = lit(value, "STRING")
    }

    object Unauthorized : Values() {
        val id = prop { type(n.String).nullable() }
        val message = prop { type(n.String).nullable() }
    }
//Agent Management API

    object AccessToken : Values() {
        val accessToken = prop { externalName("access_token").type(n.String).nullable().doc("The access token to be used in calls to Mind with Bearer authentication scheme.") }
        val tokenType = prop { externalName("token_type").type(n.String).nullable().doc("The type of the access token issued.") }
        val expiresIn = prop { externalName("expires_in").type(n.Int).nullable().doc("Number of seconds before this token expires from the time of issuance.") }
        val jti = prop { type(n.String).nullable() }
        val scope = prop { type(n.String).nullable().doc("Agent permissions list.") }
    }

    object Agent : Values({ superUnit(AgentUpdate) }) {
        val id = prop { type(n.String).nullable().doc("Unique identifier of the agent") }
        val eTag = prop { type(n.String).nullable() }
        val entityId = prop { type(n.String).nullable().doc("Unique identifier of the entity") }
    }

    object AgentUpdate : Values() {
        val name = prop { type(n.String).doc("Name must be unique per tenant.") }
        val securityProfile = prop { type(SecurityProfile) }
    }

    object BadrequestIAM : Values() {
        val id = prop { type(n.String).nullable() }
        val error = prop { type(n.String).nullable() }
        val errorDescription = prop { externalName("error_description").type(n.String).nullable().doc("An error message with Correlation-ID value.") }
    }

    object ClientIdentifier : Values() {
        val clientId = prop { externalName("client_id").type(n.String) }
        val clientSecret = prop { externalName("client_secret").type(n.String).nullable().doc("Server generated client secret. Required if security profile is SHARED_SECRET.") }
        val clientSecretExpiresAt = prop { externalName("client_secret_expires_at").type(n.Int).nullable().doc("Epoch time in seconds which client secret expires at.") }
        val grantTypes = prop { externalName("grant_types").type(n.List.GT(GrantTypes)).nullable() }
        val registrationAccessToken = prop { externalName("registration_access_token").type(n.String).doc("The access token to be used at the client configuration endpoint to perform subsequent operations upon the client registration.") }
        val registrationClientUri = prop { externalName("registration_client_uri").type(n.String).doc("The fully qualified URL of the client configuration endpoint for this client.") }
        val tokenEndpointAuthMethod = prop { externalName("token_endpoint_auth_method").type(TokenEndpointAuthMethod).nullable().doc("The client authentication method.") }
    }

    object Configuration : Values() {
        val content = prop { type(OnboardingConfigurationContent).nullable() }
        val expiration = prop { type(n.Date).nullable() }
    }

    object DataPoint : Values() {
        val id = prop {
            type(n.String).doc("""Identifier of this data point. This id needs to be unique per data source configuration. Agents expected to upload timeseries value with this id, enabling backend services to match data with this data point.
This is NOT an auto generated field, enabling agents to specify it before uploading matching timeseries value.
""")
        }
        val name = prop { type(n.String) }
        val description = prop { type(n.String).nullable() }
        val type = prop { type(Type) }
        val unit = prop {
            type(n.String).doc("""Unit of data point. Can be empty.
""")
        }
        val customData = prop { type(n.Map).nullable().doc("A list of string tuples. Max 5 tuples allowed.") }
    }

    object DataSource : Values() {
        val name = prop { type(n.String) }
        val description = prop { type(n.String).nullable() }
        val dataPoints = prop { type(n.List.GT(DataPoint)) }
        val customData = prop { type(n.Map).nullable().doc("A list of string tuples. Max 5 tuples allowed.") }
    }

    object DataSourceConfiguration : Values() {
        val id = prop { type(n.String).nullable() }
        val eTag = prop { type(n.String).nullable() }
        val configurationId = prop { type(n.String).doc("Unique identifier of the datasource configuration.") }
        val dataSources = prop { type(n.List.GT(DataSource)) }
    }

    object GrantTypes : EnumType() {
        val value = prop { type(n.String).key() }
        val sharedSecret = lit(value, "SHARED_SECRET")
        val rsa3072 = lit(value, "RSA_3072")
        val clientCredentials = lit(value, "client_credentials")
    }

    object Jwks : Values() {
        val keys = prop { type(n.List.GT(Key)).nullable() }
    }

    object Key : Values() {
        val e = prop { type(ee.lang.n.String) }
        val n = prop { type(ee.lang.n.String) }
        val kty = prop { type(ee.lang.n.String) }
        val kid = prop { type(ee.lang.n.String) }
    }

    object Keys : Values() {
        val jwks = prop { type(Jwks).nullable() }
    }

    object OnboardingConfigurationContent : Values() {
        val baseUrl = prop { type(n.String).nullable() }
        val iat = prop { type(n.String).nullable() }
        val clientCredentialProfile = prop { type(n.List.GT(n.String)).nullable() }
        val clientId = prop { type(n.String).nullable() }
        val tenant = prop { type(n.String).nullable() }
    }

    object OnboardingStatus : Values() {
        val status = prop { type(Status).nullable() }
    }

    object OnlineStatus : Values() {
        val status = prop { type(Status).nullable() }
        val since = prop { type(n.Date).nullable() }
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
        val id = prop { type(n.String).nullable() }
        val message = prop { type(n.String).nullable() }
    }

    object RotationKeys : Values() {
        val clientId = prop { externalName("client_id").type(n.String).nullable() }
        val jwks = prop { type(Jwks).nullable() }
    }

    object SecurityProfile : EnumType() {
        val value = prop { type(n.String).key() }
        val sharedSecret = lit(value, "SHARED_SECRET")
        val rsa3072 = lit(value, "RSA_3072")
    }

    object Status : EnumType() {
        val value = prop { type(n.String).key() }
        val oNLINE = lit(value, "ONLINE")
        val oFFLINE = lit(value, "OFFLINE")
    }

    object TokenEndpointAuthMethod : EnumType() {
        val value = prop { type(n.String).key() }
        val clientSecretJwt = lit(value, "client_secret_jwt")
        val privateKeyJwt = lit(value, "private_key_jwt")
    }

    object TokenKey : Values({ superUnit(Key) }) {
        val alg = prop { type(n.String).nullable() }
        val use = prop { type(n.String).nullable() }
        val value = prop { type(n.String).nullable() }
    }

    object Type : EnumType() {
        val value = prop { type(n.String).key() }
        val iNT = lit(value, "INT")
        val lONG = lit(value, "LONG")
        val dOUBLE = lit(value, "DOUBLE")
        val bOOLEAN = lit(value, "BOOLEAN")
        val sTRING = lit(value, "STRING")
    }

    object UnauthorizedIAM : Values() {
        val id = prop { type(n.String).nullable() }
        val error = prop { type(n.String).nullable() }
        val errorDescription = prop { externalName("error_description").type(n.String).nullable().doc("An error message with Correlation-ID value.") }
    }
//IoT Time Series API

    object Timeseries : Values() {
        val time = prop { externalName("_time").type(n.Date).doc("Timestamp of the data points.") }
        val exampleproperty0 = prop { type(n.String).nullable().doc("Property as specified in the property set type configuration") }
        val exampleproperty0Qc = prop { externalName("exampleproperty0_qc").type(n.Int).nullable().doc("Quality code for a property as specified in the property set type configuration") }
        val exampleproperty1 = prop { type(n.Int).nullable().doc("Property as specified in the property set type configuration") }
    }
//IoT File API

    object File : Values() {
        val name = prop { type(n.String).nullable() }
        val path = prop { type(n.String).nullable() }
        val type = prop { type(n.String).nullable() }
        val size = prop { type(n.Int).nullable() }
        val timestamp = prop { type(n.Date).nullable() }
        val created = prop { type(n.Date).nullable() }
        val updated = prop { type(n.Date).nullable() }
        val createdBy = prop { type(n.String).nullable() }
        val updatedBy = prop { type(n.String).nullable() }
        val description = prop { type(n.String).nullable() }
        val eTag = prop { type(n.Int).nullable() }
    }

    object Notmodified : Values() {
        val id = prop { type(n.String).nullable() }
        val message = prop { type(n.String).nullable() }
    }
//IoT TS Aggregates API

    object Aggregate : Values() {
        val countgood = prop { type(n.Int).nullable().doc("Number of good records in the aggregated interval.") }
        val countuncertain = prop { type(n.Int).nullable().doc("Number of uncertain records in the aggregated interval.") }
        val countbad = prop { type(n.Int).nullable().doc("Number of bad records in the aggregated interval.") }
        val sum = prop { type(n.Double).nullable().doc("Sum of the values within the interval.") }
        val avg = prop { type(n.Double).nullable().doc("Average of the values within the interval.") }
        val mintime = prop { type(n.Date).nullable().doc("Timestamp of the minimum value within the interval.") }
        val minvalue = prop { type(n.Double).nullable().doc("Minimum value within the interval.") }
        val maxtime = prop { type(n.Date).nullable().doc("Timestamp of the maximum value within the interval.") }
        val maxvalue = prop { type(n.Double).nullable().doc("Maximum value within the interval.") }
        val firsttime = prop { type(n.Date).nullable().doc("Timestamp of the first value within the interval.") }
        val firstvalue = prop { type(n.Double).nullable().doc("First value within the interval.") }
        val lasttime = prop { type(n.Date).nullable().doc("Timestamp of the last measurement within the interval.") }
        val lastvalue = prop { type(n.Double).nullable().doc("Last value within the interval.") }
    }

    object Aggregates : Values() {
        val starttime = prop { type(n.Date).nullable().doc("start time of the interval (exclusive)") }
        val endtime = prop { type(n.Date).nullable().doc("end time of the interval (inclusive)") }
        val property1 = prop { type(Aggregate).nullable() }
        val property2 = prop { type(Aggregate).nullable() }
    }
//OAuth Authorization Server

    object AccessTokenResponse : Values() {
        val accessToken = prop { externalName("access_token").type(n.String).nullable() }
        val tokenType = prop { externalName("token_type").type(n.String).nullable() }
        val expiresIn = prop { externalName("expires_in").type(n.Int).nullable() }
        val scope = prop { type(n.String).nullable() }
        val jti = prop { type(n.String).nullable() }
    }

    object JsonWebKeyIncl0 : Values() {
        val kty = prop { type(ee.lang.n.String).nullable() }
        val e = prop { type(ee.lang.n.String).nullable() }
        val use = prop { type(ee.lang.n.String).nullable() }
        val kid = prop { type(ee.lang.n.String).nullable() }
        val n = prop { type(ee.lang.n.String).nullable() }
    }

    object TokenKeysResponse : Values() {
        val keys = prop { type(n.List.GT(JsonWebKeyIncl0)).nullable() }
    }
}