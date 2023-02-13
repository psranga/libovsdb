// Code generated by "libovsdb.modelgen"
// DO NOT EDIT.

package vswitchd

import "github.com/ovn-org/libovsdb/model"

const PortTable = "Port"

type (
	PortBondMode = string
	PortLACP     = string
	PortVLANMode = string
)

var (
	PortBondModeActiveBackup   PortBondMode = "active-backup"
	PortBondModeBalanceSLB     PortBondMode = "balance-slb"
	PortBondModeBalanceTCP     PortBondMode = "balance-tcp"
	PortLACPActive             PortLACP     = "active"
	PortLACPOff                PortLACP     = "off"
	PortLACPPassive            PortLACP     = "passive"
	PortVLANModeAccess         PortVLANMode = "access"
	PortVLANModeDot1qTunnel    PortVLANMode = "dot1q-tunnel"
	PortVLANModeNativeTagged   PortVLANMode = "native-tagged"
	PortVLANModeNativeUntagged PortVLANMode = "native-untagged"
	PortVLANModeTrunk          PortVLANMode = "trunk"
)

// Port defines an object in Port table
type Port struct {
	UUID            string            `ovsdb:"_uuid"`
	BondActiveSlave *string           `ovsdb:"bond_active_slave"`
	BondDowndelay   int               `ovsdb:"bond_downdelay"`
	BondFakeIface   bool              `ovsdb:"bond_fake_iface"`
	BondMode        *PortBondMode     `ovsdb:"bond_mode"`
	BondUpdelay     int               `ovsdb:"bond_updelay"`
	CVLANs          [4096]int         `ovsdb:"cvlans"`
	ExternalIDs     map[string]string `ovsdb:"external_ids"`
	FakeBridge      bool              `ovsdb:"fake_bridge"`
	Interfaces      []string          `ovsdb:"interfaces"`
	LACP            *PortLACP         `ovsdb:"lacp"`
	MAC             *string           `ovsdb:"mac"`
	Name            string            `ovsdb:"name"`
	OtherConfig     map[string]string `ovsdb:"other_config"`
	Protected       bool              `ovsdb:"protected"`
	QOS             *string           `ovsdb:"qos"`
	RSTPStatistics  map[string]int    `ovsdb:"rstp_statistics"`
	RSTPStatus      map[string]string `ovsdb:"rstp_status"`
	Statistics      map[string]int    `ovsdb:"statistics"`
	Status          map[string]string `ovsdb:"status"`
	Tag             *int              `ovsdb:"tag"`
	Trunks          [4096]int         `ovsdb:"trunks"`
	VLANMode        *PortVLANMode     `ovsdb:"vlan_mode"`
}

func (a *Port) GetUUID() string {
	return a.UUID
}

func (a *Port) GetBondActiveSlave() *string {
	return a.BondActiveSlave
}

func copyPortBondActiveSlave(a *string) *string {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalPortBondActiveSlave(a, b *string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *Port) GetBondDowndelay() int {
	return a.BondDowndelay
}

func (a *Port) GetBondFakeIface() bool {
	return a.BondFakeIface
}

func (a *Port) GetBondMode() *PortBondMode {
	return a.BondMode
}

func copyPortBondMode(a *PortBondMode) *PortBondMode {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalPortBondMode(a, b *PortBondMode) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *Port) GetBondUpdelay() int {
	return a.BondUpdelay
}

func (a *Port) GetCVLANs() [4096]int {
	return a.CVLANs
}

func (a *Port) GetExternalIDs() map[string]string {
	return a.ExternalIDs
}

func copyPortExternalIDs(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalPortExternalIDs(a, b map[string]string) bool {
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

func (a *Port) GetFakeBridge() bool {
	return a.FakeBridge
}

func (a *Port) GetInterfaces() []string {
	return a.Interfaces
}

func copyPortInterfaces(a []string) []string {
	if a == nil {
		return nil
	}
	b := make([]string, len(a))
	copy(b, a)
	return b
}

func equalPortInterfaces(a, b []string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}

func (a *Port) GetLACP() *PortLACP {
	return a.LACP
}

func copyPortLACP(a *PortLACP) *PortLACP {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalPortLACP(a, b *PortLACP) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *Port) GetMAC() *string {
	return a.MAC
}

func copyPortMAC(a *string) *string {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalPortMAC(a, b *string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *Port) GetName() string {
	return a.Name
}

func (a *Port) GetOtherConfig() map[string]string {
	return a.OtherConfig
}

func copyPortOtherConfig(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalPortOtherConfig(a, b map[string]string) bool {
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

func (a *Port) GetProtected() bool {
	return a.Protected
}

func (a *Port) GetQOS() *string {
	return a.QOS
}

func copyPortQOS(a *string) *string {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalPortQOS(a, b *string) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *Port) GetRSTPStatistics() map[string]int {
	return a.RSTPStatistics
}

func copyPortRSTPStatistics(a map[string]int) map[string]int {
	if a == nil {
		return nil
	}
	b := make(map[string]int, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalPortRSTPStatistics(a, b map[string]int) bool {
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

func (a *Port) GetRSTPStatus() map[string]string {
	return a.RSTPStatus
}

func copyPortRSTPStatus(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalPortRSTPStatus(a, b map[string]string) bool {
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

func (a *Port) GetStatistics() map[string]int {
	return a.Statistics
}

func copyPortStatistics(a map[string]int) map[string]int {
	if a == nil {
		return nil
	}
	b := make(map[string]int, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalPortStatistics(a, b map[string]int) bool {
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

func (a *Port) GetStatus() map[string]string {
	return a.Status
}

func copyPortStatus(a map[string]string) map[string]string {
	if a == nil {
		return nil
	}
	b := make(map[string]string, len(a))
	for k, v := range a {
		b[k] = v
	}
	return b
}

func equalPortStatus(a, b map[string]string) bool {
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

func (a *Port) GetTag() *int {
	return a.Tag
}

func copyPortTag(a *int) *int {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalPortTag(a, b *int) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *Port) GetTrunks() [4096]int {
	return a.Trunks
}

func (a *Port) GetVLANMode() *PortVLANMode {
	return a.VLANMode
}

func copyPortVLANMode(a *PortVLANMode) *PortVLANMode {
	if a == nil {
		return nil
	}
	b := *a
	return &b
}

func equalPortVLANMode(a, b *PortVLANMode) bool {
	if (a == nil) != (b == nil) {
		return false
	}
	if a == b {
		return true
	}
	return *a == *b
}

func (a *Port) DeepCopyInto(b *Port) {
	*b = *a
	b.BondActiveSlave = copyPortBondActiveSlave(a.BondActiveSlave)
	b.BondMode = copyPortBondMode(a.BondMode)
	b.ExternalIDs = copyPortExternalIDs(a.ExternalIDs)
	b.Interfaces = copyPortInterfaces(a.Interfaces)
	b.LACP = copyPortLACP(a.LACP)
	b.MAC = copyPortMAC(a.MAC)
	b.OtherConfig = copyPortOtherConfig(a.OtherConfig)
	b.QOS = copyPortQOS(a.QOS)
	b.RSTPStatistics = copyPortRSTPStatistics(a.RSTPStatistics)
	b.RSTPStatus = copyPortRSTPStatus(a.RSTPStatus)
	b.Statistics = copyPortStatistics(a.Statistics)
	b.Status = copyPortStatus(a.Status)
	b.Tag = copyPortTag(a.Tag)
	b.VLANMode = copyPortVLANMode(a.VLANMode)
}

func (a *Port) DeepCopy() *Port {
	b := new(Port)
	a.DeepCopyInto(b)
	return b
}

func (a *Port) CloneModelInto(b model.Model) {
	c := b.(*Port)
	a.DeepCopyInto(c)
}

func (a *Port) CloneModel() model.Model {
	return a.DeepCopy()
}

func (a *Port) Equals(b *Port) bool {
	return a.UUID == b.UUID &&
		equalPortBondActiveSlave(a.BondActiveSlave, b.BondActiveSlave) &&
		a.BondDowndelay == b.BondDowndelay &&
		a.BondFakeIface == b.BondFakeIface &&
		equalPortBondMode(a.BondMode, b.BondMode) &&
		a.BondUpdelay == b.BondUpdelay &&
		a.CVLANs == b.CVLANs &&
		equalPortExternalIDs(a.ExternalIDs, b.ExternalIDs) &&
		a.FakeBridge == b.FakeBridge &&
		equalPortInterfaces(a.Interfaces, b.Interfaces) &&
		equalPortLACP(a.LACP, b.LACP) &&
		equalPortMAC(a.MAC, b.MAC) &&
		a.Name == b.Name &&
		equalPortOtherConfig(a.OtherConfig, b.OtherConfig) &&
		a.Protected == b.Protected &&
		equalPortQOS(a.QOS, b.QOS) &&
		equalPortRSTPStatistics(a.RSTPStatistics, b.RSTPStatistics) &&
		equalPortRSTPStatus(a.RSTPStatus, b.RSTPStatus) &&
		equalPortStatistics(a.Statistics, b.Statistics) &&
		equalPortStatus(a.Status, b.Status) &&
		equalPortTag(a.Tag, b.Tag) &&
		a.Trunks == b.Trunks &&
		equalPortVLANMode(a.VLANMode, b.VLANMode)
}

func (a *Port) EqualsModel(b model.Model) bool {
	c := b.(*Port)
	return a.Equals(c)
}

var _ model.CloneableModel = &Port{}
var _ model.ComparableModel = &Port{}