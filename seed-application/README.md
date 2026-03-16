# Seed Application
Ultimately this should be a docker container that adds the data from `data.md` to the database, by sending a POST request to the backend server.

Please write your own shell script that parses the file and posts each goal to the backend server.

It would be nice if it is possible to pass the definition and values files as parameters to the script.

If an object is **invalid**, it can be skipped. Invalid objects are those that:
- are missing required properties (both `target` and `goal` must always be present)
- reference a value ID that does not exist in the values file

If an object has extra properties, they can simply be ignored. Only `target` and `goal` are required for the POST request.

## Mix & Match Data Format (.mmd/.mmv)
This project uses the "Mix and Match" data format consisting of two separate files:

- Definition file (.mmd)
- Values file (.mmv)

This format separates entries and their values, requiring a lookup to match value IDs to actual data.

### Definition file (`.mmd`)

- Each entry represents a goal object and is identified by an id (t1, t2, t3, etc.)
- Each entry can have multiple properties, such as target and goal.
- Each property references a `value_id`, which points to a value in the `.mmv` file.

**Example goal object:**

```text
entry t1 property target value_id v4
entry t1 property goal   value_id v11
```

The values can be looked up in the corresponding `.mvv` file.

### Values file (`.mvv`)

- Contains the actual values corresponding to each value_id.
- Each value_id must be matched to the entry property in .mmd.

**Example (goals.mmv):**

```text
v11 Candy for all!
v12 Teacher will perform a parody on some internet meme (within reason)
v4 13%
v1 25%
```