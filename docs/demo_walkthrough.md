# Awareness — Demo Walkthrough

A ~5 minute live demo. Two computers, same setup script.

Two repos involved:
- **Test project** (the repo we demo on): `mustafatamyapar/super-cool-test-project`
- **Awareness tool** (the CLI we built): `stg-tud/impl-ss26-pf3-better-awareness-for-git`

## Setup (run on each computer)

```bash
awareness demo-setup
cd /tmp/demo/repo
awareness daemon --watch --privacy full 30s &
```

That's it. It clones the test project, installs the hook, and tells you what to do next.
All demo steps below run from `/tmp/demo/repo`.

Both computers are now watching for file changes, auto-publishing, and pulling every 30s.

---

## 1. Person A starts working

Just edit and save a file.

```bash
printf '\n// refactoring auth\n' >> main.go
```

Say: *"No command needed. The background watcher publishes what files they
touched and line ranges, never the code itself."*

## 2. Person B touches the same file — gets a notification

```bash
printf '\n// adding logging\n' >> main.go
```

Wait up to 30s. A desktop notification pops up showing the overlap.
Check manually with:

```bash
awareness status --once
```

Point at the `!` marker.
Say: *"They find out while still editing, not at push time."*

## 3. Two-level conflict detection (pre-push hook)

Person B commits and pushes to a throwaway branch:

```bash
git checkout -b demo-push
git commit -am "demo change"
git push -u origin demo-push
```

The pre-push hook fires and shows two levels:

- **CONFLICT** if both edited overlapping lines:
  ```
  awareness: CONFLICT — overlapping line changes with a teammate:
    Ali (full, updated 1m ago):
      ✗ main.go  (you: L40-55, them: L42-60)

  Push will continue.
  ```

- **Warning** if same file but different sections:
  ```
  awareness: warning — same area, but no detected line conflict:
    Ali (full, updated 1m ago):
      ~ main.go

  Push will continue.
  ```

Say: *"It warns, it does not block. And it tells you whether it's a real
line conflict or just the same file."*

Cleanup: `git checkout main && git branch -D demo-push && git push origin --delete demo-push`

## 4. AI summary (local Ollama)

Person A makes changes and summarizes with the local model:

```bash
printf '\n// new feature\n' >> cute/cute.go
awareness summarize --publish --privacy full
```

Person B edits the same file and pushes:

```bash
printf '\n// fix cute\n' >> cute/cute.go
git checkout -b demo-push2
git commit -am "demo change 2"
git push -u origin demo-push2
```

Expected output (note the quoted AI summary):

```
awareness: CONFLICT — overlapping line changes with a teammate:
  Ali (full, updated 1m ago):
    "Adding a new feature to the cute module"
    ✗ cute/cute.go  (you: L12-15, them: L10-18)

Push will continue.
```

Say: *"The AI summary runs locally with Ollama, no data leaves the machine."*

## 5. Privacy levels

Person A stops the daemon (Ctrl-C or `kill %1`), then:

```bash
awareness publish --privacy full       # exact file + real name
awareness publish --privacy standard   # folder only + real name
awareness publish --privacy anonymous  # folder only + hashed alias
```

Person B checks after each: `awareness status --once`

Say: *"A slider from exact file+name down to folder+alias."*

## 6. Heatmap (optional)

```bash
awareness heatmap
```

Opens an HTML view of team hotspots.

## 7. Wrap up

On both computers:

```bash
awareness reset
```

---

## Talking points

- **daemon --watch** — awareness is always on, zero babysitting
- **Two-level detection** — real line conflicts vs same-file warnings
- **AI summary** — local Ollama model, no data leaves the machine
- **Privacy levels** — how much to reveal is a choice
- **Desktop notifications** — find out the moment work collides
- **status + heatmap** — visual overview

## Who built what

- **Ali:** daemon, file watcher, privacy levels, two-level conflict detection
- **Mustafa:** pre-push hook, status, heatmap, AI summarize
