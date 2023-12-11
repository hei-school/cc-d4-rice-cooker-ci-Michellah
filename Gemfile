# frozen_string_literal: true

source "https://rubygems.org"


# gem "rails"
gem 'rubocop', require: false

AllCops:
  TargetRubyVersion: 3.0
  Exclude:
    - 'vendor/**/*'

Layout/IndentationStyle:
  Enabled: true
  IndentationWidth: 2
  EnforcedStyle: spaces

Layout/SpaceAroundOperators:
  Enabled: true 
  SpaceAroundOperators:
    - '='
    - '+'
    - '-'
    - '*'
    - '/'
    - '=='

Naming/ClassAndModuleCamelCase:
  Enabled: true

Naming/MethodName:
  Enabled: true
  Style: snake_case 

Naming/ConstantName:
  Enabled: true
  EnforcedStyle: SCREAMING_SNAKE_CASE

Naming/VariableName:
  Enabled: true
  Style: snake_case 

Metrics/LineLength:
    Enabled: true
    Max: 200

Style/For:
    Enabled: true
    EnforcedStyle: each

Style/BlockDelimiters:
    Enabled: true
    EnforcedStyle: { multiline: 'do', single_line: '{}' }
  
  
