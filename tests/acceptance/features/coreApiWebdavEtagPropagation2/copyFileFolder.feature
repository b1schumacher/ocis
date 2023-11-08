Feature: propagation of etags when copying files or folders
  As a client app
  I want metadata (etags) of parent folders to change when a file of folder is copied
  So that the client app can know to re-scan and sync the content of the folder(s)

  Background:
    Given user "Alice" has been created with default attributes and without skeleton files

  @issue-4251
  Scenario Outline: copying a file does not change its etag
    Given using <dav-path-version> DAV path
    And user "Alice" has uploaded file with content "uploaded content" to "file.txt"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/file.txt"
    And user "Alice" has stored etag of element "/file.txt" on path "/renamedFile.txt"
    When user "Alice" copies file "/file.txt" to "/renamedFile.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should not have changed:
      | user  | path      |
      | Alice | /file.txt |
    And these etags should have changed:
      | user  | path             |
      | Alice | /                |
      | Alice | /renamedFile.txt |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

    @skipOnRevaMaster
    Examples:
      | dav-path-version |
      | spaces           |

  @issue-4251
  Scenario Outline: copying a file inside a folder changes its etag
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "/folder"
    And user "Alice" has uploaded file with content "uploaded content" to "file.txt"
    And user "Alice" has stored etag of element "/file.txt"
    And user "Alice" has stored etag of element "/folder"
    And user "Alice" has stored etag of element "/file.txt" on path "/folder/renamedFile.txt"
    When user "Alice" copies file "/file.txt" to "/folder/renamedFile.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should not have changed:
      | user  | path      |
      | Alice | /file.txt |
    And these etags should have changed:
      | user  | path                    |
      | Alice | /folder/renamedFile.txt |
      | Alice | /folder                 |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

    @skipOnRevaMaster
    Examples:
      | dav-path-version |
      | spaces           |

  @issue-4251
  Scenario Outline: copying a file from one folder to an other changes the etags of destination
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "/src"
    And user "Alice" has uploaded file with content "uploaded content" to "/src/file.txt"
    And user "Alice" has created folder "/dst"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/src"
    And user "Alice" has stored etag of element "/dst"
    When user "Alice" copies folder "/src/file.txt" to "/dst/file.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path |
      | Alice | /    |
      | Alice | /dst |
    And these etags should not have changed:
      | user  | path |
      | Alice | /src |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

    @skipOnRevaMaster
    Examples:
      | dav-path-version |
      | spaces           |

  @issue-4251
  Scenario Outline: copying a file into a subfolder changes the etags of all parents
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "/upload"
    And user "Alice" has created folder "/upload/sub"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/file.txt"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/file.txt"
    And user "Alice" has stored etag of element "/upload/file.txt" on path "/upload/sub/file.txt"
    And user "Alice" has stored etag of element "/upload/sub"
    When user "Alice" copies file "/upload/file.txt" to "/upload/sub/file.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path                 |
      | Alice | /                    |
      | Alice | /upload              |
      | Alice | /upload/sub          |
      | Alice | /upload/sub/file.txt |
    And these etags should not have changed:
      | user  | path             |
      | Alice | /upload/file.txt |
    @issue-4091
    Examples:
      | dav-path-version |
      | old              |
      | new              |

    @skipOnRevaMaster
    Examples:
      | dav-path-version |
      | spaces           |

  @issue-4251
  Scenario Outline: copying a file inside a publicly shared folder by public changes etag for the sharer
    Given using <dav-path-version> DAV path
    And user "Alice" has created folder "/upload"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/file.txt"
    And user "Alice" has created a public link share with settings
      | path        | upload |
      | permissions | change |
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/file.txt"
    And user "Alice" has stored etag of element "/upload/file.txt" on path "/upload/renamedFile.txt"
    When the public copies file "file.txt" to "/renamedFile.txt" using the new public WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path                    |
      | Alice | /                       |
      | Alice | /upload                 |
      | Alice | /upload/renamedFile.txt |
    And these etags should not have changed:
      | user  | path             |
      | Alice | /upload/file.txt |
    @issue-4091
    Examples:
      | dav-path-version |
      | old              |
      | new              |

    @skipOnRevaMaster
    Examples:
      | dav-path-version |
      | spaces           |

  @skipOnReva
  Scenario Outline: sharee copying a file inside a folder changes its etag for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav-path-version> DAV path
    And user "Alice" has created folder "/upload"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/file.txt"
    And user "Alice" has shared folder "/upload" with user "Brian"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/file.txt"
    And user "Alice" has stored etag of element "/upload/file.txt" on path "/upload/renamed.txt"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/upload"
    And user "Brian" has stored etag of element "/Shares/upload/file.txt"
    And user "Brian" has stored etag of element "/Shares/upload/file.txt" on path "/Shares/upload/renamed.txt"
    When user "Brian" copies file "/Shares/upload/file.txt" to "/Shares/upload/renamed.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path                       |
      | Alice | /                          |
      | Alice | /upload                    |
      | Alice | /upload/renamed.txt        |
      | Brian | /                          |
      | Brian | /Shares                    |
      | Brian | /Shares/upload             |
      | Brian | /Shares/upload/renamed.txt |
    And these etags should not have changed:
      | user  | path                    |
      | Alice | /upload/file.txt        |
      | Brian | /Shares/upload/file.txt |
    Examples:
      | dav-path-version |
      | old              |
      | new              |

  @issue-4251 @skipOnReva
  Scenario Outline: sharer copying a file inside a folder changes its etag for all collaborators
    Given user "Brian" has been created with default attributes and without skeleton files
    And using <dav-path-version> DAV path
    And user "Alice" has created folder "/upload"
    And user "Alice" has uploaded file with content "uploaded content" to "/upload/file.txt"
    And user "Alice" has shared folder "/upload" with user "Brian"
    And user "Alice" has stored etag of element "/"
    And user "Alice" has stored etag of element "/upload"
    And user "Alice" has stored etag of element "/upload/file.txt"
    And user "Alice" has stored etag of element "/upload/file.txt" on path "/upload/renamed.txt"
    And user "Brian" has stored etag of element "/"
    And user "Brian" has stored etag of element "/Shares"
    And user "Brian" has stored etag of element "/Shares/upload"
    And user "Brian" has stored etag of element "/Shares/upload/file.txt"
    And user "Brian" has stored etag of element "/Shares/upload/file.txt" on path "/Shares/upload/renamed.txt"
    When user "Alice" copies file "/upload/file.txt" to "/upload/renamed.txt" using the WebDAV API
    Then the HTTP status code should be "201"
    And these etags should have changed:
      | user  | path                       |
      | Alice | /                          |
      | Alice | /upload                    |
      | Alice | /upload/renamed.txt        |
      | Brian | /                          |
      | Brian | /Shares                    |
      | Brian | /Shares/upload             |
      | Brian | /Shares/upload/renamed.txt |
    And these etags should not have changed:
      | user  | path                    |
      | Alice | /upload/file.txt        |
      | Brian | /Shares/upload/file.txt |
    Examples:
      | dav-path-version |
      | old              |
      | new              |