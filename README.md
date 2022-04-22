# Breathe

This CLI utility written in Go helps in breathing better. It's inspired by [the book Breath from James Nestor](https://www.mrjamesnestor.com/).

The main purpose of this project is for me to learn Go.

## Breathing exercises

### Ideal breathing

```bash
breathe ideal
```

### Box breathing

```bash
breathe box
```

### Box breathing (long variant)

```bash
breathe box-long
```

### Long breathing training

This command is for training having longer breaths.

```bash
breathe long
breathe long --inhaleStartSeconds=4 --inhaleEndSeconds=10 --cyclesPerStartSeconds=3
```

## CLI options

### Sound

By default, the sound in the CLI is disabled. You can enable it with the `--sound` option.

```bash
breathe ideal --sound=all
```
