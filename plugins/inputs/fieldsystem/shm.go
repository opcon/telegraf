package fieldsystem

//Functions to read FS share memory

import (
	"errors"
	"strings"
	"unsafe"

	"svipc"
)

// TODO: These should be generated from "/usr2/fs/include/ipckeys.h"
const (
	SHM_PATH  = "/usr2/fs"
	CLS_PATH  = "/usr2/fs"
	SKD_PATH  = "/usr2/fs"
	BRK_PATH  = "/usr2/fs"
	SEM_PATH  = "/usr2/fs"
	NSEM_PATH = "/usr2/fs"
	GO_PATH   = "/usr2/fs"

	SHM_ID  = 1
	CLS_ID  = 2
	SKD_ID  = 3
	BRK_ID  = 4
	SEM_ID  = 5
	NSEM_ID = 6
	GO_ID   = 7
)

func GetSHM() (fs *Fscom, err error) {
	key, err := svipc.Ftok(SHM_PATH, SHM_ID)
	if err != nil {
		return
	}
	id, err := svipc.Shmget(key, unsafe.Sizeof(&Fscom{}), 0666)
	if err != nil {
		return
	}
	ptr, err := svipc.Shmat(id, 0, svipc.SHM_RDONLY)
	if err != nil {
		return
	}
	fs = (*Fscom)(ptr)
	return
}

// FS stores a list of names for semephores in the NSEM group. The list is in
// list in fs.Sem.  This function queies the semephones in that group by name.
func (fs *Fscom) SemLocked(semname string) (locked bool, err error) {
	key, err := svipc.Ftok(NSEM_PATH, NSEM_ID)
	if err != nil {
		return
	}

	sid, err := svipc.Semget(key, 0, 0)
	if err != nil {
		return
	}

	semnum := -1
	semname = strings.TrimSpace(semname)
	for i := 0; i < int(fs.Sem.Allocated); i++ {
		s := strings.TrimSpace(string(fs.Sem.Name[i][:]))
		if s == semname {
			semnum = i
			break
		}
	}
	if semnum == -1 {
		err = errors.New("sem not found")
		return
	}

	ret, err := svipc.Semctl(sid, semnum, svipc.GETVAL)
	if err != nil {
		return
	}
	// FS uses sems in a strange way
	locked = (int(ret) == 0)
	return
}
