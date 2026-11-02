<p style="text-align:center;">
  <a href="https://slurp.emohruo.cloud/alec" style="display:inline-flex;font-family:monospace;font-size:24px;">
    SL
    <span>
      <svg xmlns="http://www.w3.org/2000/svg" width="1em" height="1em" viewBox="0 0 24 24">
        <path d="M0 0h24v24H0z" fill="none" />
        <g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5">
          <path d="M7.08 17.576c-1.312-.718-3.156-2.126-3.86-4.573c-.232-.805-.348-1.207-.046-1.605S3.975 11 4.973 11h14.038c.998 0 1.497 0 1.8.398c.301.398.186.8-.046 1.605c-.704 2.446-2.548 3.854-3.86 4.572l-.153.084c-.436.256-.6.677-.45 1.158l.038.11c.026.076.04.114.049.144a1.502 1.502 0 0 1-1.385 1.928L14.85 21h-5.7c-.09 0-.134 0-.168-.002a1.5 1.5 0 0 1-1.384-1.913l.053-.159l.038-.115c.145-.477-.018-.893-.45-1.148z" />
          <path stroke-linecap="round" d="m11.492 3l.5.5a1.414 1.414 0 0 1 0 2a1.414 1.414 0 0 0 0 2l.5.5m-6-5l.5.5a1.414 1.414 0 0 1 0 2a1.414 1.414 0 0 0 0 2l.5.5m8.5 3l3-6" />
        </g>
      </svg>
    </span>
    RP
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
git clone https://github.com/stewalec/slurp
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