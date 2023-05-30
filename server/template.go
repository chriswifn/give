package server

var uploadTemplate = `
<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1.0" />
    <meta http-equiv="X-UA-Compatible" content="ie=edge" />
    <title>Upload File</title>
    <style>
    html { text-align: center; }
    </style>
  </head>
  <body>
    <form
      enctype="multipart/form-data"
      action="/upload"
      method="post"
    >
      <input type="file" name="myFile" /><br>
      <input type="submit" value="upload" />
    </form>
  </body>
</html>
  `
