# TimeScape

TimeScape is a Git-inspired version control engine built in Go to deeply understand how modern version control systems work internally using immutable objects, snapshot-based history, and content-addressable storage.

---

## Why TimeScape?

Most developers use Git daily but never truly understand how it works under the hood.

TimeScape was built as a systems engineering project to explore:

- Content-addressable object storage
- Immutable snapshots
- Tree-based filesystem representation
- Commit graph traversal
- Snapshot restoration
- Version history management
- CLI tooling architecture

Instead of treating Git as a black box, TimeScape rebuilds its core ideas from scratch in Go.

---
# ⏳ TimeScape

> A Git-inspired version control engine built from scratch in Go.

## 🎥 Demo

📺 Watch the complete project walkthrough:

https://youtu.be/O26ys6P3A7k


# Core Architecture

TimeScape follows a simplified Git-inspired object model:

```txt
Working Files
      ↓
Blob Objects
      ↓
Index / Staging
      ↓
Tree Snapshots
      ↓
Commit Objects
      ↓
Refs / History
```

---

# Object Model

## Blob Object

Blob objects store immutable file contents.

Example internal structure:

```txt
blob <size>\0<raw_content>
```

Blobs do not store:
- filenames
- directory names
- metadata

They only store file content.

---

## Tree Object

Tree objects represent filesystem snapshots.

A tree maps:

```txt
filename → blob hash
```

This allows reconstruction of repository state at any point in history.

---

## Commit Object

Commit objects represent historical snapshots.

Each commit stores:

- Tree hash
- Parent commit hash
- Commit message
- Timestamp

Example structure:

```txt
tree <treeHash>
parent <parentHash>
message <message>
timestamp <unix_time>
```

This creates a parent-linked commit graph used for history traversal.

---

# Features Implemented

- Repository initialization
- SHA1 content hashing
- Blob object creation
- Object database storage
- Tree snapshot generation
- Index/staging system
- Commit history
- Parent-linked traversal
- Commit log system
- Snapshot checkout / restoration
- CLI command architecture

---

# Commands

## Initialize Repository

```bash
timescape init
```

---

## Add File To Staging

```bash
timescape add <file>
```

---

## Create Commit

```bash
timescape commit "message"
```

---

## View Commit History

```bash
timescape log
```

---

## Restore Historical Snapshot

```bash
timescape checkout <commitHash>
```

---

# Internal Project Structure

```txt
internal/
├── hashing/
├── objects/
├── repository/
├── storage/
├── utils/
```

---

# Example Workflow

```bash
timescape init

timescape add test.txt

timescape commit "initial commit"

timescape log

timescape checkout <commitHash>
```

---

# Technical Concepts Explored

- Immutable data architecture
- Snapshot-based version control
- Filesystem modeling
- Content-addressable storage
- Commit graph traversal
- Serialization/deserialization
- CLI tooling systems
- Persistent object databases

---

# Planned Features

- Branching support
- Merge support
- Recursive tree structures
- Diff engine
- Safer checkout previews
- Visual commit graph
- Better repository inspection tools

---

# Built With

- Go
- Standard Library
- Filesystem APIs
- SHA1 hashing

---

# Project Goal

TimeScape is not intended to fully replace Git.

The goal is to deeply understand and explore the architecture behind modern version control systems by rebuilding the core ideas from first principles.
