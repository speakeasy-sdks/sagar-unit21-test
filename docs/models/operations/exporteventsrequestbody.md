# ExportEventsRequestBody


## Fields

| Field                                                                                        | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `EventIds`                                                                                   | []*int64*                                                                                    | :heavy_minus_sign:                                                                           | Array of the unique identifiers of the event IDs.                                            |
| `Filters`                                                                                    | [*ExportEventsRequestBodyFilters](../../models/operations/exporteventsrequestbodyfilters.md) | :heavy_minus_sign:                                                                           | Filter to narrow down events in export                                                       |