# Test Checklist

Run the same setup on two machines (or two terminals). Each person enters their name.

## Setup

Run this on each machine. It asks for your name so both people run the same script.

```bash
read -p "Your name: " NAME
EMAIL=$(echo "$NAME" | tr '[:upper:]' '[:lower:]')@e.com
rm -rf /tmp/demo && mkdir -p /tmp/demo
git clone -q https://github.com/mustafatamyapar/super-cool-test-project.git /tmp/demo/repo
cd /tmp/demo/repo
git config user.name "$NAME"
git config user.email "$EMAIL"
BIN="/Users/amiraliebrahimi/Desktop/IMPL Project/impl-ss26-pf3-better-awareness-for-git/awareness"
"$BIN" install-hook
```

## Tests

### 1. Publish + Status

```bash
# Person A: edit main.go, publish
printf '\n// test\n' >> main.go
"$BIN" publish --privacy full

# Person B: check status
"$BIN" status --once
```

Expected: Person B sees Person A and `main.go` with `!` marker.

### 2. Two-level conflict (CONFLICT vs warning)

```bash
# Person B: edit same file, push
printf '\n// overlap\n' >> main.go
git checkout -b test && git commit -am "test" && git push -u origin test
```

Expected: pre-push hook prints CONFLICT or warning depending on line overlap.

Cleanup: `git checkout main && git branch -D test && git push origin --delete test`

### 3. AI summary (Ollama)

```bash
printf '\n// new stuff\n' >> cute/cute.go
"$BIN" summarize --privacy full
```

Expected: prints a one-line summary from local qwen2.5. No API key needed.

### 4. AI summary in pre-push output

```bash
# Person A: publish with summary
"$BIN" summarize --publish --privacy full

# Person B: edit same file, push
printf '\n// fix\n' >> cute/cute.go
git checkout -b test2 && git commit -am "test2" && git push -u origin test2
```

Expected: pre-push output shows quoted AI summary under Person A's name.

### 5. Daemon + watch + notification

```bash
# Both: start watcher
"$BIN" daemon --watch --privacy full 30s &

# Person A: save a file
printf '\n// auto\n' >> main.go

# Person B: edit same file after ~30s (next pull)
printf '\n// auto2\n' >> main.go
```

Expected: desktop notification pops up on Person B showing overlap.

Kill daemons: `kill %1` in both terminals.

### 6. Privacy levels

```bash
# Person A publishes at each level, Person B checks after each
"$BIN" publish --privacy full       # exact file + real name
"$BIN" publish --privacy standard   # folder only + real name
"$BIN" publish --privacy anonymous  # folder only + hashed alias
```

Person B checks each with: `"$BIN" status --once`

### 7. Unit tests

```bash
cd "/Users/amiraliebrahimi/Desktop/IMPL Project/impl-ss26-pf3-better-awareness-for-git"
go test ./cmd/awareness/... -count=1
```

Expected: all tests pass.
