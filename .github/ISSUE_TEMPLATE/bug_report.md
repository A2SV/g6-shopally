name: Bug Report
description: Report a problem or unexpected behavior
title: "[Bug] < title here >"
labels: ["bug"]
assignees: []

body:
  - type: textarea
    attributes:
      label: Describe the bug
      description: A clear and concise description of what the bug is.
    validations:
      required: true

  - type: textarea
    attributes:
      label: Steps to Reproduce
      description: What steps did you take that led to the issue?
    validations:
      required: true

  - type: input
    attributes:
      label: Expected behavior
      description: What did you expect to happen?
    validations:
      required: true

  - type: input
    attributes:
      label: Environment
      description: Include device/OS/browser if applicable (e.g., Android 12, Chrome 117)
