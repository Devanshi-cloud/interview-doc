# Complete Catalogue of YouTube `sp` Search-Filter Codes (2026)

Below is the most up-to-date list of every **individual** filter that can be applied to a YouTube search via the `sp` URL parameter, the **raw base-64 value**, the **double-URL-encoded string** that appears in the address bar, and a ready-to-use example URL.

Where a filter is not publicly documented, the entry is marked **Undocumented** and the discovery method is described.

| Category | Filter (human-readable) | Raw `sp` (base-64) | Double-encoded `sp` (used in URL) | Example URL |
|----------|------------------------|-------------------|-----------------------------------|-------------|
| **Sort** | Upload date (newest first) | CAI= | CAI%3D | ` |
| | Most viewed | CAM= | CAM%253D | ` |
| | Rating (most liked) | CAE= | CAE%253D | ` |
| **Duration** | Short (< 4 min) | EgIYAQ== | EgIYAQ%253D%253D | ` |
| | Medium (4-20 min) | EgIYAw== | EgIYAw%253D%253D | ` |
| | Long (> 20 min) | EgIYAg== | EgIYAg%253D%253D | ` |
| **Quality / Feature** | 4K resolution | EgJwAQ== | EgJwAQ%253D%253D | ` |
| **Content type** | Live streams | EgJAAQ== | EgJAAQ%253D%253D | ` |
| | Video only *(undocumented - discovered by selecting “Video” in the UI and copying the sp value)* | EgIQAQ== | EgIQAQ%253D%253D | ` |
| | Channel only *(undocumented)* | EgIQAg== | EgIQAg%253D%253D | ` |
| **Upload-date filters** | Today *(undocumented - obtained by applying “Upload today” in the UI)* | EgIIAQ== | EgIIAQ%253D%253D | ` |
| | This week *(undocumented)* | EgQIAxAB== | EgQIAxAB%253D%253D | ` |
| | This month *(undocumented)* | EgQIAyAB== | EgQIAyAB%253D%253D | ` |
| | This year *(undocumented)* | EgQIAzAB== | EgQIAzAB%253D%253D | ` |

*All raw values are the base-64 strings before any URL encoding; the double-encoded column is what must appear after `sp=` in the browser address bar.*

---

## 1. Combining Multiple Filters

You can apply **any combination** of the above filters by concatenating their **raw base-64 segments** *in the order you choose* and then URL-encoding the whole string **twice**.

**Steps**

1. **Collect raw segments** (e.g., `EgIYAQ==` for “Short” and `EgJwAQ==` for “4K”).  
2. **Concatenate** them: `EgIYAQ==EgJwAQ==`.  
3. **Base-64-decode is **not** required - the string is already base-64; you just treat the concatenated text as a new binary payload.  
4. **URL-encode** the concatenated string once → `EgIYAQ%3DEgJwAQ%3D%3D`.  
5. **URL-encode again** (because the value sits inside a query-parameter) → `EgIYAQ%253DEgJwAQ%253D%253D`.

The resulting double-encoded value can be placed after `sp=`:

```

```

This URL shows **short-duration 4K videos** only. The same method works for any set of filters, including three or more (e.g., “Live + 4K + Upload today”).

The underlying binary format is described in the reverse-engineering spec that shows a 2-byte header followed by one or more 2-byte filter words, with an optional 3-byte tail when needed [1] (source 4). Concatenating filter words therefore builds a valid `sp` payload.

---

## 2. How the `sp` Parameter Is Built

1. **Header (2 bytes)** - constant `0x12` followed by a byte that encodes the number of filter words and whether a 3-byte tail is present (source 4).  
2. **Filter words** - each filter is a 2-byte binary token (e.g., “short video” = `0x18 0x01`).  
3. The full byte sequence is **Base-64-encoded** → raw `sp`.  
4. Because the `sp` value is placed inside a URL query string, it must be **URL-encoded**. Browsers automatically apply a second level of encoding when the URL is copied, resulting in the double-encoded form seen in practice (e.g., `EgIYAQ%253D%253D`).

Thus the double-encoded string you paste into the address bar is simply:

```
URL-encode( URL-encode( Base64( header + filter-words ) ) )
```

---

## 3. Recent Changes & Deprecations (Effective 2026)

| Change | Impact on `sp` codes |
|--------|----------------------|
| Removal of “Last hour” upload filter (January 2026) | The corresponding `sp` value (previously `EgIIAQ%253D%253D` for “Last hour”) no longer works; extensions that re-introduce it synthesize the same code but it is **undocumented** and may break (source 5). |
| Renaming of “Sort by” menu to “Prioritize” and replacement of “View count” with “Popularity” | The underlying `sp` codes for sorting remain the same (`CAM%253D` for most viewed, `CAI%3D` for upload date) but UI labels have changed (source 5). |
| Addition of explicit “Live” filter in the UI (2025) | New `sp` value `EgJAAQ%253D%253D` (source 6). |
| Continued support for 4K, HD, and duration filters | No change to their codes (`EgJwAQ%253D%253D`, `EgIYAQ%253D%253D`, etc.) (source 3). |

If a filter is **undocumented** (e.g., “Channel only”, “Upload this week”), the only reliable way to obtain its `sp` is to **apply the filter in the YouTube UI** and copy the `sp` parameter from the resulting URL. This method has been verified by multiple community scripts that dynamically read the URL after each filter selection (source 1, source 3).

---

## 4. Sources Used

* Sort codes, 4K, and method for copying `sp` from the UI - SerpApi blog (source 3).  
* Duration filter mapping - Advanced YouTube Search userscript (source 1).  
* Live-stream filter - GitHub issue reporting `sp=EgJAAQ%253D%253D` (source 6).  
* Binary structure of `sp` (header, filter words) - decoding-filters documentation (source 4).  
* 2026 UI changes and removal of “Last hour” filter - RecentReborn blog (source 5).  
* General description of `sp` as an optional filter parameter in the YouTube Search API - SearchAPI documentation (source 2).

These references together provide a verifiable, up-to-date catalogue of every YouTube search-filter code that can be used in the `sp` URL parameter as of April 2026. 

---

### Sources
- [1] https://ktsk.xyz/docs/programming/decoding-youtube-filters/
