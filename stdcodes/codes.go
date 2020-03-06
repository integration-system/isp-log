package stdcodes

const (
	ModuleReadLocalConfigError      = 31
	ModuleInvalidLocalConfig        = 32
	ModuleInvalidRemoteConfig       = 33
	ModuleOverrideRemoteConfigError = 34
	ModuleRunFatalError             = 35
	ModuleManualShutdown            = 36
	ModuleDefaultRCReadError        = 37

	ModuleInternalGrpcServiceError  = 40
	ModuleGrpcServiceStart          = 41
	ModuleGrpcServiceStartError     = 42
	ModuleGrpcServiceManualShutdown = 43
	ModuleInternalHttpServiceError  = 44
	ModuleMetricServiceError        = 45

	ReceiveErrorFromConfig                = 51
	ReceiveErrorOnGettingConfigFromConfig = 52

	InitializingDbError = 53

	InitializingRabbitMqError = 54
	RabbitMqClientError       = 55
	RabbitMqBlockedConnection = 56

	ConfigServiceSendConfigSchema = 71
	ConfigServiceSendRequirements = 72
	ConfigServiceSendModuleReady  = 73

	ConfigServiceReceiveConfiguration         = 74
	ConfigServiceReceiveRoutes                = 75
	ConfigServiceReceiveRequiredModuleAddress = 76

	ConfigServiceSendDataError         = 77
	ConfigServiceConnection            = 86
	ConfigServiceDisconnection         = 78
	ConfigServiceConnectionError       = 79
	RemoteConfigIsNotReceivedByTimeout = 80
	ConfigServiceInvalidDataReceived   = 81

	ProfilingStart          = 82
	ProfilingStop           = 83
	ProfilingSendNewProfile = 84
	ProfilingError          = 85
)
