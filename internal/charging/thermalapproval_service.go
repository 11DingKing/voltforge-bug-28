package charging

import "errors"

var ErrThermalApprovalState = errors.New("invalid thermalapproval state")

type ThermalApprovalState struct {
	Mitigated, ProtocolOK, CableOK, ThermalOK, PowerOK bool
	Status                                             string
}

func (s ThermalApprovalState) CanCertify() bool {
	return s.Mitigated && s.ProtocolOK && s.CableOK && s.PowerOK
}
