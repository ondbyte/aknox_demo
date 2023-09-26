### demo by yadunandan for accuknox

### to run using docker

at the root of the project, to build docker image run
```docker build -t aknox_demo .```

and to run the image
```docker run -p 8000:8000 aknox_demo```

server will be available at ```localhost:8000```
api is similar to the blueprint given in the assignment pdf

### available endpoints
/login (POST)
/signup (POST)
/notes(GET,POST,DELETE)

if possible i'll add more documentation