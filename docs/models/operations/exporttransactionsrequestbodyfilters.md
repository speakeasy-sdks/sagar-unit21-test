# ExportTransactionsRequestBodyFilters

Filter to narrow down events in export


## Fields

| Field                                       | Type                                        | Required                                    | Description                                 | Example                                     |
| ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- | ------------------------------------------- |
| `Currency`                                  | **string*                                   | :heavy_minus_sign:                          | Currency used in the transaction.           | USD                                         |
| `EndDate`                                   | **string*                                   | :heavy_minus_sign:                          | Event creation date end.                    | 2021-11-05 04:13:46                         |
| `EntityIds`                                 | []*int64*                                   | :heavy_minus_sign:                          | Numerical IDs of the entities in the event. |                                             |
| `MaximumAmount`                             | **int64*                                    | :heavy_minus_sign:                          | Maximum amount in the event transaction.    | 100000                                      |
| `MinimumAmount`                             | **int64*                                    | :heavy_minus_sign:                          | Minimum amount in the event transaction.    | 1000                                        |
| `StartDate`                                 | **string*                                   | :heavy_minus_sign:                          | Event creation date start.                  | 2019-11-05 04:13:46                         |
| `Status`                                    | **string*                                   | :heavy_minus_sign:                          | Status of the event.                        | active                                      |
| `Statuses`                                  | []*string*                                  | :heavy_minus_sign:                          | Status for the events.                      |                                             |
| `TagIds`                                    | []*int64*                                   | :heavy_minus_sign:                          | Numerical IDs of the tags.                  |                                             |