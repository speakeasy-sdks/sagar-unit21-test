# RelationshipsDataProperties

An entity's associated intruments and entities (also called relationships)


## Fields

| Field                                                           | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `Children`                                                      | [][RelatedEntity](../../models/shared/relatedentity.md)         | :heavy_minus_sign:                                              | List of children entities.                                      |
| `Instruments`                                                   | [][RelatedInstrument](../../models/shared/relatedinstrument.md) | :heavy_minus_sign:                                              | List of associated instruments.                                 |
| `Parents`                                                       | [][RelatedEntity](../../models/shared/relatedentity.md)         | :heavy_minus_sign:                                              | List of parent entities.                                        |
| `Undirected`                                                    | [][RelatedEntity](../../models/shared/relatedentity.md)         | :heavy_minus_sign:                                              | List of undirected entities.                                    |