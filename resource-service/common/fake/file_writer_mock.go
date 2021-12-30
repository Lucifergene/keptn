// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package common_mock

import (
	"path/filepath"
	"sync"
)

// IFileSystemMock is a mock implementation of common.IFileSystem.
//
// 	func TestSomethingThatUsesIFileSystem(t *testing.T) {
//
// 		// make and configure a mocked common.IFileSystem
// 		mockedIFileSystem := &IFileSystemMock{
// 			DeleteFileFunc: func(path string) error {
// 				panic("mock out the DeleteFile method")
// 			},
// 			FileExistsFunc: func(path string) bool {
// 				panic("mock out the FileExists method")
// 			},
// 			MakeDirFunc: func(path string) error {
// 				panic("mock out the MakeDir method")
// 			},
// 			ReadFileFunc: func(filename string) ([]byte, error) {
// 				panic("mock out the ReadFile method")
// 			},
// 			WalkPathFunc: func(path string, walkFunc filepath.WalkFunc) error {
// 				panic("mock out the WalkPath method")
// 			},
// 			WriteBase64EncodedFileFunc: func(path string, content string) error {
// 				panic("mock out the WriteBase64EncodedFile method")
// 			},
// 			WriteFileFunc: func(path string, content []byte) error {
// 				panic("mock out the WriteFile method")
// 			},
// 		}
//
// 		// use mockedIFileSystem in code that requires common.IFileSystem
// 		// and then make assertions.
//
// 	}
type IFileSystemMock struct {
	// DeleteFileFunc mocks the DeleteFile method.
	DeleteFileFunc func(path string) error

	// FileExistsFunc mocks the FileExists method.
	FileExistsFunc func(path string) bool

	// MakeDirFunc mocks the MakeDir method.
	MakeDirFunc func(path string) error

	// ReadFileFunc mocks the ReadFile method.
	ReadFileFunc func(filename string) ([]byte, error)

	// WalkPathFunc mocks the WalkPath method.
	WalkPathFunc func(path string, walkFunc filepath.WalkFunc) error

	// WriteBase64EncodedFileFunc mocks the WriteBase64EncodedFile method.
	WriteBase64EncodedFileFunc func(path string, content string) error

	// WriteFileFunc mocks the WriteFile method.
	WriteFileFunc func(path string, content []byte) error

	// calls tracks calls to the methods.
	calls struct {
		// DeleteFile holds details about calls to the DeleteFile method.
		DeleteFile []struct {
			// Path is the path argument value.
			Path string
		}
		// FileExists holds details about calls to the FileExists method.
		FileExists []struct {
			// Path is the path argument value.
			Path string
		}
		// MakeDir holds details about calls to the MakeDir method.
		MakeDir []struct {
			// Path is the path argument value.
			Path string
		}
		// ReadFile holds details about calls to the ReadFile method.
		ReadFile []struct {
			// Filename is the filename argument value.
			Filename string
		}
		// WalkPath holds details about calls to the WalkPath method.
		WalkPath []struct {
			// Path is the path argument value.
			Path string
			// WalkFunc is the walkFunc argument value.
			WalkFunc filepath.WalkFunc
		}
		// WriteBase64EncodedFile holds details about calls to the WriteBase64EncodedFile method.
		WriteBase64EncodedFile []struct {
			// Path is the path argument value.
			Path string
			// Content is the content argument value.
			Content string
		}
		// WriteFile holds details about calls to the WriteFile method.
		WriteFile []struct {
			// Path is the path argument value.
			Path string
			// Content is the content argument value.
			Content []byte
		}
	}
	lockDeleteFile             sync.RWMutex
	lockFileExists             sync.RWMutex
	lockMakeDir                sync.RWMutex
	lockReadFile               sync.RWMutex
	lockWalkPath               sync.RWMutex
	lockWriteBase64EncodedFile sync.RWMutex
	lockWriteFile              sync.RWMutex
}

// DeleteFile calls DeleteFileFunc.
func (mock *IFileSystemMock) DeleteFile(path string) error {
	if mock.DeleteFileFunc == nil {
		panic("IFileSystemMock.DeleteFileFunc: method is nil but IFileSystem.DeleteFile was just called")
	}
	callInfo := struct {
		Path string
	}{
		Path: path,
	}
	mock.lockDeleteFile.Lock()
	mock.calls.DeleteFile = append(mock.calls.DeleteFile, callInfo)
	mock.lockDeleteFile.Unlock()
	return mock.DeleteFileFunc(path)
}

// DeleteFileCalls gets all the calls that were made to DeleteFile.
// Check the length with:
//     len(mockedIFileSystem.DeleteFileCalls())
func (mock *IFileSystemMock) DeleteFileCalls() []struct {
	Path string
} {
	var calls []struct {
		Path string
	}
	mock.lockDeleteFile.RLock()
	calls = mock.calls.DeleteFile
	mock.lockDeleteFile.RUnlock()
	return calls
}

// FileExists calls FileExistsFunc.
func (mock *IFileSystemMock) FileExists(path string) bool {
	if mock.FileExistsFunc == nil {
		panic("IFileSystemMock.FileExistsFunc: method is nil but IFileSystem.FileExists was just called")
	}
	callInfo := struct {
		Path string
	}{
		Path: path,
	}
	mock.lockFileExists.Lock()
	mock.calls.FileExists = append(mock.calls.FileExists, callInfo)
	mock.lockFileExists.Unlock()
	return mock.FileExistsFunc(path)
}

// FileExistsCalls gets all the calls that were made to FileExists.
// Check the length with:
//     len(mockedIFileSystem.FileExistsCalls())
func (mock *IFileSystemMock) FileExistsCalls() []struct {
	Path string
} {
	var calls []struct {
		Path string
	}
	mock.lockFileExists.RLock()
	calls = mock.calls.FileExists
	mock.lockFileExists.RUnlock()
	return calls
}

// MakeDir calls MakeDirFunc.
func (mock *IFileSystemMock) MakeDir(path string) error {
	if mock.MakeDirFunc == nil {
		panic("IFileSystemMock.MakeDirFunc: method is nil but IFileSystem.MakeDir was just called")
	}
	callInfo := struct {
		Path string
	}{
		Path: path,
	}
	mock.lockMakeDir.Lock()
	mock.calls.MakeDir = append(mock.calls.MakeDir, callInfo)
	mock.lockMakeDir.Unlock()
	return mock.MakeDirFunc(path)
}

// MakeDirCalls gets all the calls that were made to MakeDir.
// Check the length with:
//     len(mockedIFileSystem.MakeDirCalls())
func (mock *IFileSystemMock) MakeDirCalls() []struct {
	Path string
} {
	var calls []struct {
		Path string
	}
	mock.lockMakeDir.RLock()
	calls = mock.calls.MakeDir
	mock.lockMakeDir.RUnlock()
	return calls
}

// ReadFile calls ReadFileFunc.
func (mock *IFileSystemMock) ReadFile(filename string) ([]byte, error) {
	if mock.ReadFileFunc == nil {
		panic("IFileSystemMock.ReadFileFunc: method is nil but IFileSystem.ReadFile was just called")
	}
	callInfo := struct {
		Filename string
	}{
		Filename: filename,
	}
	mock.lockReadFile.Lock()
	mock.calls.ReadFile = append(mock.calls.ReadFile, callInfo)
	mock.lockReadFile.Unlock()
	return mock.ReadFileFunc(filename)
}

// ReadFileCalls gets all the calls that were made to ReadFile.
// Check the length with:
//     len(mockedIFileSystem.ReadFileCalls())
func (mock *IFileSystemMock) ReadFileCalls() []struct {
	Filename string
} {
	var calls []struct {
		Filename string
	}
	mock.lockReadFile.RLock()
	calls = mock.calls.ReadFile
	mock.lockReadFile.RUnlock()
	return calls
}

// WalkPath calls WalkPathFunc.
func (mock *IFileSystemMock) WalkPath(path string, walkFunc filepath.WalkFunc) error {
	if mock.WalkPathFunc == nil {
		panic("IFileSystemMock.WalkPathFunc: method is nil but IFileSystem.WalkPath was just called")
	}
	callInfo := struct {
		Path     string
		WalkFunc filepath.WalkFunc
	}{
		Path:     path,
		WalkFunc: walkFunc,
	}
	mock.lockWalkPath.Lock()
	mock.calls.WalkPath = append(mock.calls.WalkPath, callInfo)
	mock.lockWalkPath.Unlock()
	return mock.WalkPathFunc(path, walkFunc)
}

// WalkPathCalls gets all the calls that were made to WalkPath.
// Check the length with:
//     len(mockedIFileSystem.WalkPathCalls())
func (mock *IFileSystemMock) WalkPathCalls() []struct {
	Path     string
	WalkFunc filepath.WalkFunc
} {
	var calls []struct {
		Path     string
		WalkFunc filepath.WalkFunc
	}
	mock.lockWalkPath.RLock()
	calls = mock.calls.WalkPath
	mock.lockWalkPath.RUnlock()
	return calls
}

// WriteBase64EncodedFile calls WriteBase64EncodedFileFunc.
func (mock *IFileSystemMock) WriteBase64EncodedFile(path string, content string) error {
	if mock.WriteBase64EncodedFileFunc == nil {
		panic("IFileSystemMock.WriteBase64EncodedFileFunc: method is nil but IFileSystem.WriteBase64EncodedFile was just called")
	}
	callInfo := struct {
		Path    string
		Content string
	}{
		Path:    path,
		Content: content,
	}
	mock.lockWriteBase64EncodedFile.Lock()
	mock.calls.WriteBase64EncodedFile = append(mock.calls.WriteBase64EncodedFile, callInfo)
	mock.lockWriteBase64EncodedFile.Unlock()
	return mock.WriteBase64EncodedFileFunc(path, content)
}

// WriteBase64EncodedFileCalls gets all the calls that were made to WriteBase64EncodedFile.
// Check the length with:
//     len(mockedIFileSystem.WriteBase64EncodedFileCalls())
func (mock *IFileSystemMock) WriteBase64EncodedFileCalls() []struct {
	Path    string
	Content string
} {
	var calls []struct {
		Path    string
		Content string
	}
	mock.lockWriteBase64EncodedFile.RLock()
	calls = mock.calls.WriteBase64EncodedFile
	mock.lockWriteBase64EncodedFile.RUnlock()
	return calls
}

// WriteFile calls WriteFileFunc.
func (mock *IFileSystemMock) WriteFile(path string, content []byte) error {
	if mock.WriteFileFunc == nil {
		panic("IFileSystemMock.WriteFileFunc: method is nil but IFileSystem.WriteFile was just called")
	}
	callInfo := struct {
		Path    string
		Content []byte
	}{
		Path:    path,
		Content: content,
	}
	mock.lockWriteFile.Lock()
	mock.calls.WriteFile = append(mock.calls.WriteFile, callInfo)
	mock.lockWriteFile.Unlock()
	return mock.WriteFileFunc(path, content)
}

// WriteFileCalls gets all the calls that were made to WriteFile.
// Check the length with:
//     len(mockedIFileSystem.WriteFileCalls())
func (mock *IFileSystemMock) WriteFileCalls() []struct {
	Path    string
	Content []byte
} {
	var calls []struct {
		Path    string
		Content []byte
	}
	mock.lockWriteFile.RLock()
	calls = mock.calls.WriteFile
	mock.lockWriteFile.RUnlock()
	return calls
}
