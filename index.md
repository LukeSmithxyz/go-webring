<!--
SPDX-FileCopyrightText: 2021 Amolith <amolith@secluded.site>

SPDX-License-Identifier: CC0-1.0
-->

---
title: Example webpage
subtitle: Basic input for homepage
---

Some basic boilerplate up here introducing the ring

## Membership Rules

- Rule 1
- Rule 2
- Rule 3

A caveat about not being able to disallow all content that would be
inappropriate

## Joining

* Step 1
* Step 2
* Step 3

Table of required links

| Destination     | Link                                                                                         |
|-----------------|----------------------------------------------------------------------------------------------|
| Previous member | [https://example.com/previous?host=YOURDOMAIN](https://example.com/previous?host=YOURDOMAIN) |
| Next member     | [https://example.com/next?host=YOURDOMAIN](https://example.com/next?host=YOURDOMAIN)         |
| Home page       | [https://example.com](https://example.com)                                                   |

Some example HTML for linking to the next and previous ring members.

```html
<p>
    <a href="https://example.com/previous?host=example.com">←</a>
    <a href="https://example.com/">Example Webring</a>
    <a href="https://example.com/next?host=example.com">→</a>
</p>
```

Raw HTML to show up in the webpage as demonstration

<p>
    <a href="https://example.com/previous?host=example.com">←</a>
    <a href="https://example.com/">Example Webring</a>
    <a href="https://example.com/next?host=example.com">→</a>
</p>

# Table of members
<table>
  <tr>
    <th>Member UUID</th>
    <th>Member domain/URI</th>
  </tr>
  {{ . }}
</table>

This webring is powered by [go-webring](https://git.sr.ht/~amolith/go-webring).