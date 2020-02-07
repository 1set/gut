/*
Package yos is yet another wrapper of platform-independent interface to operating system functionality.


System Info

Check architecture of current system:
  - IsOn32bitArch
  - IsOn64bitArch

Check current operating system:
  - IsOnLinux
  - IsOnMacOS
  - IsOnMacOS


File Operations

Basic operations:

	| Ops \ Type  | Directory      | File            | Symbolic Link      |
	| ----------- | -------------- | --------------- | ------------------ |
	| Check Exist | ExistDir       | ExistFile       | ExistSymlink       |
	| Check Empty | IsDirEmpty     | IsFileEmpty     | -                  |
	| Check Same  | SameDirEntries | SameFileContent | SameSymlinkContent |
	| Get Size    | GetDirSize     | GetFileSize     | GetSymlinkSize     |
	| List        | ListDir        | ListFile        | ListSymlink        |
	| Copy        | CopyDir        | CopyFile        | CopySymlink        |
	| Move        | MoveDir        | MoveFile        | MoveSymlink        |

Miscellaneous operations:
  - ListMatch
  - JoinPath
  - Exist
  - NotExist

Sorting helpers for a slice of *FilePathInfo:
  - SortListByName
  - SortListBySize
  - SortListByModTime

*/
package yos
