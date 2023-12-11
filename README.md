[![Review Assignment Due Date](https://classroom.github.com/assets/deadline-readme-button-24ddc0f5d75046c5622901739e7c5dd533143b0c8e959d652212380cedb1ea36.svg)](https://classroom.github.com/a/PHq8Kfj_)

# Ruby Project Setup Guide

This guide outlines the steps to create and manage a Ruby project using Bundler.

## Installation

### Step 1: Install Ruby

Ensure that Ruby is installed on your system. Verify by running `ruby -v` in your terminal. If Ruby is not installed, you can download and install it from the [official Ruby website](https://www.ruby-lang.org/en/downloads/) or use a version manager like RVM or rbenv.

### Step 2: Create Project Directory

Create a new directory to serve as the main folder for your Ruby project. Use the command line (`mkdir project_name`) or your file explorer to create this directory.

### Step 3: Initialize the Project with Bundler

Navigate to your project directory in the terminal and initialize a new Ruby project using Bundler:

    ```bash
    bundle init

### Step 4: Add Dependencies to Gemfile

Edit the `Gemfile` to include the required gems (libraries) for your project. Open the `Gemfile` using a text editor and add the following lines:

    ```ruby
    # Gemfile
    source 'https://rubygems.org'
    gem 'gem_name_1'
    gem 'gem_name_2'
    # Add other gems as needed

### Step 6: Running the Code

To execute your Ruby code, open a terminal or command prompt, navigate to your project directory, and run the following command:

    ```bash
    ruby your_main_file.rb

## Linter Setup

For maintaining code quality and adherence to coding standards, you can integrate a linter into your Ruby project.

### Using RuboCop

[RuboCop](https://github.com/rubocop/rubocop) is a popular linter and code analyzer for Ruby. To set it up:

1. Install RuboCop by running:
   ```bash
   gem install rubocop

2. Create a .rubocop.yml file in the root directory of your project. You can customize the rules and configurations as needed. 

3. Run RuboCop to analyze your code: 
    ```bash
    rubocop

You' re ready to use ruby !