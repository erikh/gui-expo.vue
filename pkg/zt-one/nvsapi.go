package one

type NetworkStatus struct {
	Address           string `json:"address"`
	PublicIdentity    string `json:"publicIdentity"`
	WorldID           int64  `json:"worldId"`
	WorldTimestamp    int64  `json:"worldTimestamp"`
	Online            bool   `json:"online"`
	TCPFallbackActive bool   `json:"tcpFallbackActive"`
	// Enum: ALWAYS / TRUSTED / NEVER
	RelayPolicy  string `json:"relayPolicy"`
	VersionMajor int64  `json:"versionMajor"`
	VersionMinor int64  `json:"versionMinor"`
	VersionRev   int64  `json:"versionRev"`
	Version      string `json:"version"`
	Clock        int64  `json:"clock"`
}

type Network struct {
	ID                string   `json:"id"`
	LegacyNetworkID   string   `json:"nwid"`
	MAC               string   `json:"mac"`
	Name              string   `json:"name"`
	Status            string   `json:"status"`
	Type              string   `json:"type"`
	MTU               int      `json:"mtu"`
	DHCP              bool     `json:"dhcp"`
	Bridge            bool     `json:"bridge"`
	BroadcastEnabled  bool     `json:"broadcastEnabled"`
	PortError         int      `json:"portError"`
	NetconfRevision   int64    `json:"netconfRevision"`
	AssignedAddresses []string `json:"assignedAddresses"`
	Routes            []Route  `json:"routes"`
	PortDeviceName    string   `json:"portDeviceName"`
	AllowManaged      bool     `json:"allowManaged"`
	AllowGlobal       bool     `json:"allowGlobal"`
	AllowDefault      bool     `json:"allowDefault"`
	AllowDNS          bool     `json:"allowDNS"`
}

type Route struct {
	Target string `json:"target"`
	Via    string `json:"via"`
	Flags  int64  `json:"flags"`
	Metric int64  `json:"metric"`
}

func (c *Client) Status() (*NetworkStatus, error) {
	ns := &NetworkStatus{}
	return ns, c.wrapJSON("/status", ns)
}

func (c *Client) Networks() ([]*Network, error) {
	nws := []*Network{}
	return nws, c.wrapJSON("/network", &nws)
}
