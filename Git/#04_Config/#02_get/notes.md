# GET

We've used `--list` to see _all_ the configuration vlues, but the `--get` flag is useful for getting a single value.

    git config --get <key>

Keys are in the format `<section>.<keyname>`. For example:

- `user.name`
- `webflyx.ceo`
