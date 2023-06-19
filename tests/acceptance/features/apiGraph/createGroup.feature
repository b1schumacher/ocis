@api
Feature: create group
  As an admin
  I want to create a group
  So that I can add users to the group

  Background:
    Given user "Alice" has been created with default attributes and without skeleton files
    And the administrator has assigned the role "Admin" to user "Alice" using the Graph API


  Scenario Outline: admin user creates a group
    When user "Alice" creates a group "<groupname>" using the Graph API
    Then the HTTP status code should be "200"
    And group "<groupname>" should exist
    Examples:
    | groupname       |
    | simplegroup     |
    | España§àôœ€     |
    | नेपाली            |
    | $x<=>[y*z^2+1]! |
    | 😅 😆           |
    | comma,grp1      |
    | Finance (NP)    |
    | slash\Middle    |

  @issue-3516
  Scenario: admin user tries to create a group that already exists
    Given group "mygroup" has been created
    When user "Alice" tries to create a group "mygroup" using the Graph API
    Then the HTTP status code should be "400"
    And group "mygroup" should exist

  @issue-5938
  Scenario Outline: user other than the admin can't create a group
    Given user "Brian" has been created with default attributes and without skeleton files
    And the administrator has assigned the role "<userRole>" to user "Brian" using the Graph API
    When user "Brian" tries to create a group "mygroup" using the Graph API
    Then the HTTP status code should be "403"
    And group "mygroup" should not exist
    Examples:
      | userRole    |
      | Space Admin |
      | User        |
      | User Light  |

  @issue-5050
  Scenario: admin user tries to create a group that is the empty string
    When user "Alice" tries to create a group "" using the Graph API
    Then the HTTP status code should be "400"
