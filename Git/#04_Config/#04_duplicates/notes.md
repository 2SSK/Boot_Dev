# DUPLICATES

Typically, ina key/value store, like a Python dictionary, you aren't allowed to have duplicate keys. Strangely enough, Git doesn't care.

### UNSET ALL

The `--unset-all` flag is useful if you ever _really_ want to purge all instances of a key from your configuration. Conversely, the `--unset` flag ony works with a single instance of a key.

    git config --unset-all example.key
