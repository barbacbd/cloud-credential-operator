package security

import (
    i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f "github.com/microsoft/kiota-abstractions-go"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// EdiscoveryReviewTag 
type EdiscoveryReviewTag struct {
    Tag
    // Indicates whether a single or multiple child tags can be associated with a document. Possible values are: One, Many.  This value controls whether the UX presents the tags as checkboxes or a radio button group.
    childSelectability *ChildSelectability
    // Returns the tags that are a child of a tag.
    childTags []EdiscoveryReviewTagable
    // Returns the parent tag of the specified tag.
    parent EdiscoveryReviewTagable
}
// NewEdiscoveryReviewTag instantiates a new EdiscoveryReviewTag and sets the default values.
func NewEdiscoveryReviewTag()(*EdiscoveryReviewTag) {
    m := &EdiscoveryReviewTag{
        Tag: *NewTag(),
    }
    odataTypeValue := "#microsoft.graph.security.ediscoveryReviewTag";
    m.SetOdataType(&odataTypeValue);
    return m
}
// CreateEdiscoveryReviewTagFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEdiscoveryReviewTagFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewEdiscoveryReviewTag(), nil
}
// GetChildSelectability gets the childSelectability property value. Indicates whether a single or multiple child tags can be associated with a document. Possible values are: One, Many.  This value controls whether the UX presents the tags as checkboxes or a radio button group.
func (m *EdiscoveryReviewTag) GetChildSelectability()(*ChildSelectability) {
    return m.childSelectability
}
// GetChildTags gets the childTags property value. Returns the tags that are a child of a tag.
func (m *EdiscoveryReviewTag) GetChildTags()([]EdiscoveryReviewTagable) {
    return m.childTags
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EdiscoveryReviewTag) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.Tag.GetFieldDeserializers()
    res["childSelectability"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetEnumValue(ParseChildSelectability , m.SetChildSelectability)
    res["childTags"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetCollectionOfObjectValues(CreateEdiscoveryReviewTagFromDiscriminatorValue , m.SetChildTags)
    res["parent"] = i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.SetObjectValue(CreateEdiscoveryReviewTagFromDiscriminatorValue , m.SetParent)
    return res
}
// GetParent gets the parent property value. Returns the parent tag of the specified tag.
func (m *EdiscoveryReviewTag) GetParent()(EdiscoveryReviewTagable) {
    return m.parent
}
// Serialize serializes information the current object
func (m *EdiscoveryReviewTag) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.Tag.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetChildSelectability() != nil {
        cast := (*m.GetChildSelectability()).String()
        err = writer.WriteStringValue("childSelectability", &cast)
        if err != nil {
            return err
        }
    }
    if m.GetChildTags() != nil {
        cast := i2ae4187f7daee263371cb1c977df639813ab50ffa529013b7437480d1ec0158f.CollectionCast[i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable](m.GetChildTags())
        err = writer.WriteCollectionOfObjectValues("childTags", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("parent", m.GetParent())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetChildSelectability sets the childSelectability property value. Indicates whether a single or multiple child tags can be associated with a document. Possible values are: One, Many.  This value controls whether the UX presents the tags as checkboxes or a radio button group.
func (m *EdiscoveryReviewTag) SetChildSelectability(value *ChildSelectability)() {
    m.childSelectability = value
}
// SetChildTags sets the childTags property value. Returns the tags that are a child of a tag.
func (m *EdiscoveryReviewTag) SetChildTags(value []EdiscoveryReviewTagable)() {
    m.childTags = value
}
// SetParent sets the parent property value. Returns the parent tag of the specified tag.
func (m *EdiscoveryReviewTag) SetParent(value EdiscoveryReviewTagable)() {
    m.parent = value
}
