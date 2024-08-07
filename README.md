
# VR Merge

So I was farting around with Pigasus and making VR images in Blender and Daz Studio.

I was tired of combining the images in an image editor and saving them out.  So being the nerd I am, I wrote VR_Merge to do it for me.  I am learning GoLang, so I took this as an opportunity to increase my skillpoints in GoLang.

This auto-appends the _3DHF / _3DVF to the filename.  3d Viewing programs will detect this extension and adjust settings automatically.  IF you have 360 degree spherical projection images, you will need to add _360 to the end yourself. (or just add it in the output.)

Usage:

```bash
  go run main.go <image1> <image2> <output>
```

Output does not need an extension!

Example:
```bash
go run main.go scene_left.jpg scene_right.jpg scenery
```

# TODO:

I will eventually create the pipelines to automatically build windows/osx/linux binaries. But until I do, here you go.
