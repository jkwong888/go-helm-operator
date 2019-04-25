package v1alpha1

type LibertyAppAutoscalingSpec struct {
	TargetCPUUtilizationPercentage *int  `json:"targetCPUUtilizationPercentage,omitempty" yaml:"targetCPUUtilizationPercentage,omitempty"`
	Enabled                        *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	MinReplicas                    *int  `json:"minReplicas,omitempty" yaml:"minReplicas,omitempty"`
	MaxReplicas                    *int  `json:"maxReplicas,omitempty" yaml:"maxReplicas,omitempty"`
}
type LibertyAppResourcesConstraintsSpec struct {
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
type LibertyAppResourcesRequestsSpec struct {
	Cpu    *string `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory *string `json:"memory,omitempty" yaml:"memory,omitempty"`
}
type LibertyAppResourcesLimitsSpec struct {
	Cpu    *string `json:"cpu,omitempty" yaml:"cpu,omitempty"`
	Memory *string `json:"memory,omitempty" yaml:"memory,omitempty"`
}
type LibertyAppResourcesSpec struct {
	Constraints *LibertyAppResourcesConstraintsSpec `json:"constraints,omitempty" yaml:"constraints,omitempty"`
	Requests    *LibertyAppResourcesRequestsSpec    `json:"requests,omitempty" yaml:"requests,omitempty"`
	Limits      *LibertyAppResourcesLimitsSpec      `json:"limits,omitempty" yaml:"limits,omitempty"`
}
type LibertyAppRbacSpec struct {
	Install *bool `json:"install,omitempty" yaml:"install,omitempty"`
}
type LibertyAppImageSpec struct {
	Lifecycle                    *map[string]string   `json:"lifecycle,omitempty" yaml:"lifecycle,omitempty"`
	License                      *string              `json:"license,omitempty" yaml:"license,omitempty"`
	ServerOverridesConfigMapName *string              `json:"serverOverridesConfigMapName,omitempty" yaml:"serverOverridesConfigMapName,omitempty"`
	Security                     *map[string]string   `json:"security,omitempty" yaml:"security,omitempty"`
	Repository                   *string              `json:"repository,omitempty" yaml:"repository,omitempty"`
	PullPolicy                   *string              `json:"pullPolicy,omitempty" yaml:"pullPolicy,omitempty"`
	LivenessProbe                *map[string]string   `json:"livenessProbe,omitempty" yaml:"livenessProbe,omitempty"`
	ExtraEnvs                    *[]map[string]string `json:"extraEnvs,omitempty" yaml:"extraEnvs,omitempty"`
	ExtraVolumeMounts            *[]map[string]string `json:"extraVolumeMounts,omitempty" yaml:"extraVolumeMounts,omitempty"`
	Tag                          *string              `json:"tag,omitempty" yaml:"tag,omitempty"`
	ReadinessProbe               *map[string]string   `json:"readinessProbe,omitempty" yaml:"readinessProbe,omitempty"`
	PullSecret                   *string              `json:"pullSecret,omitempty" yaml:"pullSecret,omitempty"`
}
type LibertyAppPodSpec struct {
	ExtraInitContainers *[]map[string]string `json:"extraInitContainers,omitempty" yaml:"extraInitContainers,omitempty"`
	ExtraContainers     *[]map[string]string `json:"extraContainers,omitempty" yaml:"extraContainers,omitempty"`
	ExtraVolumes        *[]map[string]string `json:"extraVolumes,omitempty" yaml:"extraVolumes,omitempty"`
	Annotations         *map[string]string   `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Security            *map[string]string   `json:"security,omitempty" yaml:"security,omitempty"`
	Labels              *map[string]string   `json:"labels,omitempty" yaml:"labels,omitempty"`
}
type LibertyAppServiceSpec struct {
	Type           *string              `json:"type,omitempty" yaml:"type,omitempty"`
	ExtraSelectors *map[string]string   `json:"extraSelectors,omitempty" yaml:"extraSelectors,omitempty"`
	Name           *string              `json:"name,omitempty" yaml:"name,omitempty"`
	TargetPort     *int                 `json:"targetPort,omitempty" yaml:"targetPort,omitempty"`
	Port           *int                 `json:"port,omitempty" yaml:"port,omitempty"`
	ExtraPorts     *[]map[string]string `json:"extraPorts,omitempty" yaml:"extraPorts,omitempty"`
	Enabled        *bool                `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Annotations    *map[string]string   `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Labels         *map[string]string   `json:"labels,omitempty" yaml:"labels,omitempty"`
}
type LibertyAppJmsServiceSpec struct {
	Enabled    *bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Port       *int    `json:"port,omitempty" yaml:"port,omitempty"`
	TargetPort *int    `json:"targetPort,omitempty" yaml:"targetPort,omitempty"`
	Type       *string `json:"type,omitempty" yaml:"type,omitempty"`
}
type LibertyAppSslSpec struct {
	Enabled                       *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	UseClusterSSLConfiguration    *bool `json:"useClusterSSLConfiguration,omitempty" yaml:"useClusterSSLConfiguration,omitempty"`
	CreateClusterSSLConfiguration *bool `json:"createClusterSSLConfiguration,omitempty" yaml:"createClusterSSLConfiguration,omitempty"`
}
type LibertyAppEnvSpec struct {
	JvmArgs *string `json:"jvmArgs,omitempty" yaml:"jvmArgs,omitempty"`
}
type LibertyAppArchSpec struct {
	Ppc64le *string `json:"ppc64le,omitempty" yaml:"ppc64le,omitempty"`
	S390x   *string `json:"s390x,omitempty" yaml:"s390x,omitempty"`
	Amd64   *string `json:"amd64,omitempty" yaml:"amd64,omitempty"`
}
type LibertyAppSessioncacheHazelcastImageSpec struct {
	PullPolicy *string `json:"pullPolicy,omitempty" yaml:"pullPolicy,omitempty"`
	Repository *string `json:"repository,omitempty" yaml:"repository,omitempty"`
	Tag        *string `json:"tag,omitempty" yaml:"tag,omitempty"`
}
type LibertyAppSessioncacheHazelcastSpec struct {
	Enabled  *bool                                     `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	Embedded *bool                                     `json:"embedded,omitempty" yaml:"embedded,omitempty"`
	Image    *LibertyAppSessioncacheHazelcastImageSpec `json:"image,omitempty" yaml:"image,omitempty"`
}
type LibertyAppSessioncacheSpec struct {
	Hazelcast *LibertyAppSessioncacheHazelcastSpec `json:"hazelcast,omitempty" yaml:"hazelcast,omitempty"`
}
type LibertyAppOidcClientSpec struct {
	Enabled          *bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	ClientId         *string `json:"clientId,omitempty" yaml:"clientId,omitempty"`
	ClientSecretName *string `json:"clientSecretName,omitempty" yaml:"clientSecretName,omitempty"`
	DiscoveryURL     *string `json:"discoveryURL,omitempty" yaml:"discoveryURL,omitempty"`
}
type LibertyAppDeploymentSpec struct {
	Labels      *map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Annotations *map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
}
type LibertyAppIiopServiceSpec struct {
	SecurePort          *int    `json:"securePort,omitempty" yaml:"securePort,omitempty"`
	SecureTargetPort    *int    `json:"secureTargetPort,omitempty" yaml:"secureTargetPort,omitempty"`
	Type                *string `json:"type,omitempty" yaml:"type,omitempty"`
	Enabled             *bool   `json:"enabled,omitempty" yaml:"enabled,omitempty"`
	NonSecurePort       *int    `json:"nonSecurePort,omitempty" yaml:"nonSecurePort,omitempty"`
	NonSecureTargetPort *int    `json:"nonSecureTargetPort,omitempty" yaml:"nonSecureTargetPort,omitempty"`
}
type LibertyAppMicroprofileHealthSpec struct {
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
type LibertyAppMicroprofileSpec struct {
	Health *LibertyAppMicroprofileHealthSpec `json:"health,omitempty" yaml:"health,omitempty"`
}
type LibertyAppMonitoringSpec struct {
	Enabled *bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
type LibertyAppIngressSpec struct {
	RewriteTarget *string            `json:"rewriteTarget,omitempty" yaml:"rewriteTarget,omitempty"`
	Path          *string            `json:"path,omitempty" yaml:"path,omitempty"`
	Host          *string            `json:"host,omitempty" yaml:"host,omitempty"`
	SecretName    *string            `json:"secretName,omitempty" yaml:"secretName,omitempty"`
	Annotations   *map[string]string `json:"annotations,omitempty" yaml:"annotations,omitempty"`
	Labels        *map[string]string `json:"labels,omitempty" yaml:"labels,omitempty"`
	Enabled       *bool              `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
type LibertyAppPersistenceSelectorSpec struct {
	Label *string `json:"label,omitempty" yaml:"label,omitempty"`
	Value *string `json:"value,omitempty" yaml:"value,omitempty"`
}
type LibertyAppPersistenceSpec struct {
	Name                   *string                            `json:"name,omitempty" yaml:"name,omitempty"`
	Size                   *string                            `json:"size,omitempty" yaml:"size,omitempty"`
	UseDynamicProvisioning *bool                              `json:"useDynamicProvisioning,omitempty" yaml:"useDynamicProvisioning,omitempty"`
	StorageClassName       *string                            `json:"storageClassName,omitempty" yaml:"storageClassName,omitempty"`
	Selector               *LibertyAppPersistenceSelectorSpec `json:"selector,omitempty" yaml:"selector,omitempty"`
}
type LibertyAppLogsSpec struct {
	PersistTransactionLogs *bool   `json:"persistTransactionLogs,omitempty" yaml:"persistTransactionLogs,omitempty"`
	ConsoleFormat          *string `json:"consoleFormat,omitempty" yaml:"consoleFormat,omitempty"`
	ConsoleLogLevel        *string `json:"consoleLogLevel,omitempty" yaml:"consoleLogLevel,omitempty"`
	ConsoleSource          *string `json:"consoleSource,omitempty" yaml:"consoleSource,omitempty"`
	PersistLogs            *bool   `json:"persistLogs,omitempty" yaml:"persistLogs,omitempty"`
}
type LibertyAppSpec struct {
	ResourceNameOverride *string                     `json:"resourceNameOverride,omitempty" yaml:"resourceNameOverride,omitempty"`
	Pod                  *LibertyAppPodSpec          `json:"pod,omitempty" yaml:"pod,omitempty"`
	Service              *LibertyAppServiceSpec      `json:"service,omitempty" yaml:"service,omitempty"`
	IiopService          *LibertyAppIiopServiceSpec  `json:"iiopService,omitempty" yaml:"iiopService,omitempty"`
	Microprofile         *LibertyAppMicroprofileSpec `json:"microprofile,omitempty" yaml:"microprofile,omitempty"`
	Ingress              *LibertyAppIngressSpec      `json:"ingress,omitempty" yaml:"ingress,omitempty"`
	Image                *LibertyAppImageSpec        `json:"image,omitempty" yaml:"image,omitempty"`
	Env                  *LibertyAppEnvSpec          `json:"env,omitempty" yaml:"env,omitempty"`
	Deployment           *LibertyAppDeploymentSpec   `json:"deployment,omitempty" yaml:"deployment,omitempty"`
	Autoscaling          *LibertyAppAutoscalingSpec  `json:"autoscaling,omitempty" yaml:"autoscaling,omitempty"`
	Resources            *LibertyAppResourcesSpec    `json:"resources,omitempty" yaml:"resources,omitempty"`
	Arch                 *LibertyAppArchSpec         `json:"arch,omitempty" yaml:"arch,omitempty"`
	Sessioncache         *LibertyAppSessioncacheSpec `json:"sessioncache,omitempty" yaml:"sessioncache,omitempty"`
	OidcClient           *LibertyAppOidcClientSpec   `json:"oidcClient,omitempty" yaml:"oidcClient,omitempty"`
	Monitoring           *LibertyAppMonitoringSpec   `json:"monitoring,omitempty" yaml:"monitoring,omitempty"`
	ReplicaCount         *int                        `json:"replicaCount,omitempty" yaml:"replicaCount,omitempty"`
	Rbac                 *LibertyAppRbacSpec         `json:"rbac,omitempty" yaml:"rbac,omitempty"`
	JmsService           *LibertyAppJmsServiceSpec   `json:"jmsService,omitempty" yaml:"jmsService,omitempty"`
	Ssl                  *LibertyAppSslSpec          `json:"ssl,omitempty" yaml:"ssl,omitempty"`
	Persistence          *LibertyAppPersistenceSpec  `json:"persistence,omitempty" yaml:"persistence,omitempty"`
	Logs                 *LibertyAppLogsSpec         `json:"logs,omitempty" yaml:"logs,omitempty"`
}
