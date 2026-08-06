# Architecture

The cafe is three packages and a web front end.

`internal/menu` holds the board and the prices. `internal/orders` takes an order
at the counter and keeps the barista queue. `internal/billing` turns a total
into a receipt, with VAT.

`cats` is the resident roster, which is not really part of the ordering flow but
everyone wants it in the app anyway.

`web` is the public page. It reads the menu at build time for now.
