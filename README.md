## Google I/O 2015 web app

### Setup

1. `git clone https://github.com/GoogleChrome/ioweb2015.git`
2. `cd ioweb2015`
3. `npm install`
4. `gulp setup`

If you plan on modifying source code, be a good citizen and:

1. Install [EditorConfig plugin](http://editorconfig.org/#download) for your favourite browser.
   The plugin should automatically pick up the [.editorconfig](.editorconfig) settings.
2. Obey the pre-commit hook that's installed as part of `gulp setup`.
   It will check for JavaScript and code style errors before committing to the `master` branch.

To run a backend server you'll need:

1. [Go 1.4](https://golang.org/dl/).
2. Optional: [gcloud tool](https://cloud.google.com/sdk/#Quick_Start)
   and [app](https://cloud.google.com/sdk/gcloud-app#Installation) component
   to run and deploy GAE-based backend.

### Running

Start any static web server in `app/` or `gulp serve` to leverage live-reload,
serve with a standalone backend via `gulp serve:backend`
or using App Engine dev server: `gulp serve:gae`.

**Note**: If you're using a static server, you'll have to run `gulp` or `gulp sass` at least once
to generate CSS from the .scss files.

### Building

Run `gulp`. Then hit `http://localhost:<PORT>/dist/app/`. The unbuilt version is still viewable at `http://localhost:<PORT>/app/` but will not contain minfied JS or vulcanized HTML Imports.

**Note**: Build won't succeed if either `gulp jshint` or `gulp jscs` reports errors.

## Backend

`gulp backend` will build a self-sufficient backend server and place the binary in `backend/bin/server`.

`gulp backend:test` will run backend server tests. If, while working on the backend, you feel tired
of running the command again and again, use `gulp backend:test --watch` to watch for file changes
and re-run tests automatically.

To deploy complete application on App Engine:

1. Run `gulp` which will build both frontend and backend in `dist` directory.
2. Run `gcloud preview app deploy [--version <v>] dist/backend`.
