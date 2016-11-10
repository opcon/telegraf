package fieldsystem

//Functions to read FS share memory

/*
#include <stdlib.h>
#include <string.h>
#include <sys/types.h>
#include <sys/ipc.h>
#include <sys/shm.h>
#include <sys/sem.h>

key_t ftok(const char *pathname, int proj_id);
int shmget(key_t key, size_t size, int shmflg);
void *shmat(int shmid, const void *shmaddr, int shmflg);
int shmdt(const void *shmaddr);
int shmctl(int shmid, int cmd, struct shmid_ds *buf);

int get_sem_val(int sid, int semnum) {
        return(semctl(sid, semnum, GETVAL, 0));
}

*/
import "C"

import (
	"errors"
	"fmt"
	"strings"
	"unsafe"
)

// TODO: These should be generated from "/usr2/fs/include/ipckeys.h"
const (
	SHM_PATH  = "/usr2/fs"
	SHM_ID    = 1
	CLS_PATH  = "/usr2/fs"
	CLS_ID    = 2
	SKD_PATH  = "/usr2/fs"
	SKD_ID    = 3
	BRK_PATH  = "/usr2/fs"
	BRK_ID    = 4
	SEM_PATH  = "/usr2/fs"
	SEM_ID    = 5
	NSEM_PATH = "/usr2/fs"
	NSEM_ID   = 6
	GO_PATH   = "/usr2/fs"
	GO_ID     = 7
)

func GetSHM() (*Fscom, error) {
	cpath := C.CString(SHM_PATH)
	defer C.free(unsafe.Pointer(cpath))
	rckey, err := C.ftok(cpath, C.int(SHM_ID))
	if int64(rckey) == -1 {
		return nil, err
	}
	id, err := C.shmget(rckey, C.size_t(unsafe.Sizeof(&Fscom{})), C.int(0444))
	if int64(id) == -1 {
		return nil, err
	}
	ptr, err := C.shmat(id, nil, C.int(C.SHM_RDONLY))
	if *(*int)(unsafe.Pointer(ptr)) == -1 {
		return nil, err
	}
	return (*Fscom)(ptr), nil
}

// FS stores a list of names for semephores in the NSEM group. The list is in
// list in fs.Sem.  This function queies the semephones in that group by name.
func (fs *Fscom) SemLocked(semname string) (bool, error) {
	cpath := C.CString(NSEM_PATH)
	defer C.free(unsafe.Pointer(cpath))
	key, err := C.ftok(cpath, C.int(NSEM_ID))
	if int64(key) == -1 || err != nil {
		return false, err
	}

	sid, err := C.semget(key, 0, 0)
	if int(sid) < 0 || err != nil {
		return false, err
	}
	fmt.Printf("%d\n", sid)
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
		return false, errors.New("sem not found")
	}
	ret, err := C.get_sem_val(sid, C.int(semnum))
	if err != nil {
		panic(err)
	}
	// FS uses sems in a strange way
	return int(ret) == 0, nil
}
