// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package vswitchd

import "github.com/ovn-org/libovsdb/model"

const CTZoneTable = "CT_Zone"

// CTZone defines an object in CT_Zone table
type CTZone struct {
	UUID          string            `ovsdb:"_uuid"`
	ExternalIDs   map[string]string `ovsdb:"external_ids"`
	TimeoutPolicy *string           `ovsdb:"timeout_policy"`
}

func (a *CTZone) GetUUID() string {
	return a.UUID
}

func (a *CTZone) GetExternalIDs() map[string]string {
	return a.ExternalIDs
}

func copyCTZoneExternalIDs(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalCTZoneExternalIDs(a, b map[string]string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if w, ok := b[k]; !ok || v != w {
			return false
		}
	}
	return true
}

func (a *CTZone) GetTimeoutPolicy() *string {
	return a.TimeoutPolicy
}

func copyCTZoneTimeoutPolicy(a *string) *string {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalCTZoneTimeoutPolicy(a, b *string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *CTZone) DeepCopyInto(b *CTZone) {
	*b = *a
	b.ExternalIDs = copyCTZoneExternalIDs(a.ExternalIDs)
	b.TimeoutPolicy = copyCTZoneTimeoutPolicy(a.TimeoutPolicy)
}

func (a *CTZone) DeepCopy() *CTZone {
	b := new(CTZone)
	a.DeepCopyInto(b)
	return b
}

func (a *CTZone) CloneModelInto(b model.Model) {
	c := b.(*CTZone)
	a.DeepCopyInto(c)
}

func (a *CTZone) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *CTZone) Equals(b *CTZone) bool {
	return a.UUID == b.UUID &&
		equalCTZoneExternalIDs(a.ExternalIDs, b.ExternalIDs) &&
		equalCTZoneTimeoutPolicy(a.TimeoutPolicy, b.TimeoutPolicy)
}

func (a *CTZone) EqualsModel(b model.Model) bool {
	c := b.(*CTZone)
	return a.Equals(c)
}

var _ model.CloneableModel = &CTZone{}
var _ model.ComparableModel = &CTZone{}
