package settings

import "errors"

var (
	ErrCityNotValid                    = errors.New("the city specified is not valid")
	ErrControlServerAddress            = errors.New("listening address it not valid")
	ErrControlServerPort               = errors.New("listening port it not valid")
	ErrControlServerPrivilegedPort     = errors.New("cannot use privileged port without running as root")
	ErrCountryNotValid                 = errors.New("the country specified is not valid")
	ErrFirewallZeroPort                = errors.New("cannot have a zero port to block")
	ErrHostnameNotValid                = errors.New("the hostname specified is not valid")
	ErrISPNotValid                     = errors.New("the ISP specified is not valid")
	ErrNameNotValid                    = errors.New("the server name specified is not valid")
	ErrOpenVPNClientCertMissing        = errors.New("client certificate is missing")
	ErrOpenVPNClientCertNotValid       = errors.New("client certificate is not valid")
	ErrOpenVPNClientKeyMissing         = errors.New("client key is missing")
	ErrOpenVPNClientKeyNotValid        = errors.New("client key is not valid")
	ErrOpenVPNConfigFile               = errors.New("custom configuration file error")
	ErrOpenVPNCustomPortNotAllowed     = errors.New("custom endpoint port is not allowed")
	ErrOpenVPNEncryptionPresetNotValid = errors.New("PIA encryption preset is not valid")
	ErrOpenVPNInterfaceNotValid        = errors.New("interface name is not valid")
	ErrOpenVPNMSSFixIsTooHigh          = errors.New("mssfix option value is too high")
	ErrOpenVPNPasswordIsEmpty          = errors.New("password is empty")
	ErrOpenVPNTCPNotSupported          = errors.New("TCP protocol is not supported")
	ErrOpenVPNUserIsEmpty              = errors.New("user is empty")
	ErrOpenVPNVerbosityIsOutOfBounds   = errors.New("verbosity value is out of bounds")
	ErrOpenVPNVersionIsNotValid        = errors.New("version is not valid")
	ErrPortForwardingEnabled           = errors.New("port forwarding cannot be enabled")
	ErrPortForwardingFilepathNotValid  = errors.New("port forwarding filepath given is not valid")
	ErrPublicIPFilepathNotValid        = errors.New("public IP address file path is not valid")
	ErrPublicIPPeriodTooShort          = errors.New("public IP address check period is too short")
	ErrRegionNotValid                  = errors.New("the region specified is not valid")
	ErrServerAddressNotValid           = errors.New("server listening address is not valid")
	ErrSystemPGIDNotValid              = errors.New("process group id is not valid")
	ErrSystemPUIDNotValid              = errors.New("process user id is not valid")
	ErrSystemTimezoneNotValid          = errors.New("timezone is not valid")
	ErrUpdaterPeriodTooSmall           = errors.New("VPN server data updater period is too small")
	ErrVPNProviderNameNotValid         = errors.New("VPN provider name is not valid")
	ErrVPNTypeNotValid                 = errors.New("VPN type is not valid")
	ErrWireguardEndpointIPNotSet       = errors.New("endpoint IP is not set")
	ErrWireguardEndpointPortNotAllowed = errors.New("endpoint port is not allowed")
	ErrWireguardEndpointPortNotSet     = errors.New("endpoint port is not set")
	ErrWireguardInterfaceAddressNotSet = errors.New("interface address is not set")
	ErrWireguardInterfaceNotValid      = errors.New("interface name is not valid")
	ErrWireguardPreSharedKeyNotSet     = errors.New("pre-shared key is not set")
	ErrWireguardPreSharedKeyNotValid   = errors.New("pre-shared key is not valid")
	ErrWireguardPrivateKeyNotSet       = errors.New("private key is not set")
	ErrWireguardPrivateKeyNotValid     = errors.New("private key is not valid")
	ErrWireguardPublicKeyNotSet        = errors.New("public key is not set")
	ErrWireguardPublicKeyNotValid      = errors.New("public key is not valid")
)