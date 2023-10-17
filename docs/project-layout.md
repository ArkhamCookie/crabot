# Project Layout

## `/.github`

GitHub meta files.

## `/docs`

Documentation

### `/docs/examples`

Examples

## `/internal`

Private application and library code.

> [!IMPORTANT]
> This is not meant to be used by any other projects, so you can **not** rely on it in your projects.

## `/lib`

This contain library code that's ok to use by external applications.
While a lot of Golang programmers use `/pkg`, I use `/lib` for two reasons.
First, it makes it clear that you **can** use this in your project.
The other reason is I've just used lib for most of my programming projects;
I don't like cluttering my root directory for a project.

<!-- ## `/web` -->
