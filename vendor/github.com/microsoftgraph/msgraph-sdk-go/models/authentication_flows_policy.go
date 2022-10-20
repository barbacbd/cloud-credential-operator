package models

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AuthenticationFlowsPolicy 
type AuthenticationFlowsPolicy struct {
    Entity
    // Inherited property. A description of the policy. Optional. Read-only.
    description *string
    // Inherited property. The human-readable name of the policy. Optional. Read-only.
    displayName *string
    // Contains selfServiceSignUpAuthenticationFlowConfiguration settings that convey whether self-service sign-up is enabled or disabled. Optional. Read-only.
    selfServiceSignUp SelfServiceSignUpAuthenticationFlowConfigurationable
}
// NewAuthenticationFlowsPolicy instantiates a new authenticationFlowsPolicy and sets the default values.
func NewAuthenticationFlowsPolicy()(*AuthenticationFlowsPolicy) {
    m := &AuthenticationFlowsPolicy{
        Entity: *NewEntity(),
    }
    odataTypeValue := "#microsoft.graph.authenticationFlowsPolicy";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateAuthenticationFlowsPolicyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAuthenticationFlowsPolicyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAuthenticationFlowsPolicy(), nil
}
// GetDescription gets the description property value. Inherited property. A description of the policy. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) GetDescription()(*string) {
    return m.description
}
// GetDisplayName gets the displayName property value. Inherited property. The human-readable name of the policy. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) GetDisplayName()(*string) {
    return m.displayName
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AuthenticationFlowsPolicy) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["description"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDescription)
    res["displayName"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetStringValue(m.SetDisplayName)
    res["selfServiceSignUp"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateSelfServiceSignUpAuthenticationFlowConfigurationFromDiscriminatorValue , m.SetSelfServiceSignUp)
    return res
}
// GetSelfServiceSignUp gets the selfServiceSignUp property value. Contains selfServiceSignUpAuthenticationFlowConfiguration settings that convey whether self-service sign-up is enabled or disabled. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) GetSelfServiceSignUp()(SelfServiceSignUpAuthenticationFlowConfigurationable) {
    return m.selfServiceSignUp
}
// Serialize serializes information the current object
func (m *AuthenticationFlowsPolicy) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("description", m.GetDescription())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("selfServiceSignUp", m.GetSelfServiceSignUp())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetDescription sets the description property value. Inherited property. A description of the policy. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) SetDescription(value *string)() {
    m.description = value
}
// SetDisplayName sets the displayName property value. Inherited property. The human-readable name of the policy. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) SetDisplayName(value *string)() {
    m.displayName = value
}
// SetSelfServiceSignUp sets the selfServiceSignUp property value. Contains selfServiceSignUpAuthenticationFlowConfiguration settings that convey whether self-service sign-up is enabled or disabled. Optional. Read-only.
func (m *AuthenticationFlowsPolicy) SetSelfServiceSignUp(value SelfServiceSignUpAuthenticationFlowConfigurationable)() {
    m.selfServiceSignUp = value
}
