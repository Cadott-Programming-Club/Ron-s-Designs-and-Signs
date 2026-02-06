module github.com/BubblePlayzTHEREAL/Ron-s-Designs-and-Signs/api

go 1.24

// Replace the main module with the parent directory so this function
// module can import local packages from the repository during per-function
// builds (Vercel runs `go` inside the `api/` folder).
replace github.com/BubblePlayzTHEREAL/Ron-s-Designs-and-Signs => ../
