package charging

func ApplyThermalApproval(state ThermalApprovalState) ThermalApprovalState {
	if state.CanCertify() {
		state.Status = "certified"
		return state
	}
	state.Status = "retest_required"
	return state
}
func IsThermalApprovalReady(state ThermalApprovalState) bool { return state.Status == "certified" }
