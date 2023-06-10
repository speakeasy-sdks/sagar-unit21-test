# DigitalDataArrayProperties

Associated digital properties - IP, device, browser, client info etc.


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ClientFingerprints`                                         | []*string*                                                   | :heavy_minus_sign:                                           | List of unique addresses.                                    |
| `IPAddresses`                                                | []*string*                                                   | :heavy_minus_sign:                                           | List of IP addresses. MUST be in either IPv4 or IPv6 format. |