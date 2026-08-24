# Awareness — Demo Walkthrough

A ~5 minute live demo. Two computers, same setup script.

## Setup (run on each computer)

Paste this once. It asks for your name, clones the repo, builds the binary, and starts the watcher.

```bash
read -p "Your name: " NAME
EMAIL=$(echo "$NAME" | tr '[:upper:]' '[:lower:]')@e.com
rm -rf /tmp/demo && mkdir -p /tmp/demo
git clone -q https://github.com/mustafatamyapar/super-cool-test-project.git /tmp/demo/repo
cd /tmp/demo/repo
git config user.name "$NAME"
git config user.email "$EMAIL"
git clone -q https://github.com/stg-tud/impl-ss26-pf3-better-awareness-for-git.git /tmp/demo/tool
cd /tmp/demo/tool && go build -o /tmp/demo/aw ./cmd/awareness
cd /tmp/demo/repo
alias aw=/tmp/demo/aw
aw install-hook
aw daemon --watch --privacy full 30s &
echo "Ready. Working in /tmp/demo/repo as $NAME"
```

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
aw status --once
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
aw summarize --publish --privacy full
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
aw publish --privacy full       # exact file + real name
aw publish --privacy standard   # folder only + real name
aw publish --privacy anonymous  # folder only + hashed alias
```

Person B checks after each: `aw status --once`

Say: *"A slider from exact file+name down to folder+alias."*

## 6. Heatmap (optional)

```bash
aw heatmap
```

Opens an HTML view of team hotspots.

## 7. Wrap up

On both computers:

```bash
aw reset
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
