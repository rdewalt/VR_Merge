
# VR Merge

So I was farting around with Pigasus and making VR images in Blender and Daz Studio.

I was tired of combining the images in an image editor and saving them out.  So being the nerd I am, I wrote VR_Merge to do it for me.  I am learning GoLang, so I took this as an opportunity to increase my skillpoints in 

Usage:

```bash
  go run main.go <image1> <image2> <output>
```

Output does not need an extension!

Example:
```bash
go run main.go scene_left.jpg scene_right.jpg scenery
```

I will eventually create the pipelines to automatically build windows/osx/linux binaries. But until I do, here you go.
