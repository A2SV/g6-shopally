name: ✨ Feature Request
description: Suggest a new idea or improvement
title: "[Feature] < title here >"
labels: ["enhancement"]
assignees: []

body:
  - type: textarea
    attributes:
      label: Describe the feature
      description: What do you want to add or improve?
    validations:
      required: true

  - type: textarea
    attributes:
      label: Why is this needed?
      description: Explain the benefit of this feature or the problem it solves.

  - type: textarea
    attributes:
      label: Additional context
      description: Add any other details or mockups here (optional).
