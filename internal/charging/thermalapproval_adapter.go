package charging

func ApplyThermalApproval(state ThermalApprovalState) ThermalApprovalState {
	if state.Mitigated && state.ProtocolOK && state.CableOK && state.PowerOK {
		state.Status = "certified"
	} else {
		state.Status = "retest_required"
	}
	return state
}
func IsThermalApprovalReady(state ThermalApprovalState) bool { return state.Status == "certified" }
