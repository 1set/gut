/*
Package yos is yet another wrapper of platform-independent interface to operating system functionality.


Environments

Check architecture of current system:
  - IsOn32bitArch
  - IsOn64bitArch

Check current operating system:
  - IsOnLinux
  - IsOnMacOS
  - IsOnMacOS

Change working directory:
  - ChangeExeDir


File Operations

Basic operations:

	| Ops \ Type  | Directory      | File            | Symbolic Link      |
	| ----------- | -------------- | --------------- | ------------------ |
	| Check Exist | ExistDir       | ExistFile       | ExistSymlink       |
	| Check Empty | IsDirEmpty     | IsFileEmpty     | -                  |
	| Get Size    | GetDirSize     | GetFileSize     | GetSymlinkSize     |
	| Compare     | SameDirEntries | SameFileContent | SameSymlinkContent |
	| List        | ListDir        | ListFile        | ListSymlink        |
	| Copy        | CopyDir        | CopyFile        | CopySymlink        |
	| Move        | MoveDir        | MoveFile        | MoveSymlink        |

Miscellaneous operations:
  - ListMatch
  - JoinPath
  - Exist
  - NotExist
  - MakeDir

Sorting helpers for a slice of *FilePathInfo:
  - SortListByName
  - SortListBySize
  - SortListByModTime

*/
package yos
