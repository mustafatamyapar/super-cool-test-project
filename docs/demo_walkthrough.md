# Awareness — Demo Walkthrough

A ~5 minute live demo on the real test project
(`super-cool-test-project`). Two terminals, two users: **Ali** and **Mustafa**.

## Setup (once, before the meeting)

Build the binary:

```bash
cd "/Users/amiraliebrahimi/Desktop/IMPL Project/impl-ss26-pf3-better-awareness-for-git"
go build -o /tmp/awareness ./cmd/awareness
```

Two clones of the test repo live in `/tmp/demo` (`pc1` and `pc2`).
If they are missing or you want a clean slate, run the reset block at the bottom.

Open two terminals:

- **Terminal 1 (Ali):** `cd /tmp/demo/pc1 && alias aw=/tmp/awareness`
- **Terminal 2 (Mustafa):** `cd /tmp/demo/pc2 && alias aw=/tmp/awareness`

The project contains `main.go`, `cute/cute.go`, `go.mod`, `README.md`.

---

## 0. Turn awareness on (background)

In **both** terminals:

```bash
aw daemon --watch --privacy full 30s
```

The watcher publishes automatically on save and pulls teammates every 30s.
Desktop notifications appear when a new overlap is detected.

## 1. Mustafa starts working

**Mustafa (Terminal 2):** just edit and save.

```bash
printf '\n// mustafa: refactor\n' >> main.go
```

Say: *"No command needed. The background watcher publishes a summary of
the files he touched and line ranges, never the code itself."*

## 2. Ali touches the same file — gets a notification

**Ali (Terminal 1):**

```bash
printf '\n// ali: add logging\n' >> main.go
```

Wait for the next pull (up to 30s). A desktop notification pops up showing the
overlap. Check it with:

```bash
aw status --once
```

Point at the `!` marker and `(+lines -lines)`.
Say: *"Ali finds out while he is still editing, not at push time."*

## 3. Two-level conflict detection (pre-push hook)

**Ali (Terminal 1):** commit and push to a throwaway branch.

```bash
git checkout -b ali-demo
git commit -am "ali: add logging"
git push -u origin ali-demo
```

The pre-push hook fires and shows two levels:

- **CONFLICT** if both edited overlapping lines:
  ```
  awareness: CONFLICT — overlapping line changes with a teammate:
    Mustafa (full, updated 1m ago):
      ✗ main.go  (you: L40-55, them: L42-60)

  Push will continue.
  ```

- **Warning** if same file but different sections:
  ```
  awareness: warning — same area, but no detected line conflict:
    Mustafa (full, updated 1m ago):
      ~ main.go

  Push will continue.
  ```

Say: *"It warns, it does not block. And it tells you whether it's a real
line conflict or just the same file."*

Cleanup after the demo (optional):
```bash
git checkout main && git branch -D ali-demo && git push origin --delete ali-demo
```

## 4. AI summary (local Ollama)

**Mustafa (Terminal 2):** make changes and summarize with the local model.

```bash
printf '\n// mustafa: new feature\n' >> cute/cute.go
aw summarize --publish --privacy full
```

This calls the local Ollama model (qwen2.5) and attaches a summary note.

**Ali (Terminal 1):** edit an overlapping file and push.

```bash
printf '\n// ali: fix cute\n' >> cute/cute.go
git checkout -b ali-demo2
git commit -am "ali: fix cute"
git push -u origin ali-demo2
```

Expected output (note the quoted AI summary):

```
awareness: CONFLICT — overlapping line changes with a teammate:
  Mustafa (full, updated 1m ago):
    "Adding a new feature to the cute module"
    ✗ cute/cute.go  (you: L12-15, them: L10-18)

Push will continue.
```

Say: *"The AI summary runs locally with Ollama, no data leaves the machine.
Teammates see a one-line description of what you're working on."*

## 5. Privacy levels

Stop the watcher on Mustafa first so we can publish at chosen levels:

**Mustafa (Terminal 2):** Ctrl-C the daemon, then:

```bash
aw publish --privacy standard
```

**Ali (Terminal 1):**

```bash
aw status --once
```

Shows folder instead of exact file:
```
  Mustafa (standard, updated 1m ago):
    ! cute/
```

Then anonymous:

```bash
aw publish --privacy anonymous
```

Shows hashed alias instead of name:
```
  anon-xxxxxxxx (anonymous, updated 1m ago):
    ! cute/
```

Say: *"This is the answer to 'how much do we reveal' — a slider from exact
file+name down to folder+alias."*

## 6. Heatmap (optional finisher)

**Ali (Terminal 1):**

```bash
aw heatmap
```

Opens an HTML view of team hotspots (files shaded by number of contributors).

## 7. Wrap up

In both terminals:

```bash
aw reset
```

---

## Talking points (map to what Fabian asked)

- **daemon --watch** — awareness is always on, zero babysitting.
- **Two-level detection** — distinguishes real line conflicts from same-file warnings.
- **AI summary** — local Ollama model describes what each person is working on.
- **Privacy levels** — answer his "how much to hide or show" question.
- **Desktop notifications** — you find out the moment work starts colliding.
- **status + heatmap** — the "more visual" direction.

## Who built what

- **Ali:** publishing side — daemon / awareness branch, file watcher, privacy levels, line counts, two-level conflict detection.
- **Mustafa:** consuming side — pre-push hook, `status` + auto-pull, heatmap, AI summarize.

---

## Reset to a clean state (re-run anytime)

```bash
BIN=/tmp/awareness
rm -rf /tmp/demo && mkdir -p /tmp/demo && cd /tmp/demo
git clone -q https://github.com/mustafatamyapar/super-cool-test-project.git pc1
cd pc1 && git config user.name Ali && git config user.email ali@e.com
cd /tmp/demo
git clone -q https://github.com/mustafatamyapar/super-cool-test-project.git pc2
cd pc2 && git config user.name Mustafa && git config user.email mustafa@e.com
```
