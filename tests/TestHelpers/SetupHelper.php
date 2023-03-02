<?php declare(strict_types=1);
/**
 * ownCloud
 *
 * @author Artur Neumann <artur@jankaritech.com>
 * @copyright Copyright (c) 2017 Artur Neumann artur@jankaritech.com
 *
 * This code is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License,
 * as published by the Free Software Foundation;
 * either version 3 of the License, or any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>
 *
 */
namespace TestHelpers;

use Behat\Testwork\Hook\Scope\HookScope;
use GuzzleHttp\Exception\GuzzleException;
use GuzzleHttp\Exception\ServerException;
use Exception;
use Psr\Http\Message\ResponseInterface;
use SimpleXMLElement;

/**
 * Helper to setup UI / Integration tests
 *
 * @author Artur Neumann <artur@jankaritech.com>
 *
 */
class SetupHelper extends \PHPUnit\Framework\Assert {
	/**
	 * @var string
	 */
	private static $ocPath = null;
	/**
	 * @var string
	 */
	private static $baseUrl = null;
	/**
	 * @var string
	 */
	private static $adminUsername = null;
	/**
	 * @var string
	 */
	private static $adminPassword = null;

	/**
	 * creates a user
	 *
	 * @param string|null $userName
	 * @param string|null $password
	 * @param string|null $xRequestId
	 * @param string|null $displayName
	 * @param string|null $email
	 *
	 * @return string[] associated array with "code", "stdOut", "stdErr"
	 * @throws Exception
	 */
	public static function createUser(
		?string $userName,
		?string   $password,
		?string  $xRequestId = '',
		?string  $displayName = null,
		?string   $email = null
	):array {
		$occCommand = ['user:add', '--password-from-env'];
		if ($displayName !== null) {
			$occCommand = \array_merge($occCommand, ["--display-name", $displayName]);
		}
		if ($email !== null) {
			$occCommand = \array_merge($occCommand, ["--email", $email]);
		}
		\putenv("OC_PASS=" . $password);
		return ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
	}

	/**
	 * deletes a user
	 *
	 * @param string|null $userName
	 * @param string|null $xRequestId
	 *
	 * @return string[] associated array with "code", "stdOut", "stdErr"
	 * @throws Exception
	 */
	public static function deleteUser(
		?string $userName,
		?string $xRequestId = ''
	):array {
		return ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
	}

	/**
	 *
	 * @param string|null $userName
	 * @param string|null $app
	 * @param string|null $key
	 * @param string|null $value
	 * @param string|null $xRequestId
	 *
	 * @return string[]
	 * @throws Exception
	 */
	public static function changeUserSetting(
		?string $userName,
		?string $app,
		?string $key,
		?string $value,
		?string $xRequestId = ''
	):array {
		return ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
	}

	/**
	 * creates a group
	 *
	 * @param string|null $groupName
	 * @param string|null $xRequestId
	 *
	 * @return string[] associated array with "code", "stdOut", "stdErr"
	 * @throws Exception
	 */
	public static function createGroup(
		?string $groupName,
		?string $xRequestId = ''
	):array {
		return ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
	}

	/**
	 * adds an existing user to a group, creating the group if it does not exist
	 *
	 * @param string|null $groupName
	 * @param string|null $userName
	 * @param string|null $xRequestId
	 *
	 * @return string[] associated array with "code", "stdOut", "stdErr"
	 * @throws Exception
	 */
	public static function addUserToGroup(
		?string $groupName,
		?string $userName,
		?string $xRequestId = ''
	):array {
		return ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
	}

	/**
	 * removes a user from a group
	 *
	 * @param string|null $groupName
	 * @param string|null $userName
	 * @param string|null $xRequestId
	 *
	 * @return string[] associated array with "code", "stdOut", "stdErr"
	 * @throws Exception
	 */
	public static function removeUserFromGroup(
		?string $groupName,
		?string $userName,
		?string $xRequestId = ''
	):array {
		return ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
	}

	/**
	 * deletes a group
	 *
	 * @param string|null $groupName
	 * @param string|null $xRequestId
	 *
	 * @return string[] associated array with "code", "stdOut", "stdErr"
	 * @throws Exception
	 */
	public static function deleteGroup(
		?string $groupName,
		?string $xRequestId = ''
	):array {
		return ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
	}

	/**
	 *
	 * @param string|null $xRequestId
	 *
	 * @return string[]
	 * @throws Exception
	 */
	public static function getGroups(
		?string $xRequestId = ''
	):array {
		return \json_decode(
			['code' => '', 'stdOut' => '', 'stdErr' => '' ]['stdOut']
		);
	}
	/**
	 *
	 * @param HookScope $scope
	 *
	 * @return array of suite context parameters
	 */
	public static function getSuiteParameters(HookScope $scope):array {
		return $scope->getEnvironment()->getSuite()
			->getSettings() ['context'] ['parameters'];
	}

	/**
	 * Fixup OC path so that it always starts with a "/" and does not end with
	 * a "/".
	 *
	 * @param string|null $ocPath
	 *
	 * @return string
	 */
	private static function normaliseOcPath(?string $ocPath):string {
		return '/' . \trim($ocPath, '/');
	}

	/**
	 *
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 * @param string|null $baseUrl
	 * @param string|null $ocPath
	 *
	 * @return void
	 */
	public static function init(
		?string $adminUsername,
		?string $adminPassword,
		?string $baseUrl,
		?string $ocPath
	): void {
		foreach (\func_get_args() as $variableToCheck) {
			if (!\is_string($variableToCheck)) {
				throw new \InvalidArgumentException(
					"mandatory argument missing or wrong type ($variableToCheck => "
					. \gettype($variableToCheck) . ")"
				);
			}
		}
		self::$adminUsername = $adminUsername;
		self::$adminPassword = $adminPassword;
		self::$baseUrl = \rtrim($baseUrl, '/');
		self::$ocPath = self::normaliseOcPath($ocPath);
	}

	/**
	 *
	 * @return string path to the testing app occ endpoint
	 * @throws Exception if ocPath has not been set yet
	 */
	public static function getOcPath():?string {
		if (self::$ocPath === null) {
			throw new Exception(
				"getOcPath called before ocPath is set by init"
			);
		}

		return self::$ocPath;
	}

	/**
	 *
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 * @param string|null $xRequestId
	 *
	 * @return SimpleXMLElement
	 * @throws GuzzleException
	 */
	public static function getSysInfo(
		?string $baseUrl,
		?string $adminUsername,
		?string $adminPassword,
		?string $xRequestId = ''
	):SimpleXMLElement {
		$result = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			"GET",
			"/apps/testing/api/v1/sysinfo",
			$xRequestId
		);
		if ($result->getStatusCode() !== 200) {
			throw new \Exception(
				"could not get sysinfo " . $result->getReasonPhrase()
			);
		}
		return HttpRequestHelper::getResponseXml($result, __METHOD__)->data;
	}

	/**
	 *
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 * @param string|null $xRequestId
	 *
	 * @return string
	 * @throws GuzzleException
	 */
	public static function getServerRoot(
		?string $baseUrl,
		?string $adminUsername,
		?string $adminPassword,
		?string $xRequestId = ''
	):string {
		$sysInfo = self::getSysInfo(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			$xRequestId
		);
		// server_root is a SimpleXMLElement object that "wraps" a string.
		/// We want the bare string.
		return (string) $sysInfo->server_root;
	}

	/**
	 * @param string|null $adminUsername
	 * @param string|null $callerName
	 *
	 * @return string
	 * @throws Exception
	 */
	private static function checkAdminUsername(?string $adminUsername, ?string $callerName):?string {
		if (self::$adminUsername === null
			&& $adminUsername === null
		) {
			throw new Exception(
				"$callerName called without adminUsername - pass the username or call SetupHelper::init"
			);
		}
		if ($adminUsername === null) {
			$adminUsername = self::$adminUsername;
		}
		return $adminUsername;
	}

	/**
	 * @param string|null $adminPassword
	 * @param string|null $callerName
	 *
	 * @return string
	 * @throws Exception
	 */
	private static function checkAdminPassword(?string $adminPassword, ?string $callerName):string {
		if (self::$adminPassword === null
			&& $adminPassword === null
		) {
			throw new Exception(
				"$callerName called without adminPassword - pass the password or call SetupHelper::init"
			);
		}
		if ($adminPassword === null) {
			$adminPassword = self::$adminPassword;
		}
		return $adminPassword;
	}

	/**
	 * @param string|null $baseUrl
	 * @param string|null $callerName
	 *
	 * @return string
	 * @throws Exception
	 */
	private static function checkBaseUrl(?string $baseUrl, ?string $callerName):?string {
		if (self::$baseUrl === null
			&& $baseUrl === null
		) {
			throw new Exception(
				"$callerName called without baseUrl - pass the baseUrl or call SetupHelper::init"
			);
		}
		if ($baseUrl === null) {
			$baseUrl = self::$baseUrl;
		}
		return $baseUrl;
	}

	/**
	 *
	 * @param string|null $dirPathFromServerRoot e.g. 'apps2/myapp/appinfo'
	 * @param string|null $xRequestId
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 *
	 * @return void
	 * @throws GuzzleException
	 * @throws Exception
	 */
	public static function mkDirOnServer(
		?string $dirPathFromServerRoot,
		?string $xRequestId = '',
		?string $baseUrl = null,
		?string $adminUsername = null,
		?string $adminPassword = null
	):void {
		$baseUrl = self::checkBaseUrl($baseUrl, "mkDirOnServer");
		$adminUsername = self::checkAdminUsername($adminUsername, "mkDirOnServer");
		$adminPassword = self::checkAdminPassword($adminPassword, "mkDirOnServer");
		$result = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			"POST",
			"/apps/testing/api/v1/dir",
			$xRequestId,
			['dir' => $dirPathFromServerRoot]
		);

		if ($result->getStatusCode() !== 200) {
			throw new \Exception(
				"could not create directory $dirPathFromServerRoot " . $result->getReasonPhrase()
			);
		}
	}

	/**
	 *
	 * @param string|null $dirPathFromServerRoot e.g. 'apps2/myapp/appinfo'
	 * @param string|null $xRequestId
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 *
	 * @return void
	 * @throws GuzzleException
	 * @throws Exception
	 */
	public static function rmDirOnServer(
		?string $dirPathFromServerRoot,
		?string $xRequestId = '',
		?string $baseUrl = null,
		?string $adminUsername = null,
		?string $adminPassword = null
	):void {
		$baseUrl = self::checkBaseUrl($baseUrl, "rmDirOnServer");
		$adminUsername = self::checkAdminUsername($adminUsername, "rmDirOnServer");
		$adminPassword = self::checkAdminPassword($adminPassword, "rmDirOnServer");
		$result = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			"DELETE",
			"/apps/testing/api/v1/dir",
			$xRequestId,
			['dir' => $dirPathFromServerRoot]
		);

		if ($result->getStatusCode() !== 200) {
			throw new \Exception(
				"could not delete directory $dirPathFromServerRoot " . $result->getReasonPhrase()
			);
		}
	}

	/**
	 *
	 * @param string|null $filePathFromServerRoot e.g. 'app2/myapp/appinfo/info.xml'
	 * @param string|null $content
	 * @param string|null $xRequestId
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 *
	 * @return void
	 * @throws GuzzleException
	 */
	public static function createFileOnServer(
		?string $filePathFromServerRoot,
		?string $content,
		?string $xRequestId = '',
		?string $baseUrl = null,
		?string $adminUsername = null,
		?string $adminPassword = null
	):void {
		$baseUrl = self::checkBaseUrl($baseUrl, "createFileOnServer");
		$adminUsername = self::checkAdminUsername($adminUsername, "createFileOnServer");
		$adminPassword = self::checkAdminPassword($adminPassword, "createFileOnServer");
		$result = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			"POST",
			"/apps/testing/api/v1/file",
			$xRequestId,
			[
				'file' => $filePathFromServerRoot,
				'content' => $content
			]
		);

		if ($result->getStatusCode() !== 200) {
			throw new \Exception(
				"could not create file $filePathFromServerRoot " . $result->getReasonPhrase()
			);
		}
	}

	/**
	 *
	 * @param string|null $filePathFromServerRoot e.g. 'app2/myapp/appinfo/info.xml'
	 * @param string|null $xRequestId
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 *
	 * @return void
	 * @throws GuzzleException
	 * @throws Exception
	 */
	public static function deleteFileOnServer(
		?string $filePathFromServerRoot,
		?string $xRequestId = '',
		?string $baseUrl = null,
		?string $adminUsername = null,
		?string $adminPassword = null
	):void {
		$baseUrl = self::checkBaseUrl($baseUrl, "deleteFileOnServer");
		$adminUsername = self::checkAdminUsername($adminUsername, "deleteFileOnServer");
		$adminPassword = self::checkAdminPassword($adminPassword, "deleteFileOnServer");
		$result = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			"DELETE",
			"/apps/testing/api/v1/file",
			$xRequestId,
			[
				'file' => $filePathFromServerRoot
			]
		);

		if ($result->getStatusCode() !== 200) {
			throw new \Exception(
				"could not delete file $filePathFromServerRoot " . $result->getReasonPhrase()
			);
		}
	}

	/**
	 * @param string|null $fileInCore e.g. 'app2/myapp/appinfo/info.xml'
	 * @param string|null $xRequestId
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 *
	 * @return string
	 * @throws GuzzleException
	 * @throws Exception
	 */
	public static function readFileFromServer(
		?string $fileInCore,
		?string $xRequestId = '',
		?string $baseUrl  = null,
		?string $adminUsername = null,
		?string $adminPassword = null
	):string {
		$baseUrl = self::checkBaseUrl($baseUrl, "readFile");
		$adminUsername = self::checkAdminUsername(
			$adminUsername,
			"readFile"
		);
		$adminPassword = self::checkAdminPassword(
			$adminPassword,
			"readFile"
		);

		$response = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			'GET',
			"/apps/testing/api/v1/file?file={$fileInCore}",
			$xRequestId
		);
		self::assertSame(
			200,
			$response->getStatusCode(),
			"Failed to read the file {$fileInCore}"
		);
		$localContent = HttpRequestHelper::getResponseXml($response, __METHOD__);
		$localContent = (string)$localContent->data->element->contentUrlEncoded;
		return \urldecode($localContent);
	}

	/**
	 * returns the content of a file in a skeleton folder
	 *
	 * @param string|null $fileInSkeletonFolder
	 * @param string|null $xRequestId
	 * @param string|null $baseUrl
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 *
	 * @return string content of the file
	 * @throws GuzzleException
	 * @throws Exception
	 */
	public static function readSkeletonFile(
		?string $fileInSkeletonFolder,
		?string $xRequestId = '',
		?string $baseUrl = null,
		?string $adminUsername = null,
		?string $adminPassword = null
	):string {
		$baseUrl = self::checkBaseUrl($baseUrl, "readSkeletonFile");
		$adminUsername = self::checkAdminUsername(
			$adminUsername,
			"readSkeletonFile"
		);
		$adminPassword = self::checkAdminPassword(
			$adminPassword,
			"readSkeletonFile"
		);

		//find the absolute path of the currently set skeletondirectory
		$occResponse = ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
		if ((int) $occResponse['code'] !== 0) {
			throw new \Exception(
				"could not get current skeletondirectory. " . $occResponse['stdErr']
			);
		}
		$skeletonRoot = \trim($occResponse['stdOut']);

		$fileInSkeletonFolder = \rawurlencode("$skeletonRoot/$fileInSkeletonFolder");
		$response = OcsApiHelper::sendRequest(
			$baseUrl,
			$adminUsername,
			$adminPassword,
			'GET',
			"/apps/testing/api/v1/file?file={$fileInSkeletonFolder}&absolute=true",
			$xRequestId
		);
		self::assertSame(
			200,
			$response->getStatusCode(),
			"Failed to read the file {$fileInSkeletonFolder}"
		);
		$localContent = HttpRequestHelper::getResponseXml($response, __METHOD__);
		$localContent = (string)$localContent->data->element->contentUrlEncoded;
		$localContent = \urldecode($localContent);
		return $localContent;
	}

	/**
	 * enables an app
	 *
	 * @param string|null $appName
	 * @param string|null $xRequestId
	 *
	 * @return string[] associated array with "code", "stdOut", "stdErr"
	 * @throws Exception
	 */
	public static function enableApp(
		?string $appName,
		?string $xRequestId = ''
	):array {
		return ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
	}

	/**
	 * disables an app
	 *
	 * @param string|null $appName
	 * @param string|null $xRequestId
	 *
	 * @return string[] associated array with "code", "stdOut", "stdErr"
	 * @throws Exception
	 */
	public static function disableApp(
		?string $appName,
		?string $xRequestId = ''
	):array {
		return ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
	}

	/**
	 * checks if an app is currently enabled
	 *
	 * @param string|null $appName
	 * @param string|null $xRequestId
	 *
	 * @return bool true if enabled, false if disabled or nonexistent
	 * @throws Exception
	 */
	public static function isAppEnabled(
		?string $appName,
		?string $xRequestId = ''
	):bool {
		$result = ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
		return \strtolower(\substr($result['stdOut'], 0, 7)) === 'enabled';
	}

	/**
	 * lists app status (enabled or disabled)
	 *
	 * @param string|null $appName
	 * @param string|null $xRequestId
	 *
	 * @return bool true if the app needed to be enabled, false otherwise
	 * @throws Exception
	 */
	public static function enableAppIfNotEnabled(
		?string $appName,
		?string $xRequestId = ''
	):bool {
		if (!self::isAppEnabled($appName, $xRequestId)) {
			self::enableApp(
				$appName,
				$xRequestId
			);
			return true;
		}

		return false;
	}

	/**
	 * Runs a list of occ commands at once
	 *
	 * @param array|null $commands
	 * @param string|null $xRequestId
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 * @param string|null $baseUrl
	 *
	 * @return array
	 * @throws GuzzleException
	 * @throws Exception
	 */
	public static function runBulkOcc(
		?array $commands,
		?string $xRequestId = '',
		?string $adminUsername = null,
		?string $adminPassword = null,
		?string $baseUrl = null
	):array {
		return [];
	}

	/**
	 * @param string $baseUrl
	 * @param string $user
	 * @param string $password
	 * @param string $xRequestId
	 *
	 * @return ResponseInterface|null
	 * @throws GuzzleException
	 */
	public static function resetOpcache(
		?string $baseUrl,
		?string $user,
		?string $password,
		?string $xRequestId = ''
	):?ResponseInterface {
		try {
			return OcsApiHelper::sendRequest(
				$baseUrl,
				$user,
				$password,
				"DELETE",
				"/apps/testing/api/v1/opcache",
				$xRequestId
			);
		} catch (ServerException $e) {
			echo "could not reset opcache, if tests fail try to set " .
				"'opcache.revalidate_freq=0' in the php.ini file\n";
		}
		return null;
	}

	/**
	 * Create local storage mount
	 *
	 * @param string|null $mount (name of local storage mount)
	 * @param string|null $xRequestId
	 *
	 * @return string[] associated array with "code", "stdOut", "stdErr", "storageId"
	 * @throws GuzzleException
	 */
	public static function createLocalStorageMount(
		?string $mount,
		?string $xRequestId = ''
	):array {
		$mountPath = TEMPORARY_STORAGE_DIR_ON_REMOTE_SERVER . "/$mount";
		SetupHelper::mkDirOnServer(
			$mountPath,
			$xRequestId
		);
		// files_external:create requires absolute path
		$serverRoot = self::getServerRoot(
			self::$baseUrl,
			self::$adminUsername,
			self::$adminPassword,
			$xRequestId
		);
		$result = ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
		// stdOut should have a string like "Storage created with id 65"
		$storageIdWords = \explode(" ", \trim($result['stdOut']));
		if (\array_key_exists(4, $storageIdWords)) {
			$result['storageId'] = (int)$storageIdWords[4];
		} else {
			// presumably something went wrong with the files_external:create command
			// so return "unknown" to the caller. The result array has the command exit
			// code and stdErr output etc., so the caller can process what it likes
			// of that information to work out what went wrong.
			$result['storageId'] = "unknown";
		}
		return $result;
	}

	/**
	 * Get a system config setting, including status code, output and standard
	 * error output.
	 *
	 * @param string|null $key
	 * @param string|null $xRequestId
	 * @param string|null $output e.g. json
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 * @param string|null $baseUrl
	 * @param string|null $ocPath
	 *
	 * @return string[] associated array with "code", "stdOut", "stdErr"
	 * @throws GuzzleException
	 */
	public static function getSystemConfig(
		?string $key,
		?string $xRequestId = '',
		?string $output = null,
		?string $adminUsername = null,
		?string $adminPassword = null,
		?string $baseUrl = null,
		?string $ocPath = null
	):array {
		$args = [];
		$args[] = 'config:system:get';
		$args[] = $key;

		if ($output !== null) {
			$args[] = '--output';
			$args[] = $output;
		}

		$args[] = '--no-ansi';

		return ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
	}

	/**
	 * Set a system config setting
	 *
	 * @param string|null $key
	 * @param string|null $value
	 * @param string|null $xRequestId
	 * @param string|null $type e.g. boolean or json
	 * @param string|null $output e.g. json
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 * @param string|null $baseUrl
	 * @param string|null $ocPath
	 *
	 * @return string[] associated array with "code", "stdOut", "stdErr"
	 * @throws GuzzleException
	 */
	public static function setSystemConfig(
		?string $key,
		?string $value,
		?string $xRequestId = '',
		?string $type = null,
		?string $output = null,
		?string $adminUsername = null,
		?string $adminPassword = null,
		?string $baseUrl = null,
		?string $ocPath = null
	):array {
		$args = [];
		$args[] = 'config:system:set';
		$args[] = $key;
		$args[] = '--value';
		$args[] = $value;

		if ($type !== null) {
			$args[] = '--type';
			$args[] = $type;
		}

		if ($output !== null) {
			$args[] = '--output';
			$args[] = $output;
		}

		$args[] = '--no-ansi';
		if ($baseUrl === null) {
			$baseUrl = self::$baseUrl;
		}
		return ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
	}

	/**
	 * Get the value of a system config setting
	 *
	 * @param string|null $key
	 * @param string|null $xRequestId
	 * @param string|null $output e.g. json
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 * @param string|null $baseUrl
	 * @param string|null $ocPath
	 *
	 * @return string
	 * @throws GuzzleException if parameters have not been provided yet or the testing app is not enabled
	 */
	public static function getSystemConfigValue(
		?string $key,
		?string $xRequestId = '',
		?string $output = null,
		?string $adminUsername = null,
		?string $adminPassword = null,
		?string $baseUrl = null,
		?string $ocPath = null
	):string {
		if ($baseUrl === null) {
			$baseUrl = self::$baseUrl;
		}
		return self::getSystemConfig(
			$key,
			$xRequestId,
			$output,
			$adminUsername,
			$adminPassword,
			$baseUrl,
			$ocPath
		)['stdOut'];
	}

	/**
	 * Finds all lines containing the given text
	 *
	 * @param string|null $input stdout or stderr output
	 * @param string|null $text text to search for
	 *
	 * @return array array of lines that matched
	 */
	public static function findLines(?string $input, ?string $text):array {
		$results = [];
		foreach (\explode("\n", $input) as $line) {
			if (\strpos($line, $text) !== false) {
				$results[] = $line;
			}
		}
		return $results;
	}

	/**
	 * Delete a system config setting
	 *
	 * @param string|null $key
	 * @param string|null $xRequestId
	 * @param string|null $adminUsername
	 * @param string|null $adminPassword
	 * @param string|null $baseUrl
	 * @param string|null $ocPath
	 *
	 * @return string[] associated array with "code", "stdOut", "stdErr"
	 * @throws GuzzleException if parameters have not been provided yet or the testing app is not enabled
	 */
	public static function deleteSystemConfig(
		?string $key,
		?string $xRequestId = '',
		?string $adminUsername = null,
		?string $adminPassword = null,
		?string $baseUrl = null,
		?string $ocPath = null
	):array {
		$args = [];
		$args[] = 'config:system:delete';
		$args[] = $key;

		$args[] = '--no-ansi';
		if ($baseUrl === null) {
			$baseUrl = self::$baseUrl;
		}
		return ['code' => '', 'stdOut' => '', 'stdErr' => '' ];
	}
}
