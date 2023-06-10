# ExportAlertsRequestBody

To filter your response.


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `AlertIds`                                                                                   | []*int64*                                                                                    | :heavy_minus_sign:                                                                           | Array of the unique identifiers of the alert IDs.                                            |
| `Filters`                                                                                    | [*ExportAlertsRequestBodyFilters](../../models/operations/exportalertsrequestbodyfilters.md) | :heavy_minus_sign:                                                                           | Filter to narrow down alerts in export                                                       |