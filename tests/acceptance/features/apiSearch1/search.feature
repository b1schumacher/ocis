Feature: Search
  As a user
  I want to search for resources in the space
  So that I can get them quickly

  Background:
    Given these users have been created with default attributes and without skeleton files:
      | username |
      | Alice    |
      | Brian    |
    And using spaces DAV path
    And the administrator has assigned the role "Space Admin" to user "Alice" using the Graph API
    And user "Alice" has created a space "project101" with the default quota using the Graph API
    And user "Alice" has created a folder "folderMain/SubFolder1/subFOLDER2" in space "project101"
    And user "Alice" has uploaded a file inside space "project101" with content "some content" to "folderMain/SubFolder1/subFOLDER2/insideTheFolder.txt"


  Scenario Outline: user can search items from the project space
    Given using <dav-path-version> DAV path
    When user "Alice" searches for "*SUB*" using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result should contain "2" entries
    And the search result of user "Alice" should contain these entries:
      | /folderMain/SubFolder1            |
      | /folderMain/SubFolder1/subFOLDER2 |
    But the search result of user "Alice" should not contain these entries:
      | /folderMain                                           |
      | /folderMain/SubFolder1/subFOLDER2/insideTheFolder.txt |
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: user can search items from the shares
    Given using <dav-path-version> DAV path
    And user "Alice" has sent the following share invitation:
      | resource        | folderMain |
      | space           | project101 |
      | sharee          | Brian      |
      | shareType       | user       |
      | permissionsRole | Viewer     |
    When user "Brian" searches for "*folder*" using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result should contain "4" entries
    And the search result of user "Brian" should contain these entries:
      | folderMain/SubFolder1                                |
      | folderMain/SubFolder1/subFOLDER2                     |
      | folderMain/SubFolder1/subFOLDER2/insideTheFolder.txt |
    And for user "Brian" the search result should contain space "mountpoint/folderMain"
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: user can search hidden files
    Given using <dav-path-version> DAV path
    And user "Alice" has created a folder ".space" in space "project101"
    When user "Alice" searches for "*.sp*" using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result should contain "1" entries
    And the search result of user "Alice" should contain these entries:
      | /.space |
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: user cannot search pending share
    Given user "Brian" has disabled auto-accepting
    And using <dav-path-version> DAV path
    And user "Alice" has sent the following share invitation:
      | resource        | folderMain |
      | space           | project101 |
      | sharee          | Brian      |
      | shareType       | user       |
      | permissionsRole | Viewer     |
    When user "Brian" searches for "*folder*" using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result should contain "0" entries
    And the search result of user "Brian" should not contain these entries:
      | /SubFolder1                                |
      | /SubFolder1/subFOLDER2                     |
      | /SubFolder1/subFOLDER2/insideTheFolder.txt |
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: user cannot search declined share
    Given using <dav-path-version> DAV path
    And user "Alice" has sent the following share invitation:
      | resource        | folderMain |
      | space           | project101 |
      | sharee          | Brian      |
      | shareType       | user       |
      | permissionsRole | Viewer     |
    And user "Brian" has disabled sync of last shared resource
    When user "Brian" searches for "*folder*" using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result should contain "0" entries
    And the search result of user "Brian" should not contain these entries:
      | /SubFolder1                                |
      | /SubFolder1/subFOLDER2                     |
      | /SubFolder1/subFOLDER2/insideTheFolder.txt |
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: user cannot search deleted items
    Given using <dav-path-version> DAV path
    And user "Alice" has removed the folder "folderMain" from space "project101"
    When user "Alice" searches for "*folderMain*" using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result should contain "0" entries
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: user can search project space by name
    Given using <dav-path-version> DAV path
    When user "Alice" searches for '*project101*' using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result should contain "1" entries
    And for user "Alice" the search result should contain space "project101"
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: user can search inside folder in space
    Given using <dav-path-version> DAV path
    When user "Alice" searches for "*folder*" inside folder "/folderMain" in space "project101" using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result should contain "3" entries
    And the search result of user "Alice" should contain only these entries:
      | folderMain/SubFolder1                                |
      | folderMain/SubFolder1/subFOLDER2                     |
      | folderMain/SubFolder1/subFOLDER2/insideTheFolder.txt |
    But the search result of user "Alice" should not contain these entries:
      | /folderMain |
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: search inside folder in shares
    Given using <dav-path-version> DAV path
    And user "Alice" has sent the following share invitation:
      | resource        | folderMain |
      | space           | project101 |
      | sharee          | Brian      |
      | shareType       | user       |
      | permissionsRole | Viewer     |
    When user "Brian" searches for "*folder*" inside folder "/folderMain" in space "Shares" using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result of user "Brian" should contain only these entries:
      | folderMain/SubFolder1                                |
      | folderMain/SubFolder1/subFOLDER2                     |
      | folderMain/SubFolder1/subFOLDER2/insideTheFolder.txt |
    But the search result of user "Brian" should not contain these entries:
      | /folderMain |
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  @skipOnStable3.0
  Scenario Outline: search files inside the folder
    Given using <dav-path-version> DAV path
    And user "Alice" has uploaded file with content "hello world inside root" to "file1.txt"
    And user "Alice" has created folder "/Folder"
    And user "Alice" has uploaded file with content "hello world inside folder" to "/Folder/file2.txt"
    And user "Alice" has created folder "/Folder/SubFolder"
    And user "Alice" has uploaded file with content "hello world inside sub-folder" to "/Folder/SubFolder/file3.txt"
    When user "Alice" searches for "*file*" inside folder "/Folder" using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result of user "Alice" should contain only these entries:
      | /Folder/file2.txt           |
      | /Folder/SubFolder/file3.txt |
    But the search result of user "Alice" should not contain these entries:
      | file1.txt |
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |

  @issue-7114
  Scenario Outline: search files inside the folder with white space character in its name
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "/New Folder"
    And user "Alice" has uploaded file with content "hello world inside folder" to "/New Folder/file.txt"
    And user "Alice" has created folder "/New Folder/Sub Folder"
    And user "Alice" has uploaded file with content "hello world inside sub folder" to "/New Folder/Sub Folder/file1.txt"
    When user "Alice" searches for "*file*" inside folder "/New Folder" using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result of user "Alice" should contain only these entries:
      | /New Folder/file.txt             |
      | /New Folder/Sub Folder/file1.txt |
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |

  @issue-7114
  Scenario Outline: search files with white space character in its name
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "/New Folder"
    And user "Alice" has uploaded file with content "hello world" to "/new file.txt"
    And user "Alice" has created folder "/New Folder/New Sub Folder"
    When user "Alice" searches for "*new*" using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result of user "Alice" should contain only these entries:
      | /New Folder                |
      | /New Folder/New Sub Folder |
      | /new file.txt              |
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |

  @issue-enterprise-6000 @issue-7028 @issue-7092
  Scenario Outline: sharee cannot find resources that are not shared
    Given using <dav-path-version> DAV path
    And user "Alice" has created a folder "foo/sharedToBrian" in space "Alice Hansen"
    And user "Alice" has created a folder "sharedToCarol" in space "Alice Hansen"
    And user "Alice" has sent the following share invitation:
      | resource        | foo      |
      | space           | Personal |
      | sharee          | Brian    |
      | shareType       | user     |
      | permissionsRole | Viewer   |
    When user "Brian" searches for "shared*" using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result of user "Brian" should contain these entries:
      | foo/sharedToBrian |
    But the search result of user "Brian" should not contain these entries:
      | /sharedToCarol |
    Examples:
      | dav-path-version |
      | old              |
      | new              |
      | spaces           |


  Scenario Outline: search resources using different search patterns (KQL feature)
    Given using spaces DAV path
    And user "Alice" has created a folder "subfolder" in space "project101"
    When user "Alice" searches for '<pattern>' using the WebDAV API
    Then the HTTP status code should be "207"
    And the search result should contain "1" entries
    And the search result of user "Alice" should contain these entries:
      | <search-result> |
    Examples:
      | pattern      | search-result                     | description                     |
      | fold*        | /folderMain                       | starts with                     |
      | *der1        | /folderMain/SubFolder1            | ends with                       |
      | subfolder    | /subfolder                        | exact search                    |
      | name:*der2   | /folderMain/SubFolder1/subFOLDER2 | patern 'name:''                 |
      | name:"*der2" | /folderMain/SubFolder1/subFOLDER2 | pattern 'name:""' (with quotes) |

  @issue-7812 @issue-8442
  Scenario: try to search with invalid patterns
    Given using spaces DAV path
    And user "Alice" has uploaded file with content "test file" to "testFile.txt"
    When user "Alice" searches for 'AND mediatype:document' using the WebDAV API
    Then the HTTP status code should be "400"
    And the value of the item "/d:error/s:message" in the response should be "error: bad request: the expression can't begin from a binary operator: 'AND'"
