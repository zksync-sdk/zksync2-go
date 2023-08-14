module.exports =
    {
        "extends": [
            "@commitlint/config-conventional"
        ],
        "rules": {
            "type-enum": [2, "always", ["feat", "fix", "chore"]],
            "footer-max-line-length": [0, "always"],
            "header-max-length": [2, "always", 120],
            "subject-case": [0]
        }
    }
