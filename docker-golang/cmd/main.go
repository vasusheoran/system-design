package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	if len(os.Args) < 2 { // Basic check for enough arguments
		panic("Usage: go-container run <command> [args...]")
	}
	switch os.Args[1] {
	case "run":
		parent()
	case "child":
		child()
	default:
		panic(fmt.Sprintf("Unknown command: %s. Available: run, child", os.Args[1]))
	}
}

func parent() {
	if len(os.Args) < 3 { // Ensure there's a command to run in the container
		panic("Usage: go-container run <command> [args...]")
	}
	// The first argument to child will be "child", followed by the actual command and its args
	childArgs := append([]string{"child"}, os.Args[2:]...)
	cmd := exec.Command("/proc/self/exe", childArgs...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | // New UTS namespace (hostname, domainname)
			syscall.CLONE_NEWPID | // New PID namespace (processes see their own PID tree starting from 1)
			syscall.CLONE_NEWNS, // New mount namespace (processes have their own view of the filesystem hierarchy)
		//Consider adding syscall.CLONE_NEWUSER for user namespaces for better security
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Printf("Parent: Running command %v with args %v as child process\n", childArgs[1], childArgs[2:])
	if err := cmd.Run(); err != nil {
		fmt.Println("Parent ERROR running child:", err)
		os.Exit(1)
	}
}

func child() {
	fmt.Printf("Child: Running %v\n", os.Args[2:])
	// Essential: Unshare mount namespace again for the child, if not already fully isolated
	// This is often needed if the parent wasn't already in a separate mount namespace.
	// However, CLONE_NEWNS in parent() should handle this. Let's assume it's fine.

	// --- Filesystem Layer Setup ---
	// These paths should exist or be created with appropriate permissions
	// For this example, assume they are relative to where your go-container binary is run
	// or are absolute paths.

	// 1. Define paths for OverlayFS
	baseImage := "base_image_rootfs" // Your read-only base layer (e.g., a minimal busybox rootfs)
	upperDir := "upperdir"           // Writable layer for container changes
	workDir := "workdir"             // OverlayFS internal working directory
	mergedDir := "rootfs"            // Where the merged view will be mounted (this will be the new root)

	// 2. Create the necessary directories if they don't exist
	// In a real scenario, baseImage would be pre-populated.
	// For this example, we'll just ensure the directories exist.
	// Ensure `baseImage` exists and is populated with a basic root filesystem (e.g., busybox)
	// For a quick test, you might create a dummy `baseImage/bin/sh` or `baseImage/bin/ls`
	must(os.MkdirAll(baseImage, 0755)) // Should be pre-existing with content
	must(os.MkdirAll(upperDir, 0755))
	must(os.MkdirAll(workDir, 0755))
	must(os.MkdirAll(mergedDir, 0755)) // This is what we will pivot_root into

	// 3. Mount OverlayFS
	// The options string for OverlayFS: "lowerdir=<path_to_lower>,upperdir=<path_to_upper>,workdir=<path_to_work>"
	// Get absolute paths for robustness in mount options
	absLower, err := os.Getwd() // Or use a fixed absolute path
	must(err)
	absUpper, err := os.Getwd()
	must(err)
	absWork, err := os.Getwd()
	must(err)

	// Adjust these paths according to your actual directory structure
	// For simplicity, assuming they are in the current working directory
	// In a real container runtime, these paths would be more rigorously managed.
	overlayData := fmt.Sprintf("lowerdir=%s/%s,upperdir=%s/%s,workdir=%s/%s",
		absLower, baseImage,
		absUpper, upperDir,
		absWork, workDir)

	fmt.Printf("Child: Mounting OverlayFS with data: %s to %s\n", overlayData, mergedDir)
	// MS_NODEV, MS_NOSUID are good security additions for general mounts
	mountFlags := uintptr(syscall.MS_NODEV | syscall.MS_NOSUID)
	must(syscall.Mount("overlay", mergedDir, "overlay", mountFlags, overlayData))
	fmt.Printf("Child: OverlayFS mounted on %s\n", mergedDir)

	// --- Pivot Root ---
	// Now that `mergedDir` (e.g., "rootfs") has our layered filesystem,
	// we can pivot into it.

	// Create a directory inside the new root to mount the old root onto.
	// The path must be relative to the new root.
	oldRootInnerPath := "oldrootfs" // This will be mergedDir + "/oldrootfs"
	must(os.MkdirAll(fmt.Sprintf("%s/%s", mergedDir, oldRootInnerPath), 0700))
	fmt.Printf("Child: Created %s/%s for old root\n", mergedDir, oldRootInnerPath)

	// Pivot root: new_root is `mergedDir`, put_old is `mergedDir + "/" + oldRootInnerPath`
	fmt.Printf("Child: Pivoting root to %s, putting old root at %s/%s\n", mergedDir, mergedDir, oldRootInnerPath)
	must(syscall.PivotRoot(mergedDir, fmt.Sprintf("%s/%s", mergedDir, oldRootInnerPath)))
	fmt.Println("Child: PivotRoot successful")

	// Change current working directory to the new root "/"
	must(os.Chdir("/"))
	fmt.Println("Child: Changed directory to /")

	// Unmount the old root.
	// The path to the old root is now ` "/" + oldRootInnerPath` from within the new root.
	// MS_DETACH is often used for lazy unmounting.
	fmt.Printf("Child: Unmounting old root at /%s\n", oldRootInnerPath)
	must(syscall.Unmount("/"+oldRootInnerPath, syscall.MNT_DETACH))
	must(os.RemoveAll("/" + oldRootInnerPath)) // Clean up the mount point for oldroot
	fmt.Printf("Child: Old root /%s unmounted and removed\n", oldRootInnerPath)

	// --- Setup standard mounts like /proc, /sys, /dev (Essential for most programs) ---
	// These should be mounted *after* pivot_root, inside the new root filesystem.
	// Note: Mounting /dev typically requires more setup (e.g., using devtmpfs or populating a minimal set of device nodes)
	// For simplicity, we'll mount /proc and /sys.
	// You might also need to mount /dev/pts for pseudo-terminals.

	// Mount /proc
	// Ensure /proc directory exists in your baseImage or create it now if it's guaranteed to be empty
	// However, it's better if baseImage already has these standard mountpoints.
	// If baseImage/proc doesn't exist, OverlayFS will create it in the upperdir upon write (like mkdir).
	must(os.MkdirAll("/proc", 0755)) // Create mountpoint if not in base image
	must(syscall.Mount("proc", "/proc", "proc", 0, ""))
	fmt.Println("Child: Mounted /proc")

	// Mount /sys
	must(os.MkdirAll("/sys", 0755)) // Create mountpoint if not in base image
	must(syscall.Mount("sysfs", "/sys", "sysfs", 0, ""))
	fmt.Println("Child: Mounted /sys")

	// (Optional) Mount /dev
	// must(os.MkdirAll("/dev", 0755))
	// must(syscall.Mount("tmpfs", "/dev", "tmpfs", syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755"))
	// // You might need to create basic device nodes like /dev/null, /dev/zero, /dev/tty, /dev/console etc.
	// // or mount a devtmpfs if your kernel supports it and you have permissions.
	// fmt.Println("Child: Mounted basic /dev (tmpfs)")

	// --- Execute the command ---
	fmt.Printf("Child: Executing command %s with args %v\n", os.Args[2], os.Args[3:])
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Set Hostname (UTS Namespace)
	must(syscall.Sethostname([]byte("my-container")))

	if err := cmd.Run(); err != nil {
		fmt.Printf("Child ERROR executing command %s: %v\n", os.Args[2], err)
		// Attempt to unmount filesystems before exiting to clean up
		// This is best-effort and might fail if the process is in a bad state.
		// syscall.Unmount("/proc", 0)
		// syscall.Unmount("/sys", 0)
		// syscall.Unmount("/", syscall.MNT_DETACH) // Unmount the OverlayFS from the *original* mount point if possible,
		// but this is tricky after pivot_root.
		// The namespace handles cleanup mostly.
		os.Exit(1)
	}
}

func must(err error) {
	if err != nil {
		panic(err) // In a real tool, log more context or handle specific errors
	}
}
