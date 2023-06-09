// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type BlacklistCIDR struct {
	// CIDR IP range. Check ranges with [this calculator](https://www.ipaddressguide.com/cidr)
	Cidr string `json:"cidr"`
	// Origin of IP range
	Source *string `json:"source,omitempty"`
}
