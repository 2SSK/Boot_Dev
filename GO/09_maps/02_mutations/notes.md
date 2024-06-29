# MUTATIONS

### INSERT AN ELEMENT

    m[key] = elem

### GET AN ELEMENT

    elem = m[key]

### DELETE AN ELEMENT

    delete(m, key)

### CHECK IF A KEY EXISTS

    elem, ok := m[key]

if `key` is in `m`, then `ok` is `true`. if not, `ok` is `false`.

if `key` is not in the map, then `elem` is the zero value for the map's value for the map's element type.
