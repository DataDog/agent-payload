package process

import (
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestServiceDiscovery_WithResources_RoundTrip(t *testing.T) {
	original := &ServiceDiscovery{
		GeneratedServiceName: &ServiceName{
			Name:   "web-server",
			Source: ServiceNameSource_SERVICE_NAME_SOURCE_UNKNOWN,
		},
		DdServiceName: &ServiceName{
			Name:   "web-app",
			Source: ServiceNameSource_SERVICE_NAME_SOURCE_KUBERNETES,
		},
		TracerMetadata: []*TracerMetadata{
			{
				Language:    "go",
				RuntimeId:   "runtime-123",
				ServiceName: "nginx-traced",
			},
		},
		ApmInstrumentation: true,
		Resources: []*Resource{
			{
				Resource: &Resource_Logs{
					Logs: &LogResource{
						Path: "/var/log/nginx/access.log",
					},
				},
			},
			{
				Resource: &Resource_Logs{
					Logs: &LogResource{
						Path: "/var/log/nginx/error.log",
					},
				},
			},
			{
				Resource: &Resource_Logs{
					Logs: &LogResource{
						Path: "/var/log/application.log",
					},
				},
			},
		},
	}

	data, err := proto.Marshal(original)
	require.NoError(t, err, "Failed to marshal ServiceDiscovery")

	decoded := &ServiceDiscovery{}
	err = proto.Unmarshal(data, decoded)
	require.NoError(t, err, "Failed to unmarshal ServiceDiscovery")

	assert.Equal(t, original.GeneratedServiceName.Name, decoded.GeneratedServiceName.Name)
	assert.Equal(t, original.GeneratedServiceName.Source, decoded.GeneratedServiceName.Source)
	assert.Equal(t, original.DdServiceName.Name, decoded.DdServiceName.Name)
	assert.Equal(t, original.ApmInstrumentation, decoded.ApmInstrumentation)

	require.Len(t, decoded.Resources, 3, "Should have 3 resources")

	expectedPaths := []string{
		"/var/log/nginx/access.log",
		"/var/log/nginx/error.log",
		"/var/log/application.log",
	}
	
	actualPaths := make([]string, len(decoded.Resources))
	for i, resource := range decoded.Resources {
		logs := resource.GetLogs()
		require.NotNil(t, logs)
		actualPaths[i] = logs.Path
	}
	assert.ElementsMatch(t, expectedPaths, actualPaths)
}

func TestLogResource_SinglePath(t *testing.T) {
	resource := &Resource{
		Resource: &Resource_Logs{
			Logs: &LogResource{
				Path: "/var/log/test.log",
			},
		},
	}

	sd := &ServiceDiscovery{
		GeneratedServiceName: &ServiceName{
			Name:   "test-service",
			Source: ServiceNameSource_SERVICE_NAME_SOURCE_UNKNOWN,
		},
		Resources: []*Resource{resource},
	}

	data, err := proto.Marshal(sd)
	require.NoError(t, err)

	decoded := &ServiceDiscovery{}
	err = proto.Unmarshal(data, decoded)
	require.NoError(t, err)

	require.Len(t, decoded.Resources, 1)
	logs := decoded.Resources[0].GetLogs()
	require.NotNil(t, logs)
	assert.Equal(t, "/var/log/test.log", logs.Path)
}

func TestLogResource_EmptyPath(t *testing.T) {
	resource := &Resource{
		Resource: &Resource_Logs{
			Logs: &LogResource{
				Path: "",
			},
		},
	}

	sd := &ServiceDiscovery{
		GeneratedServiceName: &ServiceName{
			Name:   "test-service",
			Source: ServiceNameSource_SERVICE_NAME_SOURCE_UNKNOWN,
		},
		Resources: []*Resource{resource},
	}

	data, err := proto.Marshal(sd)
	require.NoError(t, err)

	decoded := &ServiceDiscovery{}
	err = proto.Unmarshal(data, decoded)
	require.NoError(t, err)

	require.Len(t, decoded.Resources, 1)
	logs := decoded.Resources[0].GetLogs()
	require.NotNil(t, logs)
	assert.Empty(t, logs.Path)
}

func TestServiceDiscovery_BackwardCompatibility(t *testing.T) {
	original := &ServiceDiscovery{
		GeneratedServiceName: &ServiceName{
			Name:   "legacy-service",
			Source: ServiceNameSource_SERVICE_NAME_SOURCE_UNKNOWN,
		},
		ApmInstrumentation: false,
	}

	data, err := proto.Marshal(original)
	require.NoError(t, err)

	decoded := &ServiceDiscovery{}
	err = proto.Unmarshal(data, decoded)
	require.NoError(t, err)

	assert.Equal(t, "legacy-service", decoded.GeneratedServiceName.Name)
	assert.False(t, decoded.ApmInstrumentation)
	assert.Nil(t, decoded.Resources, "Resources should be nil for backward compatibility")
}

func TestResource_OneofAccessors(t *testing.T) {
	logResource := &Resource{
		Resource: &Resource_Logs{
			Logs: &LogResource{
				Path: "/test.log",
			},
		},
	}

	logs := logResource.GetLogs()
	require.NotNil(t, logs, "GetLogs() should return the LogResource")
	assert.Equal(t, "/test.log", logs.Path)

	require.NotNil(t, logResource.Resource, "Resource oneof should be set")
	_, ok := logResource.Resource.(*Resource_Logs)
	assert.True(t, ok, "Resource should be of type *Resource_Logs")
}

func TestLogResource_ManyResources(t *testing.T) {
	resources := make([]*Resource, 100)
	for i := 0; i < 100; i++ {
		resources[i] = &Resource{
			Resource: &Resource_Logs{
				Logs: &LogResource{
					Path: "/var/log/file" + string(rune(i)) + ".log",
				},
			},
		}
	}

	sd := &ServiceDiscovery{
		GeneratedServiceName: &ServiceName{
			Name:   "multi-log-service",
			Source: ServiceNameSource_SERVICE_NAME_SOURCE_UNKNOWN,
		},
		Resources: resources,
	}

	data, err := proto.Marshal(sd)
	require.NoError(t, err)

	decoded := &ServiceDiscovery{}
	err = proto.Unmarshal(data, decoded)
	require.NoError(t, err)

	require.Len(t, decoded.Resources, 100, "Should preserve all 100 resources")
	for i, resource := range decoded.Resources {
		logs := resource.GetLogs()
		require.NotNil(t, logs)
		assert.Contains(t, logs.Path, "/var/log/file")
		_ = i
	}
}

func TestLogResource_SpecialCharactersInPaths(t *testing.T) {
	specialPaths := []string{
		"/var/log/app with spaces.log",
		"/var/log/special!@#$%^&*().log",
		"/var/log/unicode-日本語.log",
		"/var/log/quotes-'single'.log",
		"/var/log/quotes-\"double\".log",
	}

	resources := make([]*Resource, len(specialPaths))
	for i, path := range specialPaths {
		resources[i] = &Resource{
			Resource: &Resource_Logs{
				Logs: &LogResource{
					Path: path,
				},
			},
		}
	}

	sd := &ServiceDiscovery{
		GeneratedServiceName: &ServiceName{
			Name:   "special-service",
			Source: ServiceNameSource_SERVICE_NAME_SOURCE_UNKNOWN,
		},
		Resources: resources,
	}

	data, err := proto.Marshal(sd)
	require.NoError(t, err)

	decoded := &ServiceDiscovery{}
	err = proto.Unmarshal(data, decoded)
	require.NoError(t, err)

	require.Len(t, decoded.Resources, 5)
	actualPaths := make([]string, len(decoded.Resources))
	for i, resource := range decoded.Resources {
		logs := resource.GetLogs()
		require.NotNil(t, logs)
		actualPaths[i] = logs.Path
	}
	assert.ElementsMatch(t, specialPaths, actualPaths)
}

func TestServiceDiscovery_MultipleResourceTypes(t *testing.T) {
	sd := &ServiceDiscovery{
		GeneratedServiceName: &ServiceName{
			Name:   "future-ready-service",
			Source: ServiceNameSource_SERVICE_NAME_SOURCE_UNKNOWN,
		},
		Resources: []*Resource{
			{
				Resource: &Resource_Logs{
					Logs: &LogResource{
						Path: "/var/log/app.log",
					},
				},
			},
		},
	}

	data, err := proto.Marshal(sd)
	require.NoError(t, err)

	decoded := &ServiceDiscovery{}
	err = proto.Unmarshal(data, decoded)
	require.NoError(t, err)

	require.Len(t, decoded.Resources, 1)
	assert.NotNil(t, decoded.Resources[0].GetLogs())
}

func TestLogResource_PathPreservation(t *testing.T) {
	testPaths := []string{
		"/var/log/nginx/access.log",
		"/var/log/nginx/error.log",
		"/var/log/app.log",
	}

	resources := make([]*Resource, len(testPaths))
	for i, path := range testPaths {
		resources[i] = &Resource{
			Resource: &Resource_Logs{
				Logs: &LogResource{
					Path: path,
				},
			},
		}
	}

	sd := &ServiceDiscovery{
		GeneratedServiceName: &ServiceName{
			Name:   "test",
			Source: ServiceNameSource_SERVICE_NAME_SOURCE_UNKNOWN,
		},
		Resources: resources,
	}

	data, err := proto.Marshal(sd)
	require.NoError(t, err)

	decoded := &ServiceDiscovery{}
	err = proto.Unmarshal(data, decoded)
	require.NoError(t, err)

	require.Len(t, decoded.Resources, 3)
	actualPaths := make([]string, len(decoded.Resources))
	for i, resource := range decoded.Resources {
		logs := resource.GetLogs()
		require.NotNil(t, logs)
		actualPaths[i] = logs.Path
	}
	assert.ElementsMatch(t, testPaths, actualPaths)
}
