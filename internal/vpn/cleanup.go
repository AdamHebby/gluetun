package vpn

import (
	"context"
	"errors"
)

func (l *Loop) cleanup(vpnProvider string) {
	for _, vpnPort := range l.vpnInputPorts {
		err := l.fw.RemoveAllowedPort(context.Background(), vpnPort)
		if err != nil {
			l.logger.Error("cannot remove allowed input port from firewall: " + err.Error())
		}
	}

	err := l.publicip.ClearData()
	if err != nil {
		l.logger.Error("clearing public IP data: " + err.Error())
	}

	err = l.stopPortForwarding(vpnProvider)
	if err != nil {
		portForwardingAlreadyStopped := errors.Is(err, context.Canceled)
		if !portForwardingAlreadyStopped {
			l.logger.Error("stopping port forwarding: " + err.Error())
		}
	}
}
