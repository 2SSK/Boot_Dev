# REMOVE A SECTION

As I pointed out before, the `webflyx` section is nonsensical because Git doesn't use it for anything. While we _can_ store any key/value pairs we want in our Git configuration, it doesn't mean we _should_.

The `--remove-section` flag is used to remove an entire section from your Git configuration. For example:

    git config --remove-section section
