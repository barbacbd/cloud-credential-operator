package models
import (
    "errors"
)
// Provides operations to call the sync method.
type DeviceManagementExchangeConnectorSyncType int

const (
    // Discover all the device in Exchange.
    FULLSYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE DeviceManagementExchangeConnectorSyncType = iota
    // Discover only the device in Exchange which have updated during the delta sync window.
    DELTASYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE
)

func (i DeviceManagementExchangeConnectorSyncType) String() string {
    return []string{"fullSync", "deltaSync"}[i]
}
func ParseDeviceManagementExchangeConnectorSyncType(v string) (interface{}, error) {
    result := FULLSYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE
    switch v {
        case "fullSync":
            result = FULLSYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE
        case "deltaSync":
            result = DELTASYNC_DEVICEMANAGEMENTEXCHANGECONNECTORSYNCTYPE
        default:
            return 0, errors.New("Unknown DeviceManagementExchangeConnectorSyncType value: " + v)
    }
    return &result, nil
}
func SerializeDeviceManagementExchangeConnectorSyncType(values []DeviceManagementExchangeConnectorSyncType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
