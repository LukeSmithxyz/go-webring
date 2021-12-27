<!--
SPDX-FileCopyrightText: 2021 Amolith <amolith@secluded.site>

SPDX-License-Identifier: CC0-1.0
-->

# go-webring
Simple webring implementation

## TODO
- [ ] Implement some kind of rate-limiting or caching; the list of ring members
      is re-read on each request and one could abuse this to DOS the server.
      - It does save admins from having to restart the server when adding new
      members, though, so I would like to keep the current functionality.
      However, some sort of DOS mitigation would still be good Just In Case™.
