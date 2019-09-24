package stdcodes

const (
	ReceiveErrorFromConfig                = 51
	ReceiveErrorOnGettingConfigFromConfig = 52

	InitializingDbError = 53

	ModuleRunFatalError  = 54
	ModuleManualShutdown = 56

	ConfigServiceSendConfigSchema = 61
	ConfigServiceSendRequirements = 62
	ConfigServiceSendModuleReady  = 63

	ConfigServiceReceiveConfiguration         = 64
	ConfigServiceReceiveRoutes                = 65
	ConfigServiceReceiveRequiredModuleAddress = 66

	ConfigServiceSendDataError         = 67
	ConfigServiceDisconnection         = 68
	ConfigServiceConnectionError       = 69
	RemoteConfigIsNotReceivedByTimeout = 70
)
