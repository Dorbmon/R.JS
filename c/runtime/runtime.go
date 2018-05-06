package runtime
import(
	"os/exec"
	"log"
	"os"
	"path/filepath"
	"fmt"
)
func GetGolangPathWithName()string{
	execPath, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	//    Is Symlink
	fi, err := os.Lstat(execPath)
	if err != nil {
		log.Fatal(err)
	}
	if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
		execPath, err = os.Readlink(execPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	execDir := filepath.Dir(execPath)
	if execDir == "." {
		execDir, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}
	return execPath
}
func GetGolangPath()string{
	execPath, err := exec.LookPath(os.Args[0])
	if err != nil {
		log.Fatal(err)
	}
	//    Is Symlink
	fi, err := os.Lstat(execPath)
	if err != nil {
		log.Fatal(err)
	}
	if fi.Mode()&os.ModeSymlink == os.ModeSymlink {
		execPath, err = os.Readlink(execPath)
		if err != nil {
			log.Fatal(err)
		}
	}
	execDir := filepath.Dir(execPath)
	if execDir == "." {
		execDir, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}
	return execDir
}