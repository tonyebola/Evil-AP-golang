package files

import (
	"fmt"
	"text/template"
	"bytes"
	"os"
)

type HostAPDVariables struct {
	// Evil AP Interface
	Interface string
	// Evil AP Name
	SSID string
	// Operating Channel #
	Channel string
	// country code
	CountryCode string
}

const hostapdConf = `
interface={{.Interface}}
driver=nl80211
# ssid={{.SSID}}
ssid=NETGEAR12345
hw_mode=g
channel={{.Channel}}
macaddr_acl=0
auth_algs=1
ignore_broadcast_ssid=0
country_code={{.CountryCode}}
# wpa=1
wpa=2
wpa_passphrase=passwordhere
wpa_key_mgmt=WPA-PSK
wpa_pairwise=TKIP
rsn_pairwise=CCMP
bssid=A5:2B:8C:C1:27:84

# Country code (ISO/IEC 3166-1). Used to set regulatory domain.
# Set as needed to indicate country in which device is operating.
# This can limit available channels and transmit power.
#country_code=US
#country_code=BO

##### WPA/IEEE 802.11i configuration ##########################################

# Enable WPA. Setting this variable configures the AP to require WPA (either
# WPA-PSK or WPA-RADIUS/EAP based on other configuration). For WPA-PSK, either
# wpa_psk or wpa_passphrase must be set and wpa_key_mgmt must include WPA-PSK.
# For WPA-RADIUS/EAP, ieee8021x must be set (but without dynamic WEP keys),
# RADIUS authentication server must be configured, and WPA-EAP must be included
# in wpa_key_mgmt.
# This field is a bit field that can be used to enable WPA (IEEE 802.11i/D3.0)
# and/or WPA2 (full IEEE 802.11i/RSN):
# bit0 = WPA
# bit1 = IEEE 802.11i/RSN (WPA2) (dot11RSNAEnabled)
#wpa=1

# WPA pre-shared keys for WPA-PSK. This can be either entered as a 256-bit
# secret in hex format (64 hex digits), wpa_psk, or as an ASCII passphrase
# (8..63 characters) that will be converted to PSK. This conversion uses SSID
# so the PSK changes when ASCII passphrase is used and the SSID is changed.
# wpa_psk (dot11RSNAConfigPSKValue)
# wpa_passphrase (dot11RSNAConfigPSKPassPhrase)
#wpa_psk=0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef
#wpa_passphrase=secret passphrase

#bssid=00:13:10:95:fe:0b
`

func WriteHostAPDConfFile(filePath string, vars *HostAPDVariables) {
	var (
		err error
	)

	t := template.New("hostapd config template")

	t, err = t.Parse(hostapdConf)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing hostapd config:", err)
		panic(err)
	}

	var tpl bytes.Buffer
	if tErr := t.Execute(&tpl, *vars); tErr != nil {
		fmt.Fprintln(os.Stderr, "Error executing hostapd config template:", tErr)
    panic(tErr)
	}

	result := tpl.String()

	WriteStringToFile(filePath, result)
}
