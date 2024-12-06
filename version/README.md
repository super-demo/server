# How to manage version number of the API

Latest editted on Jan 17, 2024

### Concept

The idea is to use this file to manage version of API combined with the ability of git hook script (pre-commit).

This file will be checked by git hook (pre-commit) and auto increment the Revision if it needed.

> Condition : `pre-commit` script will allow you to commit if this file is in the staged file.

- If it isn't in the staged file (no change), then pre-commit script will automatically increase the revision version's value by default.

### Version explanation

> `constants.go` in ${ProjectRoot}/constants
```
const (
	// This value should be increased when there is a new major release.
	MajorVersion = 1

	// This value should be increased when there is a new feature / change under the current major version.
	MinorVersion = 0

	// This value will be increased when there is a fix for a current minor version (no new functionality).
	RevisionVersion = 0
)
```

### Situations
1.) Dev want to implement/fix bug and commit
```
    What to do :    Just run "git commit" one time to let the script works for you.
                    OR
                    Dev manually increase the RevisionVersion by one.
                    OR
                    Run the "increase_revision_version.sh" script in "${ProjectRoot}/constants/scripts"
    Expected   :    1.) This file must be updated by increasing value of RevisionVersion by one
                    2.) This file must be staged.
```

2.) Dev want to increase minor version.
```
    What to do  :   Run the "increase_minor_version.sh" script in "${ProjectRoot}/constants/scripts"
                    OR
                    Dev manually increase the MinorVersion by one and reset the RevisionVersion to zero.
    Expected    :   1.) This file must be updated by increasing value of MinorVersion by one.
                    2.) RevisionVersion is resetted to zero.
                    3.) This file must be staged.
```

3.) Dev want to increase major version.
```
    What to do :    Run the "increase_minor_version.sh" script in "${ProjectRoot}/constants/scripts"
                    OR
                    Dev manually increase the MajorVersion by one and reset both MinorVersion and RevisionVersion to zero.
    Expected    :   1.) This file must be updated by increasing value of MinorVersion by one.
                    2.) Both MinorVersion and RevisionVersion are resetted to zero.
                    3.) This file must be staged.
```