# EBarimt

# Шаардлага
1. `docker` - [install docker](https://docs.docker.com/engine/install/)
2. `docker-compose` - [install docker-compoose](https://docs.docker.com/compose/install/)
3. НӨАТУС-с авсан `libPosAPI.so` файл

# Тохиргоо

1. `data` folder үүсгэх
2. `poslib` folder үүсгэх
3. `poslib` дотор Татвараас ирсэн `.so` өргөтгөлтэй файлыг `libPosAPI.so` нэртэй хадгална.

Using docker

```
docker run -d \
  -p 8001:5000 \
  -v "$(pwd)"/data:/root/.vatps \
  -v "$(pwd)"/poslib:/poslib:ro \
  --name ebarimt registry.gitlab.com/endigo/ebarimt:latest
```

эсвэл

Using docker-compose
`docker-compose.yml` файл үүсгэнэ

```
version: "2"
services:
  web:
    image: registry.gitlab.com/endigo/ebarimt:latest
    ports:
      - 5000:5000
    restart: always
    volumes:
      - ./data:/root/.vatps
      - ./poslib:/poslib
```

Next step

```
docker-compose up -d
```

# Routes

1. `GET /health` - Сервис хэвийн ажиллаж байгаа эсэхийг шалгана.
2. `GET /checkApi` - POS API ажиллаж байгааг эсэхийг шалгана.

```
# Response
{
  "config": {
    "success": true
  },
  "database": {
    "success": true
  },
  "network": {
    "success": true
  },
  "success": true
}
```

Хэрэв дараах response ирсэн тохиолдолд /sendData -г дуудаж ажиллуулна уу!
```
{
    "config": {
        "message": "[100] Тохиргооны мэдээлэл олдсонгүй. Мэдээлэл илгээх хүсэлтийг ажиллуулж тохиргооны мэдээллүүдийг татна уу!!!",
        "success": false
    },
    "database": {
        "success": true
    },
    "network": {
        "success": true
    },
    "success": false
}
```
3. `GET /getInformation` -

4. `POST /callFunction` -
5. `POST /put` - Баримт хэвлэх үед ашиглана
6. `POST /returnBill` - Баримт буцаах
7. `POST /sendData` - Баримтуудыг илгээх, шинэ сугалааны дугаар явуулахад ашиглана. 24 цагт 1 удаа хамгийн багадаа ажиллаж ёстой

# Олон libPosApi.so ашиглах

`docker-compose.yml` файл үүсгэнэ

```
version: "2"
services:
  web:
    image: registry.gitlab.com/endigo/ebarimt:v1.1 # v1.1 image ашиглах
    ports:
      - 5000:5000
    restart: always
    volumes:
      - ./data:/root/.vatps
      - ./poslib:/poslib
```

`poslib` хавтас дотор доорх байдлаар `.so` файлуудаа хуулна.

```
ls -la ./ebarimt-service/poslib
-rwxr-xr-x 1 root root 94096 Jan  4 06:57 libPosApi.so # default. заавал байх ёстой
-rwxr-xr-x 1 root root 94096 Jan  4 06:57 organization1.so
-rwxr-xr-x 1 root root 94096 Jan  4 06:57 organization2.so
```


# Routes

Хүсэлт явуулахдаа `?lib={soFileName}` query parameter-г нэмж явуулах шаардлагтай.
Хэрвээ явуулаагүй тохиолдолд `libPosApi.so` файлыг ашиглах учир заавал энэ файлыг үүсгэнэ үү.

`GET /checkApi?lib={soFileName}` - POS API ажиллаж байгааг эсэхийг шалгана.
`GET /getInformation?lib={soFileName}`
`POST /callFunction?lib={soFileName}`
`POST /put?lib={soFileName}` - Баримт хэвлэх үед ашиглана
`POST /returnBill?lib={soFileName}` - Баримт буцаах
`POST /sendData?lib={soFileName}` - Баримтуудыг илгээх, шинэ сугалааны дугаар явуулахад ашиглана. 24 цагт 1 удаа хамгийн багадаа ажиллаж ёстой

## Example

```
  curl localhost:5000/checkApi?lib={organization1}
  curl localhost:5000/checkApi?lib={organization2}
```
