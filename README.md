# README

## About

Simple modmanager for recent Anno games (primarily Anno 1800, Anno 117 maybe later) using mod.io's API.
Primarily developed for use under Linux, as the mod management tools integrated in these games don't work properly
 - a version capable of running under Windows may follow.

## Development
To develop on this application
- clone the repository
- `cd <cloned project folder>`

Install the frontend dependencies:
- `cd ./frontend`
- `npm install`

Then run `wails dev` in the project directory. This will run a Vite development
server that will provide very fast hot reload of your frontend changes. If you want to develop in a browser
and have access to your Go methods, there is also a dev server that runs on http://localhost:34115. Connect
to this in your browser, and you can call your Go code from devtools.

## Building

To build a redistributable, production mode package, use `wails build` in the project directory.

Powered by [Wails](https://wails.io/), [SvelteKit](https://svelte.dev) and [Skeleton](https://skeleton.dev)
