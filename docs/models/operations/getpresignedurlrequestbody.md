# GetPreSignedURLRequestBody


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `FileName`                                                         | *string*                                                           | :heavy_check_mark:                                                 | File name.                                                         | custom_data.csv                                                    |
| `Md5Hash`                                                          | **string*                                                          | :heavy_minus_sign:                                                 | Unique hash. Requires UUID format. Used if filename is not unique. | 352bfecf-ce8e-4c3d-64b9-ba0707fc2496                               |
| `StreamName`                                                       | *string*                                                           | :heavy_check_mark:                                                 | Stream name.                                                       | insts                                                              |