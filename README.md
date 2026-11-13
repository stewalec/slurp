<p align="center">
  <a href="https://slurp.emohruo.cloud/alec">
    SLURP
  </a>
  <br>
  A simple, multi-tenant feed reader forked from <a href="https://git.j3s.sh/vore">vore</a>.
</p>

## Features

- RSS and Atom support
- Minimal, simple, reliable, fast
- Refresh your feeds automatically
- Display a chronological list of feed items
- Filter by feed name
- Toggle sort by unread items
- Mark as read while you scroll
- Open source & free of charge forever

## Slurp vs vore

Slurp intends to be as minimal and simple as vore, but with a few changes:

- Mark as read while you scroll
- Filter by feed name
- Toggle sort by unread items
- Mark when a feed item has been saved
- Registration disabled via web UI but available instead via the CLI
- Separate "saving" from "archiving" while allowing both functions
- Delete saved/archived items
- UI tweaks


## Building & Running

```
git clone --depth=1 https://git.emohruo.cloud/git/slurp/slurp
cd slurp
go build
./slurp
```

## Dev Notes

### Slurp should always trust websites as the source of authority.

- This is why posts aren't saved to disk - there's no good way to uniquely identify them over time easily. A new website might show up and reuse the post urls, and we want to avoid all of the complexities that introduces by just fetching feeds at runtime and loading them live - that way we're SURE they're fresh and accurate.

### Do not natively display posts.

- Posts always look bad away from their home websites. Instead of doing that, Slurp just takes website snapshots via archive.org and presents them to the user.

### Archived entries will NEVER change/expire.

- If a user uses the "archive" feature, the data they were looking at must never be lost. Therefore, we just copy whatever the active post state was from memory and also snapshot the website via archive.org and link to the snapshot. This way, there's always a cached version available to use.

### Slurp prefers raw URLs.

- We don't care about traditional RSS formats like OPML.

## Misc

Soup icon courtesy of [Hugeicons](https://icon-sets.iconify.design/hugeicons/soup/).