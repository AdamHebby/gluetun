// Package purevpn contains code to obtain the server information
// for the PureVPN provider.
package purevpn

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/qdm12/gluetun/internal/models"
	"github.com/qdm12/gluetun/internal/publicip"
	"github.com/qdm12/gluetun/internal/updater/openvpn"
	"github.com/qdm12/gluetun/internal/updater/resolver"
	"github.com/qdm12/gluetun/internal/updater/unzip"
)

var ErrNotEnoughServers = errors.New("not enough servers found")

func GetServers(ctx context.Context, client *http.Client,
	unzipper unzip.Unzipper, presolver resolver.Parallel, minServers int) (
	servers []models.PurevpnServer, warnings []string, err error) {
	const url = "https://s3-us-west-1.amazonaws.com/heartbleed/windows/New+OVPN+Files.zip"
	contents, err := unzipper.FetchAndExtract(ctx, url)
	if err != nil {
		return nil, nil, err
	} else if len(contents) < minServers {
		return nil, nil, fmt.Errorf("%w: %d and expected at least %d",
			ErrNotEnoughServers, len(contents), minServers)
	}

	hts := make(hostToServer)

	for fileName, content := range contents {
		if !strings.HasSuffix(fileName, ".ovpn") {
			continue
		}

		host, warning, err := openvpn.ExtractHost(content)
		if warning != "" {
			warnings = append(warnings, warning)
		}

		if err != nil {
			// treat error as warning and go to next file
			warning := err.Error() + " in " + fileName
			warnings = append(warnings, warning)
			continue
		}

		hts.add(host)
	}

	if len(hts) < minServers {
		return nil, warnings, fmt.Errorf("%w: %d and expected at least %d",
			ErrNotEnoughServers, len(hts), minServers)
	}

	hosts := hts.toHostsSlice()
	hostToIPs, newWarnings, err := resolveHosts(ctx, presolver, hosts, minServers)
	warnings = append(warnings, newWarnings...)
	if err != nil {
		return nil, warnings, err
	}

	hts.adaptWithIPs(hostToIPs)

	servers = hts.toServersSlice()

	if len(servers) < minServers {
		return nil, warnings, fmt.Errorf("%w: %d and expected at least %d",
			ErrNotEnoughServers, len(servers), minServers)
	}

	// Dedup by location
	lts := make(locationToServer)
	for _, server := range servers {
		ipInfo, err := publicip.Info(ctx, client, server.IPs[0])
		if err != nil {
			return nil, warnings, err
		}

		// TODO split servers by host
		lts.add(ipInfo.Country, ipInfo.Region, ipInfo.City, server.IPs)
	}

	if len(servers) < minServers {
		return nil, warnings, fmt.Errorf("%w: %d and expected at least %d",
			ErrNotEnoughServers, len(servers), minServers)
	}

	servers = lts.toServersSlice()

	sortServers(servers)

	return servers, warnings, nil
}