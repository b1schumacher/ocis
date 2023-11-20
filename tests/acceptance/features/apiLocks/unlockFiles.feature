Feature: unlock locked items
  As a user
  I want to unlock the resources previously locked by myself
  So that other users can make changes to the resources

  Background:
    Given user "Alice" has been created with default attributes and without skeleton files


  Scenario Outline: unlock file locked by the user
    Given using <dav-path-version> DAV path
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "textfile0.txt"
    And user "Alice" has locked file "textfile0.txt" setting the following properties
      | lockscope | exclusive |
    When user "Alice" unlocks the last created lock of file "textfile0.txt" using the WebDAV API
    Then the HTTP status code should be "204"
    And 0 locks should be reported for file "textfile0.txt" of user "Alice" by the WebDAV API
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: public tries to unlock a file in a share that was locked by the file owner
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "PARENT"
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "PARENT/parent.txt"
    And user "Alice" has created a public link share with settings
      | path        | PARENT   |
      | permissions | change   |
      | password    | %public% |
    And user "Alice" has locked file "PARENT/parent.txt" setting the following properties
      | lockscope | <lock-scope> |
    When the public unlocks file "/parent.txt" with the last created lock of file "PARENT/parent.txt" of user "Alice" using the WebDAV API
    Then the HTTP status code should be "204"
    And 0 locks should be reported for file "PARENT/parent.txt" of user "Alice" by the WebDAV API
    Examples:
      | dav-path-version | lock-scope |
      | old              | shared     |
      | old              | exclusive  |
      | new              | exclusive  |
      | new              | shared     |
      | spaces           | shared     |
      | spaces           | exclusive  |

  @issue-7599
  Scenario Outline: unlock one of multiple locks set by the user itself
    Given using <dav-path-version> DAV path
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "textfile0.txt"
    And user "Alice" has locked file "textfile0.txt" setting the following properties
      | lockscope | shared |
    And user "Alice" has locked file "textfile0.txt" setting the following properties
      | lockscope | shared |
    When user "Alice" unlocks the last created lock of file "textfile0.txt" using the WebDAV API
    Then the HTTP status code should be "204"
    And 1 locks should be reported for file "textfile0.txt" of user "Alice" by the WebDAV API
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |

  @issue-7638
  Scenario Outline: unlocking a file with the same name as another file in another part of the file system
    Given using <dav-path-version> DAV path
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "textfile0.txt"
    And user "Alice" has created folder "locked"
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "/locked/textfile0.txt"
    And user "Alice" has created folder "notlocked"
    And user "Alice" has uploaded file "filesForUpload/textfile.txt" to "/notlocked/textfile0.txt"
    And user "Alice" has locked file "locked/textfile0.txt" setting the following properties
      | lockscope | <lock-scope> |
    And user "Alice" has locked file "notlocked/textfile0.txt" setting the following properties
      | lockscope | <lock-scope> |
    And user "Alice" has locked file "textfile0.txt" setting the following properties
      | lockscope | <lock-scope> |
    When user "Alice" unlocks the last created lock of file "notlocked/textfile0.txt" using the WebDAV API
    And user "Alice" unlocks the last created lock of file "textfile0.txt" using the WebDAV API
    Then user "Alice" should be able to upload file "filesForUpload/lorem.txt" to "/notlocked/textfile0.txt"
    And user "Alice" should be able to upload file "filesForUpload/lorem.txt" to "/textfile0.txt"
    But user "Alice" should not be able to upload file "filesForUpload/lorem.txt" to "/locked/textfile0.txt"
    Examples:
      | dav-path-version | lock-scope |
      | old              | shared     |
      | old              | exclusive  |
      | new              | shared     |
      | new              | exclusive  |
      | spaces           | shared     |
      | spaces           | exclusive  |

  @issue-7696
  Scenario Outline: unlock a locked file in project space
    Given using spaces DAV path
    And the administrator has assigned the role "Space Admin" to user "Alice" using the Graph API
    And user "Alice" has created a space "project-space" with the default quota using the Graph API
    And user "Alice" has uploaded a file inside space "project-space" with content "some data" to "textfile.txt"
    And user "Alice" has locked file "textfile.txt" inside space "project-space" setting the following properties
      | lockscope | <lock-scope> |
    When user "Alice" unlocks the last created lock of file "textfile.txt" inside space "project-space" using the WebDAV API
    Then the HTTP status code should be "204"
    Examples:
      | lock-scope |
      | shared     |
      | exclusive  |
