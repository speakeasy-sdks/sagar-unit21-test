# ExportCasesRequestBody

To filter your response.


## Fields

| Field                                                                                      | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `CaseIds`                                                                                  | []*int64*                                                                                  | :heavy_minus_sign:                                                                         | Array of the unique identifiers of the case IDs.                                           |
| `Filters`                                                                                  | [*ExportCasesRequestBodyFilters](../../models/operations/exportcasesrequestbodyfilters.md) | :heavy_minus_sign:                                                                         | Filter to narrow down cases in export                                                      |