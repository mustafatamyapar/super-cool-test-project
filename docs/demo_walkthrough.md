# Awareness — Demo Walkthrough

A ~5 minute live demo. Two computers, same setup script.

The test project is embedded in the binary. No internet needed.
Uses your existing `git config user.name`.

## Setup (run on each computer)

```bash
cd impl-ss26-pf3-better-awareness-for-git
go build -o awareness ./cmd/awareness
./awareness demo-setup
```

Then follow the commands it prints. It will tell you exactly where to `cd` and how to start the daemon.

Both computers are now watching for file changes, auto-publishing, and pulling every 30s.

---

## 1. Person A starts working

Open `cute/cute.go` in an editor and add a line inside the `Hello` function (around line 7):

```go
fmt.Println("greeting", name)
```

Save the file.

Say: *"No command needed. The background watcher publishes what files they
touched and line ranges, never the code itself."*

## 2. Person B touches the same file -- gets a notification

Open `cute/cute.go` and edit the same `Hello` function (around line 10):

```go
name = "valued guest"
```

Save the file.

Wait up to 30s. A desktop notification pops up showing the overlap.
Check manually with:

```bash
awareness status --once
```

Point at the `!` marker.
Say: *"They find out while still editing, not at push time."*

## 3. Two-level conflict detection (pre-push hook)

Person B commits and pushes (the hook runs even on a local push):

```bash
git checkout -b demo-push
git commit -am "demo change"
git push
```

The pre-push hook fires and shows two levels:

- **CONFLICT** if both edited overlapping lines:
  ```
  awareness: CONFLICT — overlapping line changes with a teammate:
    <Person A's name> (full, updated 1m ago):
      ✗ main.go  (you: L40-55, them: L42-60)

  Push will continue.
  ```

- **Warning** if same file but different sections:
  ```
  awareness: warning — same area, but no detected line conflict:
    <Person A's name> (full, updated 1m ago):
      ~ main.go

  Push will continue.
  ```

Say: *"It warns, it does not block. And it tells you whether it's a real
line conflict or just the same file."*

## 4. AI summary (local Ollama)

Person A edits `main.go` (add a new print line), saves, then summarizes:

```bash
awareness summarize --publish --privacy full
```

Person B checks status:

```bash
awareness status --once
```

Expected: Person A's entry shows a quoted AI summary.

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

- **daemon --watch** -- awareness is always on, zero babysitting
- **Two-level detection** -- real line conflicts vs same-file warnings
- **AI summary** -- local Ollama model, no data leaves the machine
- **Privacy levels** -- how much to reveal is a choice
- **Desktop notifications** -- find out the moment work collides
- **status + heatmap** -- visual overview

## Who built what

- **Ali:** daemon, file watcher, privacy levels, two-level conflict detection
- **Mustafa:** pre-push hook, status, heatmap, AI summarize
