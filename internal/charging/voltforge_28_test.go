package charging

import "testing"

func TestVoltForge28(t *testing.T) {
	state := ApplyThermalApproval(ThermalApprovalState{Mitigated: true, ProtocolOK: true, CableOK: true, ThermalOK: false, PowerOK: true})
	if IsThermalApprovalReady(state) || state.Status != "retest_required" {
		t.Fatalf("unsafe state was certified: %+v", state)
	}
}
