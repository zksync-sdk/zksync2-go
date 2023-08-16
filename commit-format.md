# Commit Message Format

*This specification is inspired by and supersedes
the [AngularJS commit message format](https://gist.github.com/brianclements/841ea7bffdb01346392c).*

We have very precise rules over how our Git commit messages must be formatted.
This format leads to **easier to read commit history**.

Each commit message consists of a **header**, a **body**, and a **footer**.

```
<header>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

The `header` is mandatory and must conform to the [Commit Message Header](#commit-message-header) format.

The `body` is optional. The [Commit Message Body](#commit-message-body) format describes what the footer is used for and
the structure it must have.

The `footer` is optional. The [Commit Message Footer](#commit-message-footer) format describes what the footer is used 
for and the structure it must have.

#### Commit Message Header

```
<type>(<scope>): <short summary>
  │       │             │
  │       │             └─⫸ Summary in present tense. Not capitalized. No period at the end.
  │       │
  │       └─⫸ Commit Scope: package that is modified
  │
  └─⫸ Commit Type: build|ci|docs|feat|fix|perf|refactor|test
```

The `<type>` and `<summary>` fields are mandatory, the `(<scope>)` field is optional.

##### Type

Must be one of the following:

* **build**: Changes that affect the build system or external dependencies (example scopes: make).
* **ci**: Changes to our CI configuration files and scripts (examples: GitHub, GitLab).
* **docs**: Documentation only changes.
* **feat**: A new feature.
* **fix**: A bug fix.
* **perf**: A code change that improves performance.
* **refactor**: A code change that neither fixes a bug nor adds a feature.
* **test**: Adding missing tests or correcting existing tests.

***chore** type is replaced in favor of **build** and **ci**.*

##### Scope

The scope should be the name of the package affected (examples: `types`, `utils`, `clients`, `accounts`).

##### Summary

Use the summary field to provide a succinct description of the change:

* use the imperative, present tense: "change" not "changed" nor "changes"
* don't capitalize the first letter
* no dot (.) at the end

#### Commit Message Body

Just as in the summary, use the imperative, present tense: "fix" not "fixed" nor "fixes".

Explain the motivation for the change in the commit message body. This commit message should explain _why_ you are
making the change.
You can include a comparison of the previous behavior with the new behavior in order to illustrate the impact of the
change.

#### Commit Message Footer

The footer can contain information about breaking changes and deprecations and is also the place to reference GitHub
issues, Jira tickets, and other PRs that this commit closes or is related to.
For example:

```
BREAKING CHANGE: <breaking change description + migration instructions>
<BLANK LINE>
Fixes #<issue number>
```

or

```
DEPRECATED: <deprecation description + recommended update path>
<BLANK LINE>
Closes #<pr number>
```

Breaking Change section should start with the phrase "BREAKING CHANGE: " followed by a detailed description of the 
breaking change that also includes migration instructions.

Similarly, a Deprecation section should start with "DEPRECATED: " followed by a detailed description of the deprecation 
that also mentions the recommended update path.

### Revert commits

If the commit reverts a previous commit, it should begin with `revert: `, followed by the header of the reverted commit.

The content of the commit message body should contain:

- information about the SHA of the commit being reverted in the following format: `This reverts commit <SHA>`,
- a clear description of the reason for reverting the commit message.


### Additional resources

To determine which formats are valid and which are not, please refer to the [following documentation](https://github.com/conventional-changelog/commitlint/tree/master/@commitlint/config-conventional#type-enum) 
for examples.