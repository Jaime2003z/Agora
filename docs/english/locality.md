# Locality Identifier Specification

## Format

LocalityID := <SCOPE>

SCOPE :=
  GLOBAL |
  <ISO-3166-1> |
  <ISO-3166-1>-<ISO-3166-2> |
  <ISO-3166-1>-<ISO-3166-2>-<LOCAL-CODE>


## Definitions

### ISO-3166-1
- 2-letter country codes as defined in ISO 3166-1 alpha-2
- Examples: "US" (United States), "JP" (Japan), "CO" (Colombia)

### ISO-3166-2
- Subdivision codes as defined in ISO 3166-2
- Examples: "CA" (California), "13" (Tokyo), "ANT" (Antioquia)

### ZIP-CODE
- ZIP codes as defined in ISO 3166-2
- Examples: "12345", "12345-6789"

## Examples

1. Global scope: `GLOBAL`
2. Country scope: `CO`
3. State scope: `CO-ANT`
4. Municipality scope: `CO-ANT-05001` - (Colombia - Antioquia - Medellín)